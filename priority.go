package eval

// priority is the priority an operator has, the lower the better
type priority int

// priority constants
const (
	// + or -
	priorityAdd priority = iota
	// * or /
	priorityMul
	// %
	priorityMod
	// ^
	priorityPow
	priorityLen
)
