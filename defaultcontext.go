package eval

import (
	"github.com/eqgo/mat"
)

// MathContext is a context with the standard math functions and constants
var MathContext = &Context{
	Vars:  MathConsts,
	Funcs: MathFuncs,
}

// MathFuncs are all of the standard math functions
var MathFuncs = Funcs{
	"sin": NewFunc1(mat.Sin),
	"cos": NewFunc1(mat.Cos),
	"tan": NewFunc1(mat.Tan),
}

// MathConsts are all of the standard math constants
var MathConsts = Vars{
	"pi": mat.Pi,
	"π":  mat.Pi,
	"e":  mat.E,
}
