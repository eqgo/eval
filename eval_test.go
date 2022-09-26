package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", 2.0)
	ctx.Set("a", 3.0)
	ctx.Set("t", 4.0)
	ctx.Set("y", 5.0)
	ex := New("-true")
	err := ex.Compile(ctx)
	if err != nil {
		t.Error(err)
	}
	val, err := ex.Eval(ctx)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(val)
}
