package eval

import (
	"github.com/eqgo/mat"
)

// MathContext returns a context with the standard math functions and constants
func MathContext() *Context {
	return &Context{
		Vars:  MathConsts(),
		Funcs: MathFuncs(),
	}
}

// MathFuncs returns a set of functions with all of the standard math functions
func MathFuncs() Funcs {
	return Funcs{
		"sin":     NewFunc1(mat.Sin),
		"cos":     NewFunc1(mat.Cos),
		"tan":     NewFunc1(mat.Tan),
		"sec":     NewFunc1(mat.Sec),
		"csc":     NewFunc1(mat.Csc),
		"cot":     NewFunc1(mat.Cot),
		"arcsin":  NewFunc1(mat.Arcsin),
		"arccos":  NewFunc1(mat.Arccos),
		"arctan":  NewFunc1(mat.Arctan),
		"arcsec":  NewFunc1(mat.Arcsec),
		"arccsc":  NewFunc1(mat.Arccsc),
		"arccot":  NewFunc1(mat.Arccot),
		"sinh":    NewFunc1(mat.Sinh),
		"cosh":    NewFunc1(mat.Cosh),
		"tanh":    NewFunc1(mat.Tanh),
		"sech":    NewFunc1(mat.Sech),
		"csch":    NewFunc1(mat.Csch),
		"coth":    NewFunc1(mat.Coth),
		"arcsinh": NewFunc1(mat.Arcsinh),
		"arccosh": NewFunc1(mat.Arccosh),
		"arctanh": NewFunc1(mat.Arctanh),
		"arcsech": NewFunc1(mat.Arcsech),
		"arccsch": NewFunc1(mat.Arccsch),
		"arccoth": NewFunc1(mat.Arccoth),
		"ln":      NewFunc1(mat.Ln),
		"log":     NewFunc2(mat.Log),
		"abs":     NewFunc1(mat.Abs[float64]),
		"pow":     NewFunc2(mat.Pow),
		"exp":     NewFunc1(mat.Exp),
		"mod":     NewFunc2(mat.Mod),
		"fact":    NewFunc1(mat.Fact),
		"floor":   NewFunc1(mat.Floor),
		"ceil":    NewFunc1(mat.Ceil),
		"round":   NewFunc1(mat.Round),
		"sqrt":    NewFunc1(mat.Sqrt),
		"cbrt":    NewFunc1(mat.Cbrt),
		"min":     NewFuncV(mat.Min[float64]),
		"max":     NewFuncV(mat.Max[float64]),
		"avg":     NewFuncV(mat.Avg[float64]),
	}
}

// MathConsts returns a set of variables with all of the standard math constants
func MathConsts() Vars {
	return Vars{
		"pi": mat.Pi,
		"π":  mat.Pi,
		"e":  mat.E,
	}
}

// LogicalContext returns a context with basic logical functions
func LogicalContext() *Context {
	return &Context{
		Funcs: LogicalFuncs(),
		Vars:  NewVars(),
	}
}

// LogicalFuncs returns a set of basic logical functions
func LogicalFuncs() Funcs {
	return Funcs{
		"if": NewFunc3(func(condition bool, val1, val2 any) any {
			if condition {
				return val1
			}
			return val2
		}),
	}
}

// DefaultContext returns a context with standard math constants, standard math functions, and basic logic functions
func DefaultContext() *Context {
	return &Context{
		Vars:  DefaultVars(),
		Funcs: DefaultFuncs(),
	}
}

// DefaultVars returns a set of variables containing all of the standard math constants
func DefaultVars() Vars {
	return NewVarsFrom(MathConsts())
}

// DefaultFuncs returns a set of functions containing all of the standard math functions, and basic logic functions
func DefaultFuncs() Funcs {
	funcs := NewFuncsFrom(MathFuncs())
	funcs.Copy(LogicalFuncs())
	return funcs
}
