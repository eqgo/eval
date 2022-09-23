package eval

import "errors"

// A Stage is an evaluation stage
type Stage struct {
	Left   *Stage
	Right  *Stage
	Symbol Token
}

// Stages returns the stages for the given tokens
func Stages(t []Token) (*Stage, error) {
	p := newParser(t)
	err := p.parse()
	return p.stg, err
}

// Eval evaluates the stage
func (s *Stage) Eval() (any, error) {
	switch s.Symbol.Type {
	case NUMOP:
	case NUMPRE:
	case COMP:
	case LOGOP:
	case LOGPRE:
	}
	return nil, errors.New("symbol must be symbol")
}
