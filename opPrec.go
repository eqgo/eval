package eval

// opPrec represents the precedence of an operator
type opPrec int

const (
	valPrec opPrec = iota
	funcPrec
	prePrec
	powPrec
	addPrec
	mulPrec
	compPrec
	andPrec
	orPrec
	sepPrec
)

// opPrec returns the operator precedence of the token
func (t Token) opPrec() opPrec {
	switch t.Type {
	case NUM, BOOL:
		return valPrec
	case FUNC:
		return funcPrec
	case NUMPRE, LOGPRE:
		return prePrec
	case NUMOP:
		switch t.Value {
		case POW:
			return powPrec
		case ADD, SUB:
			return addPrec
		case MUL, DIV, MOD:
			return mulPrec
		}
	case COMP:
		return compPrec
	case SEP:
		return sepPrec
	case LOGOP:
		switch t.Value {
		case AND:
			return andPrec
		case OR:
			return orPrec
		}
	}
	return valPrec
}
