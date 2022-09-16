package eval

import (
	"fmt"
)

// NewFunc1 makes a function that can be used in expressions from a function that takes a single argument and returns a single value.
func NewFunc1[I, O any](f func(I) O) Func {
	return func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("evaluation error: function of type %T wants 1 argument, not %v arguments", f, len(args))
		}
		arg0, ok := args[0].(I)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T", f, args[0])
		}
		res := f(arg0)
		return res, nil
	}
}
