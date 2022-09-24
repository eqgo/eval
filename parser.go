package eval

import "fmt"

type parser struct {
	src []Token
	pos int
	len int
	stg *Stage
}

type parseRule func(p *parser) (*Stage, error)

var parseSep, parseAdd, parseMul, parseMod, parsePow parseRule

func init() {
	parsePow = makeOpParseFunc([]Token{{NUMOP, POW}}, parseFunc)
	parseMod = makeOpParseFunc([]Token{{NUMOP, MOD}}, parsePow)
	parseMul = makeOpParseFunc([]Token{{NUMOP, MUL}, {NUMOP, DIV}}, parseMod)
	parseAdd = makeOpParseFunc([]Token{{NUMOP, ADD}, {NUMOP, SUB}}, parseMul)
	parseSep = makeOpParseFunc([]Token{{SEP, nil}}, parseAdd)
}

func newParser(src []Token) *parser {
	return &parser{src: src, pos: 0, len: len(src), stg: nil}
}

// parse goes through tokens and sets stg to result
func (p *parser) parse() error {
	stg, err := parseSep(p)
	if err != nil {
		return err
	}
	p.stg = stg
	return nil
}

// makeOpParseFunc makes an a simple parse func for an operator situation
func makeOpParseFunc(tokens []Token, next parseRule) parseRule {
	return func(p *parser) (*Stage, error) {
		for _, t := range tokens {
			for p.pos = 0; p.pos < p.len; p.pos++ {
				cur := p.src[p.pos]
				if cur == t {
					pl := newParser(p.src[:p.pos])
					pl.parse()
					pr := newParser(p.src[p.pos+1:])
					pr.parse()
					return &Stage{
						Left:  pl.stg,
						Right: pr.stg,
						eval:  tokenStageEvalMap[t],
					}, nil
				}
			}
		}
		p.pos = 0
		return next(p)
	}

}

func parseFunc(p *parser) (*Stage, error) {
	fmt.Println(p.src)
	if p.src[0].Type != FUNC {
		return parseVal(p)
	}
	pr := newParser(p.src[p.pos+1:])
	err := pr.parse()
	if err != nil {
		return nil, err
	}
	return &Stage{
		Right: pr.stg,
		eval:  functionStage(p.src[0].Value.(string)),
	}, nil
}

func parseVal(p *parser) (*Stage, error) {
	fmt.Println(p.src)
	tok := p.src[0]
	switch tok.Type {
	case NUM, BOOL:
		return &Stage{eval: litStage(tok)}, nil
	case VAR:
		return &Stage{eval: varStage(tok.Value.(string))}, nil
	case LEFT:
		fmt.Println(p.src)
		pr := newParser(p.src[p.pos+1 : p.len-1])
		err := pr.parse()
		if err != nil {
			return nil, err
		}
		return pr.stg, nil

	case RIGHT:
		return nil, nil
	}
	return nil, nil
}
