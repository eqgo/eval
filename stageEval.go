package eval

import (
	"fmt"
	"math"
)

// stageEval represents a function that can be used as the eval function for a stage
type stageEval func(left, right any, ctx *Context) (any, error)

var tokenStageEvalMap = map[Token]stageEval{
	{SEP, nil}:   sepStage,
	{NUMOP, SUB}: subStage,
	{NUMOP, ADD}: addStage,
	{NUMOP, DIV}: divStage,
	{NUMOP, MUL}: mulStage,
	{NUMOP, MOD}: modStage,
	{NUMOP, POW}: powStage,
}

// litStage makes the stageEval for a stage that is a literal value from a token
func litStage(t Token) stageEval {
	return func(left, right any, ctx *Context) (any, error) {
		return t.Value, nil
	}

}

// varStage makes the stageEval for a stage that it is a variable
func varStage(name string) stageEval {
	return func(left, right any, ctx *Context) (any, error) {
		v, ok := ctx.Vars[name]
		if !ok {
			return nil, fmt.Errorf("var %v is not defined", name)
		}
		return v, nil
	}
}

func functionStage(name string) stageEval {
	return func(left, right any, ctx *Context) (any, error) {
		f, ok := ctx.Funcs[name]
		if !ok {
			return nil, fmt.Errorf("func %v is not defined", name)
		}
		switch right.(type) {
		case []any:
			return f(right.([]any)...)
		default:
			return f(right)
		}
	}
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

// modStage is the stageEval for a stage that is an mod op
func modStage(left, right any, ctx *Context) (any, error) {
	return math.Mod(left.(float64), right.(float64)), nil
}

// powStage is the stageEval for a stage that is an pow op
func powStage(left, right any, ctx *Context) (any, error) {
	return math.Pow(left.(float64), right.(float64)), nil
}

// sepStage is the stageEval for a stage that is a sep op
func sepStage(left, right any, ctx *Context) (any, error) {
	res := []any{}
	switch left.(type) {
	case []any:
		res = append(res, left.([]any)...)
	default:
		res = append(res, left)
	}
	switch right.(type) {
	case []any:
		res = append(res, right.([]any)...)
	default:
		res = append(res, right)
	}
	return res, nil
}
