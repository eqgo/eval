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
	ctx.Set("cos", NewFunc1(math.Cos))
	fmt.Println(Tokens("sin(x)>0 & cos(x) <1 | y +3.2 ", ctx))
}
