package eval

// A Func is a function that can be used in expressions. Functions take any number of arguments of any type, and return one argument of any type and an error. Use [eval/fun] to easily make functions.
type Func func(args ...any) (any, error)

// Funcs is a group of named functions that can be used when evaluating expressions.
type Funcs map[string]Func

// NewFuncs makes a new empty set of funcs
func NewFuncs() Funcs {
	return make(Funcs)
}

// NewFuncsFrom makes a new set of functions by copying the other functions
func NewFuncsFrom(from Funcs) Funcs {
	funcs := NewFuncs()
	funcs.Copy(from)
	return funcs
}

// Set sets the value of the given name to the given value.
func (f Funcs) Set(name string, value Func) {
	f[name] = value
}

// Delete removes the function with the given name.
func (f Funcs) Delete(name string) {
	delete(f, name)
}

// Reset resets the functions to an empty map.
func (f Funcs) Reset() {
	for k := range f {
		delete(f, k)
	}
}

// Copy copies the given functions to v. It does not remove the existing functions in v, but it will override any functions with the same name as copied functions.
func (f Funcs) Copy(other Funcs) {
	for k, d := range other {
		f[k] = d
	}
}

// SetTo sets the functions to the other given functions.
func (f Funcs) SetTo(other Funcs) {
	f.Reset()
	f.Copy(other)
}
