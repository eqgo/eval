package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", 2.0)
	expr := "xsinx"
	ex := New(expr)
	ex.Compile(ctx)
	fmt.Println(ex.Tokens)
	fmt.Println(ex.Eval(ctx))

}
