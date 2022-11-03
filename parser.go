package eval

import (
	"errors"
)

type parser struct {
	src []Token
	pos int
	len int
}

type parseRule func(p *parser) (*stage, error)

var parseSep, parseOr, parseAnd, parseComp, parseAdd, parseMul, parseMod, parsePow, parsePre parseRule

func init() {
	parsePre = makeOpParseFunc([]Token{{NUMPRE, NEG}, {LOGPRE, NOT}}, nil, parseFunc)
	parsePow = makeOpParseFunc([]Token{{NUMOP, POW}}, parseFunc, nil)
	parseMod = makeOpParseFunc([]Token{{NUMOP, MOD}}, parsePow, nil)
	parseMul = makeOpParseFunc([]Token{{NUMOP, MUL}, {NUMOP, DIV}}, parseMod, nil)
	parseAdd = makeOpParseFunc([]Token{{NUMOP, ADD}, {NUMOP, SUB}}, parseMul, nil)
	parseComp = makeOpParseFunc([]Token{
		{COMP, EQUAL}, {COMP, NOTEQUAL}, {COMP, GREATER}, {COMP, LESS}, {COMP, GEQ}, {COMP, LEQ},
	}, parseAdd, nil)
	parseAnd = makeOpParseFunc([]Token{{LOGOP, AND}}, parseComp, nil)
	parseOr = makeOpParseFunc([]Token{{LOGOP, OR}}, parseAnd, nil)
	parseSep = makeOpParseFunc([]Token{{SEP, nil}}, parseOr, nil)
}

func newParser(src []Token) *parser {
	return &parser{src: src, pos: 0, len: len(src)}
}

func (p *parser) next() Token {
	t := p.src[p.pos]
	p.pos++
	return t
}

func (p *parser) hasNext() bool {
	return p.pos < p.len
}

func (p *parser) rewind() {
	p.pos--
}

// parse goes through tokens and returns stg
func (p *parser) parse() (*stage, error) {
	if p.len == 0 {
		return nil, errors.New("parser: syntax error")
	}
	stg, err := parseSep(p)
	if err != nil {
		return nil, err
	}
	return stg, nil
}

// makeOpParseFunc makes an a simple parse func for an operator situation
func makeOpParseFunc(tokens []Token, leftRule parseRule, rightRule parseRule) parseRule {
	var f parseRule
	f = func(p *parser) (*stage, error) {
		var leftStage *stage
		if leftRule != nil {
			var err error
			leftStage, err = leftRule(p)
			if err != nil {
				return nil, err
			}
		}
		for p.hasNext() {
			cur := p.next()
			found := false
			for _, t := range tokens {
				if t == cur {
					found = true
					break
				}
			}
			if !found {
				break
			}
			rightStage, err := rightRule(p)
			if err != nil {
				return nil, err
			}
			stg := &stage{
				left:     leftStage,
				right:    rightStage,
				tok:      cur,
				evalFunc: tokenStageEvalMap[cur],
			}
			return stg, nil

		}
		p.rewind()
		return leftStage, nil
	}
	if rightRule == nil {
		rightRule = f
	}
	return f
}

func parseFunc(p *parser) (*stage, error) {
	if !p.hasNext() {
		return nil, errors.New("parser: syntax error")
	}
	tok := p.next()
	if tok.Type != FUNC {
		p.rewind()
		return parseVal(p)
	}
	stg, err := parseVal(p)
	if err != nil {
		return nil, err
	}
	return &stage{
		right:    stg,
		tok:      tok,
		evalFunc: funcStage(tok.Value.(string)),
	}, nil
}

func parseVal(p *parser) (*stage, error) {
	if !p.hasNext() {
		return nil, errors.New("parser: syntax error")
	}
	tok := p.next()
	switch tok.Type {
	case NUM, BOOL:
		return &stage{tok: tok, evalFunc: litStage(tok)}, nil
	case VAR:
		return &stage{tok: tok, evalFunc: varStage(tok.Value.(string))}, nil
	case LEFT:
		stg, err := p.parse()
		if err != nil {
			return nil, err
		}
		p.next()
		return stg, nil
	case RIGHT:
		return nil, nil
	case NUMPRE:
		p.rewind()
		return parsePre(p)
	case LOGPRE:
		p.rewind()
		return parsePre(p)
	}
	return nil, errors.New("parser: unrecognized token")
}
