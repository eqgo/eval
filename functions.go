package eval

// A Func is a function that can be used in expressions. Functions take any number of arguments of any type, and return one argument of any type and an error. Use [eval/fun] to easily make functions.
type Func func(args ...any) (any, error)

// Funcs is a group of named functions that can be used when evaluating expressions.
type Funcs map[string]Func
