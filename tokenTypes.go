package eval

// TokenType represents the type of an expression token
type TokenType int

const (
	// VAR is a token that represents a variable. Its value will be the name of the variable.
	VAR TokenType = iota
	// FUNC is a token that represents a function. Its value will be the name of the function.
	FUNC
	// NUM is a token that represents a number. Its value will be the number in float64 format.
	NUM
	// BOOL is a token that represents a bool literal. Its value will be the value of the bool.
	BOOL
	// LEFT is a token that represents a ([{ symbol. It has no value.
	LEFT
	// RIGHT is a token that represents a )]} symbol. It has no value.
	RIGHT
	// ADD is a token that represents a + symbol. It has no value.
	ADD
	// SUB is a token that represents a - symbol. It has no value.
	SUB
	// MUL is a token that represents a * symbol. It has no value.
	MUL
	// DIV is a token that represents a / symbol. It has no value.
	DIV
	// MOD is a token that represents a % symbol. It has no value.
	MOD
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
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case MOD:
		return "MOD"
	}
	return "UNKNOWN"
}
