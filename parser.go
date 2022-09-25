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

var parseSep, parseAdd, parseMul, parseMod, parsePow parseRule

func init() {
	parsePow = makeOpParseFunc([]Token{{NUMOP, POW}}, parseFunc, parseFunc)
	parseMod = makeOpParseFunc([]Token{{NUMOP, MOD}}, parsePow, nil)
	parseMul = makeOpParseFunc([]Token{{NUMOP, MUL}, {NUMOP, DIV}}, parseMod, nil)
	parseAdd = makeOpParseFunc([]Token{{NUMOP, ADD}, {NUMOP, SUB}}, parseMul, nil)
	parseSep = makeOpParseFunc([]Token{{SEP, nil}}, parseAdd, nil)
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
		leftStage, err := leftRule(p)
		if err != nil {
			return nil, err
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
		return &stage{evalFunc: litStage(tok)}, nil
	case VAR:
		return &stage{evalFunc: varStage(tok.Value.(string))}, nil
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
		stg, err := p.parse()
		if err != nil {
			return nil, err
		}
		return &stage{
			right:    stg,
			evalFunc: negStage, // negate is currently only num prefix
		}, nil
	}
	return nil, errors.New("parser: unrecognized token")
}
