package eval

import (
	"fmt"
)

// NewFuncV makes a function that can be used in expressions from a function that takes a variadic input and returns a single value.
func NewFuncV[I, O any](f func(...I) O) Func {
	return func(args ...any) (any, error) {
		newArgs := []I{}
		for i, arg := range args {
			a, ok := arg.(I)
			if !ok {
				return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument %v", f, arg, i)
			}
			newArgs = append(newArgs, a)
		}
		res := f(newArgs...)
		return res, nil
	}
}

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

// NewFunc2 makes a function that can be used in expressions from a function that takes two arguments and returns a single value.
func NewFunc2[I1, I2, O any](f func(I1, I2) O) Func {
	return func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("evaluation error: function of type %T wants 2 arguments, not %v arguments", f, len(args))
		}
		arg0, ok := args[0].(I1)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument 0", f, args[0])
		}
		arg1, ok := args[1].(I2)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument 1", f, args[1])
		}
		res := f(arg0, arg1)
		return res, nil
	}
}

// NewFunc3 makes a function that can be used in expressions from a function that takes three arguments and returns a single value.
func NewFunc3[I1, I2, I3, O any](f func(I1, I2, I3) O) Func {
	return func(args ...any) (any, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("evaluation error: function of type %T wants 3 arguments, not %v arguments", f, len(args))
		}
		arg0, ok := args[0].(I1)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument 0", f, args[0])
		}
		arg1, ok := args[1].(I2)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument 1", f, args[1])
		}
		arg2, ok := args[2].(I3)
		if !ok {
			return nil, fmt.Errorf("evaluation error: function of type %T does not accept input type %T for argument 2", f, args[2])
		}
		res := f(arg0, arg1, arg2)
		return res, nil
	}
}
