package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	ctx := NewContextFrom(DefaultContext())
	ctx.Set("x", 2.0)
	ctx.Set("a", 3.0)
	ctx.Set("t", 4.0)
	ctx.Set("y", 5.0)
	ex := New("if(x≥2,y,a)+sinx")
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
