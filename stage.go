package eval

// A stage is an evaluation stage
type stage struct {
	left     *stage
	right    *stage
	evalFunc stageEval
}

// stages returns the stages for the given tokens
func stages(t []token) (*stage, error) {
	p := newParser(t)
	err := p.parse()
	return p.stg, err
}

// eval evaluates the stage
func (s *stage) eval(ctx *Context) (any, error) {
	var left, right any
	var err error

	if s.left != nil {
		left, err = s.left.eval(ctx)
		if err != nil {
			return nil, err
		}
	}

	if s.right != nil {
		right, err = s.right.eval(ctx)
		if err != nil {
			return nil, err
		}
	}

	return s.evalFunc(left, right, ctx)
}
