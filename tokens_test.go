package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Copy(MathContext)
	ctx.Set("x", 0)
	ex := New("-3.1415sinx")
	err := ex.Compile(ctx)
	fmt.Println(ex.Tokens, err)
}
