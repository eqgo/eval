package eval

// Context is context that is given when compiling and evaluating expressions. Context contains variables and functions that can be used in expressions.
type Context struct {
	Vars  Vars
	Funcs Funcs
}

// NewContext makes a new context
func NewContext() *Context {
	return &Context{Vars: NewVars(), Funcs: NewFuncs()}
}

// NewContextFrom makes a new context by copying the other context
func NewContextFrom(from *Context) *Context {
	ctx := NewContext()
	ctx.Copy(from)
	return ctx
}

// Set calls c.Funcs.Set if value is type Func, otherwise it calls c.Vars.Set
func (c *Context) Set(name string, value any) {
	switch value := value.(type) {
	case Func:
		c.SetFunc(name, value)
	default:
		c.SetVar(name, value)
	}
}

// SetVar sets the value of the given variable to the given value
func (c *Context) SetVar(name string, value Var) {
	c.Vars.Set(name, value)
}

// SetFunc sets the value of the given function to the given value
func (c *Context) SetFunc(name string, value Func) {
	c.Funcs.Set(name, value)
}

// Delete deletes the given name from both the vars and funcs
func (c *Context) Delete(name string) {
	c.DeleteVar(name)
	c.DeleteFunc(name)
}

// DeleteVar deletes the given variable
func (c *Context) DeleteVar(name string) {
	delete(c.Vars, name)
}

// DeleteFunc deletes the given function
func (c *Context) DeleteFunc(name string) {
	delete(c.Funcs, name)
}

// Reset resets the variables and functions
func (c *Context) Reset() {
	c.ResetVars()
	c.ResetFuncs()
}

// ResetVars resets the variables
func (c *Context) ResetVars() {
	c.Vars.Reset()
}

// ResetFuncs resets the functions
func (c *Context) ResetFuncs() {
	c.Funcs.Reset()
}

// Copy copies the variables and functions of the other context and does not delete existing variables and functions
func (c *Context) Copy(other *Context) {
	c.CopyVars(other)
	c.CopyFuncs(other)
}

// CopyVars copies the vars from the other context and does not delete existing variables
func (c *Context) CopyVars(other *Context) {
	c.Vars.Copy(other.Vars)
}

// CopyFuncs copies the funcs from the other context and does not delete existing functions
func (c *Context) CopyFuncs(other *Context) {
	c.Funcs.Copy(other.Funcs)
}

// SetTo sets the context to the other context
func (c *Context) SetTo(other *Context) {
	c.SetVarsTo(other)
	c.SetFuncsTo(other)
}

// SetVarsTo sets the vars of the context to the vars of the other context
func (c *Context) SetVarsTo(other *Context) {
	c.ResetVars()
	c.CopyVars(other)
}

// SetFuncsTo sets the funcs of the context to the funcs of the other context
func (c *Context) SetFuncsTo(other *Context) {
	c.ResetFuncs()
	c.CopyFuncs(other)
}
