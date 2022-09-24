package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", 5.32)
	expr := "xarccoth78.9"
	ex := New(expr)
	ex.Compile(ctx)
	fmt.Println(ex.Tokens)
	fmt.Println(ex.Eval(ctx))

}
