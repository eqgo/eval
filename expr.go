// Package eval provides mathematical expression evaluation.
package eval

// Expr is an expression that can be evaluated
type Expr struct {
	Expr   string
	Tokens []Token
	stages *stage
}

// NewExpr makes a new expression from the given expression string
func NewExpr(expr string) *Expr {
	return &Expr{Expr: expr}
}

// New is an alias for NewExpr
func New(expr string) *Expr {
	return NewExpr(expr)
}

// Set sets the expression of the expression
func (ex *Expr) Set(expr string) {
	ex.Expr = expr
}

// Compile compiles the expression with the given context
func (ex *Expr) Compile(ctx *Context) error {
	t, err := Tokens(ex.Expr, ctx)
	if err != nil {
		return err
	}
	ex.Tokens = t
	s, err := stages(ex.Tokens)
	if err != nil {
		return err
	}
	ex.stages = s
	return nil
}

// Eval evaluates the expression with the given context
func (ex *Expr) Eval(ctx *Context) (any, error) {
	if ex.stages == nil {
		return nil, nil
	}
	return ex.stages.eval(ctx)
}
