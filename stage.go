package eval

// A stage is an evaluation stage
type stage struct {
	left     *stage
	right    *stage
	tok      Token
	evalFunc stageEval
}

// stages returns the stages for the given tokens
func stages(t []Token) (*stage, error) {
	p := newParser(t)
	stg, err := p.parse()
	stg.fixSamePrec()
	return stg, err
}

// fixSamePrec fixes stages of equal precedence that were parsed in reverse
func (s *stage) fixSamePrec() {
	var curStage *stage
	var curPrec opPrec

	var samePrecs []*stage

	nextStage := s
	prec := s.tok.opPrec()

	for nextStage != nil {
		curStage = nextStage
		nextStage = curStage.right

		if curStage.left != nil {
			curStage.left.fixSamePrec()
		}

		curPrec = curStage.tok.opPrec()

		if curPrec == prec {
			samePrecs = append(samePrecs, curStage)
			continue
		}

		if len(samePrecs) > 1 {
			mirrorStages(samePrecs)
		}

		samePrecs = []*stage{curStage}
		prec = curPrec
	}

	if len(samePrecs) > 1 {
		mirrorStages(samePrecs)
	}
}

// mirrorStages mirrors the stages
func mirrorStages(stages []*stage) {
	var rootStage, inverseStage, carryStage, frontStage *stage

	stagesLen := len(stages)

	for _, frontStage = range stages {
		carryStage = frontStage.right
		frontStage.right = frontStage.left
		frontStage.left = carryStage
	}

	rootStage = stages[0]
	frontStage = stages[stagesLen-1]

	carryStage = frontStage.left
	frontStage.left = rootStage.right
	rootStage.right = carryStage

	for i := 0; i < (stagesLen-2)/2+1; i++ {
		frontStage = stages[i+1]
		inverseStage = stages[stagesLen-i-1]
		carryStage = frontStage.right
		frontStage.right = inverseStage.right
		inverseStage.right = carryStage
	}

	for i := 0; i < stagesLen/2; i++ {
		frontStage = stages[i]
		inverseStage = stages[stagesLen-i-1]

		frontStage.swapTokAndFunc(inverseStage)
	}
}

// swapTokAndFunc swaps the token and evalFunc of the two stages
func (s *stage) swapTokAndFunc(other *stage) {
	temp := *other

	other.tok = s.tok
	other.evalFunc = s.evalFunc

	s.tok = temp.tok
	s.evalFunc = temp.evalFunc
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
