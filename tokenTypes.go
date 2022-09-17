package eval

// TokenType represents the type of an expression token
type TokenType int

// TokenType Constants
const (
	// Value is type Var
	VAR TokenType = iota
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
	// Value is type NumOp
	NUMOP
	// Value is type Comp
	COMP
	// Value is type LogOp
	LOGOP
)

// NumOp represents the value of a NUMOP token
type NumOp int

// NumOp Constants
const (
	// +
	ADD NumOp = iota
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

// Comp represents the value of a COMP token
type Comp int

// Comp Constants
const (
	// =
	EQUAL Comp = iota
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

// LogOp represents the value of a LOGOP token
type LogOp int

// LogOp Constants
const (
	// &
	AND LogOp = iota
	// |
	OR
)

func (t TokenType) String() string {
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
	case NUMOP:
		return "NUMOP"
	case COMP:
		return "COMP"
	case LOGOP:
		return "LOGOP"
	}
	return "UNKNOWN"
}

func (n NumOp) String() string {
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

func (c Comp) String() string {
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

func (l LogOp) String() string {
	switch l {
	case AND:
		return "AND"
	case OR:
		return "OR"
	}
	return "UNKNOWN"
}
