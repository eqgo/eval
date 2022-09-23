package eval

// A Stage is an evaluation stage
type Stage struct {
	Left  *Stage
	Right *Stage
	eval  stageEval
}

// Stages returns the stages for the given tokens
func Stages(t []Token) (*Stage, error) {
	p := newParser(t)
	err := p.parse()
	return p.stg, err
}

// Eval evaluates the stage
func (s *Stage) Eval(ctx *Context) (any, error) {
	var left, right any
	var err error

	if s.Left != nil {
		left, err = s.Left.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}

	if s.Right != nil {
		right, err = s.Right.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}

	return s.eval(left, right, ctx)
}
