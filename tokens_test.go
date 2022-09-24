package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", -5.32)
	expr := "xarccotx+min(x,2)-max(3, x)+sinhpi"
	ex := New(expr)
	ex.Compile(ctx)
	fmt.Println(ex.Tokens)
	fmt.Println(ex.Eval(ctx))

}
