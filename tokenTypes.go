package eval

// tokenType represents the type of an expression token
type tokenType int

// tokenType Constants
const (
	// Value is type Var
	VAR tokenType = iota
	// Value is type Func
	FUNC
	// float64
	NUM
	// literal
	BOOL
	// ({[
	LEFT
	// )}]
	RIGHT
	// ,
	SEP
	// Value is type NumOp
	NUMOP
	// Value is type NumPre
	NUMPRE
	// Value is type Comp
	COMP
	// Value is type LogOp
	LOGOP
	// Value is type LogPre
	LOGPRE
)

// numOp represents the value of a NUMOP token
type numOp int

// numOp Constants
const (
	// +
	ADD numOp = iota
	// -
	SUB
	// *
	MUL
	// /
	DIV
	// ^
	POW
	// %
	MOD
)

// numPre represents the value of a NUMPRE token
type numPre int

// numPre Constants
const (
	// -
	NEG numPre = iota
)

// comp represents the value of a COMP token
type comp int

// comp Constants
const (
	// =
	EQUAL comp = iota
	// !=
	NOTEQUAL
	// >
	GREATER
	// <
	LESS
	// >=
	GEQ
	// <=
	LEQ
)

// logOp represents the value of a LOGOP token
type logOp int

// logOp Constants
const (
	// &
	AND logOp = iota
	// |
	OR
)

// logPre represents the value of a LOGPRE token
type logPre int

// logPre Constants
const (
	// !
	NOT logPre = iota
)

func (t tokenType) String() string {
	switch t {
	case VAR:
		return "VAR"
	case FUNC:
		return "FUNC"
	case NUM:
		return "NUM"
	case BOOL:
		return "BOOL"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case SEP:
		return "SEP"
	case NUMOP:
		return "NUMOP"
	case NUMPRE:
		return "NUMPRE"
	case COMP:
		return "COMP"
	case LOGOP:
		return "LOGOP"
	case LOGPRE:
		return "LOGPRE"
	}
	return "UNKNOWN"
}

func (n numOp) String() string {
	switch n {
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case POW:
		return "POW"
	case MOD:
		return "MOD"
	}
	return "UNKNOWN"
}

func (n numPre) String() string {
	switch n {
	case NEG:
		return "NEG"
	}
	return "UNKNOWN"
}

func (c comp) String() string {
	switch c {
	case EQUAL:
		return "EQUAL"
	case NOTEQUAL:
		return "NOTEQUAL"
	case GREATER:
		return "GREATER"
	case LESS:
		return "LESS"
	case GEQ:
		return "GEQ"
	case LEQ:
		return "LEQ"
	}
	return "UNKNOWN"
}

func (l logOp) String() string {
	switch l {
	case AND:
		return "AND"
	case OR:
		return "OR"
	}
	return "UNKNOWN"
}

func (l logPre) String() string {
	switch l {
	case NOT:
		return "NOT"
	}
	return "UNKNOWN"
}
