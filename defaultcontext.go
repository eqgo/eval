package eval

import (
	"math"

	"github.com/eqgo/mat"
)

// MathContext is a context with the standard math functions and constants
var MathContext = &Context{
	Vars:  MathConsts,
	Funcs: MathFuncs,
}

// MathFuncs are all of the standard math functions
var MathFuncs = Funcs{
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
	"arcsech": NewFunc1(mat.Arcsech),
	"arccsch": NewFunc1(mat.Arccsch),
	"arccoth": NewFunc1(mat.Arccoth),
	"ln":      NewFunc1(mat.Ln),
	"log":     NewFunc2(mat.Log),
	"abs":     NewFunc1(math.Abs),
	"pow":     NewFunc2(math.Pow),
	"mod":     NewFunc2(math.Mod),
}

// MathConsts are all of the standard math constants
var MathConsts = Vars{
	"pi": mat.Pi,
	"π":  mat.Pi,
	"e":  mat.E,
}
