package eval

import "math"

// stageEval represents a function that can be used as the eval function for a stage
type stageEval func(left, right any, ctx *Context) (any, error)

var priorityTokensMap = map[priority][]Token{
	priorityAdd: {{NUMOP, ADD}, {NUMOP, SUB}},
	priorityMul: {{NUMOP, MUL}, {NUMOP, DIV}},
	priorityPow: {{NUMOP, POW}},
}

var tokenStageEvalMap = map[Token]stageEval{
	{NUMOP, SUB}: subStage,
	{NUMOP, ADD}: addStage,
	{NUMOP, DIV}: divStage,
	{NUMOP, MUL}: mulStage,
	{NUMOP, POW}: powStage,
}

// litStage makes the stageEval for a stage that is a literal value from a token
func litStage(t Token) stageEval {
	switch t.Type {
	case NUM, BOOL:
		return func(left, right any, ctx *Context) (any, error) {
			return t.Value, nil
		}
	case VAR:
		return func(left, right any, ctx *Context) (any, error) {
			return ctx.Vars[t.Value.(string)], nil
		}
	}
	return nil

}

// addStage is the stageEval for a stage that is an add op
func addStage(left, right any, ctx *Context) (any, error) {
	return left.(float64) + right.(float64), nil
}

// subStage is the stageEval for a stage that is an sub op
func subStage(left, right any, ctx *Context) (any, error) {
	return left.(float64) - right.(float64), nil
}

// mulStage is the stageEval for a stage that is an mul op
func mulStage(left, right any, ctx *Context) (any, error) {
	return left.(float64) * right.(float64), nil
}

// divStage is the stageEval for a stage that is an div op
func divStage(left, right any, ctx *Context) (any, error) {
	return left.(float64) / right.(float64), nil
}

// powStage is the stageEval for a stage that is an pow op
func powStage(left, right any, ctx *Context) (any, error) {
	return math.Pow(left.(float64), right.(float64)), nil
}
