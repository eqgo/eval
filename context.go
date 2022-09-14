package eval

// Ctx is context that is given when compiling and evaluating expressions. Ctx contains variables and functions that can be used in expressions.
type Ctx struct {
	Vars  Vars
	Funcs Funcs
}
