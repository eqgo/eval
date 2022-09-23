package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", 0.0)
	ex := New("1+3*4^2")
	ex.Compile(ctx)
	fmt.Println(ex.Eval(ctx))

}
