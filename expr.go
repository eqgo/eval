package eval

// Expr is an expression that can be evaluated
type Expr struct {
	Expr   string
	Tokens []Token
	Stages *Stage
}

// NewExpr makes a new expression from the given expression string
func NewExpr(expr string) *Expr {
	return &Expr{Expr: expr}
}

// New is an alias for NewExpr
func New(expr string) *Expr {
	return NewExpr(expr)
}

// Compile compiles the expression with the given context
func (ex *Expr) Compile(ctx *Context) error {
	t, err := Tokens(ex.Expr, ctx)
	if err != nil {
		return err
	}
	ex.Tokens = t
	s, err := Stages(ex.Tokens)
	if err != nil {
		return err
	}
	ex.Stages = s
	return nil
}
