package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContext()
	ctx.Set("x", 0)
	ctx.Set("sin", NewFunc1(math.Sin))
	fmt.Println(Tokens("sin(x%2)", ctx))
}
