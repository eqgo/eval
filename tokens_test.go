package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", -5.32)
	expr := "-x"
	ex := New(expr)
	err := ex.Compile(ctx)
	fmt.Println(err)
	fmt.Println(ex.tokens)
	fmt.Println(ex.Eval(ctx))

}
