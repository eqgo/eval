package eval

import "errors"

type parser struct {
	src []token
	pos int
	len int
	stg *stage
}

type parseRule func(p *parser) (*stage, error)

var parseSep, parseAdd, parseMul, parseMod, parsePow parseRule

func init() {
	parsePow = makeOpParseFunc([]token{{NUMOP, POW}}, parseFunc)
	parseMod = makeOpParseFunc([]token{{NUMOP, MOD}}, parsePow)
	parseMul = makeOpParseFunc([]token{{NUMOP, MUL}, {NUMOP, DIV}}, parseMod)
	parseAdd = makeOpParseFunc([]token{{NUMOP, ADD}, {NUMOP, SUB}}, parseMul)
	parseSep = makeOpParseFunc([]token{{SEP, nil}}, parseVal)
}

func newParser(src []token) *parser {
	return &parser{src: src, pos: 0, len: len(src), stg: nil}
}

// parse goes through tokens and sets stg to result
func (p *parser) parse() error {
	if len(p.src) == 0 {
		return errors.New("parser: syntax error")
	}
	stg, err := parseAdd(p)
	if err != nil {
		return err
	}
	p.stg = stg
	return nil
}

// makeOpParseFunc makes an a simple parse func for an operator situation
func makeOpParseFunc(tokens []token, next parseRule) parseRule {
	return func(p *parser) (*stage, error) {
		for _, t := range tokens {
			for p.pos = 0; p.pos < p.len; p.pos++ {
				cur := p.src[p.pos]
				if cur == t {
					pl := newParser(p.src[:p.pos])
					err := pl.parse()
					if err != nil {
						return nil, err
					}
					pr := newParser(p.src[p.pos+1:])
					err = pr.parse()
					if err != nil {
						return nil, err
					}
					return &stage{
						left:     pl.stg,
						right:    pr.stg,
						evalFunc: tokenStageEvalMap[t],
					}, nil
				}
			}
		}
		p.pos = 0
		return next(p)
	}

}

func parseFunc(p *parser) (*stage, error) {
	if p.src[0].typ != FUNC {
		return parseSep(p)
	}
	pr := newParser(p.src[p.pos+1:])
	err := pr.parse()
	if err != nil {
		return nil, err
	}
	return &stage{
		right:    pr.stg,
		evalFunc: funcStage(p.src[0].value.(string)),
	}, nil
}

func parseVal(p *parser) (*stage, error) {
	tok := p.src[0]
	switch tok.typ {
	case NUM, BOOL:
		return &stage{evalFunc: litStage(tok)}, nil
	case VAR:
		return &stage{evalFunc: varStage(tok.value.(string))}, nil
	case LEFT:
		pr := newParser(p.src[p.pos+1:])
		err := pr.parse()
		if err != nil {
			return nil, err
		}
		return pr.stg, nil
	case RIGHT:
		return nil, nil
	case NUMPRE:
		pr := newParser(p.src[p.pos+1:])
		err := pr.parse()
		if err != nil {
			return nil, err
		}
		return &stage{
			right:    pr.stg,
			evalFunc: negStage, // negate is currently only num prefix
		}, nil
	}
	return nil, errors.New("parser: unrecognized token")
}
