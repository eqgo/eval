package eval

// A Var is a variable that can be used in expressions. Variables can be of any type and hold any value.
type Var any

// Vars is a group of named variables that can be used in expressions.
type Vars map[string]Var

// NewVars makes a new empty set of vars
func NewVars() Vars {
	return make(Vars)
}

// NewVarsFrom makes a new set of variables by copying the other variables
func NewVarsFrom(from Vars) Vars {
	vars := NewVars()
	vars.Copy(from)
	return vars
}

// Set sets the value of the given name to the given value.
func (v Vars) Set(name string, value Var) {
	v[name] = value
}

// Delete removes the variable with the given name.
func (v Vars) Delete(name string) {
	delete(v, name)
}

// Reset resets the variables to an empty map.
func (v Vars) Reset() {
	for k := range v {
		delete(v, k)
	}
}

// Copy copies the given variables to v. It does not remove the existing variables in v, but it will override any variables with the same name as copied variables.
func (v Vars) Copy(other Vars) {
	for k, d := range other {
		v[k] = d
	}
}

// SetTo sets the variables to the other given variables.
func (v Vars) SetTo(other Vars) {
	v.Reset()
	v.Copy(other)
}
