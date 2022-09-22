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
	// ,
	SEP
	// Value is type NumOp
	NUMOP
	// Value is type Comp
	COMP
	// Value is type LogOp
	LOGOP
	// Value is type SLogOp
	SLOGOP
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

// SLogOp represents the value of a SLOGOP token
type SLogOp int

// SLogOp Constants
const (
	// !
	NOT SLogOp = iota
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
	case SEP:
		return "SEP"
	case NUMOP:
		return "NUMOP"
	case COMP:
		return "COMP"
	case LOGOP:
		return "LOGOP"
	case SLOGOP:
		return "SLOGOP"
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

func (sl SLogOp) String() string {
	switch sl {
	case NOT:
		return "NOT"
	}
	return "UNKNOWN"
}
