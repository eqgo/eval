package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	ctx := NewContextFrom(DefaultContext())
	ctx.Set("x", 2.0)
	ctx.Set("a", 3.0)
	ctx.Set("t", 4.0)
	ctx.Set("y", 5.0)
	type tokensTest struct {
		expr string
		want []Token
	}
	tests := []tokensTest{
		{"sinx", []Token{{FUNC, "sin"}, {LEFT, nil}, {VAR, "x"}, {RIGHT, nil}}},
		{"", nil},
		{"y=max(2sinx, 3cosx)",
			[]Token{{VAR, "y"}, {COMP, EQUAL}, {FUNC, "max"}, {LEFT, nil}, {NUM, 2}, {NUMOP, MUL}, {FUNC, "sin"}, {LEFT, nil}, {VAR, "x"},
				{RIGHT, nil}, {SEP, nil}, {NUM, 3}, {NUMOP, MUL}, {FUNC, "cos"}, {LEFT, nil}, {VAR, "x"}, {RIGHT, nil}, {RIGHT, nil}}},
		{"42.873pieaxtyarccothx3+4x^2-7y^3+min(piea9.3t,y6.82tx)",
			[]Token{{NUM, 42.873}, {NUMOP, MUL}, {VAR, "pi"}, {NUMOP, MUL}, {VAR, "e"}, {NUMOP, MUL}, {VAR, "a"}, {NUMOP, MUL},
				{VAR, "x"}, {NUMOP, MUL}, {VAR, "t"}, {NUMOP, MUL}, {VAR, "y"}, {NUMOP, MUL}, {FUNC, "arccoth"}, {LEFT, nil},
				{VAR, "x"}, {RIGHT, nil}, {NUMOP, MUL}, {NUM, 3}, {NUMOP, ADD}, {NUM, 4}, {NUMOP, MUL}, {VAR, "x"}, {NUMOP, POW},
				{NUM, 2}, {NUMOP, SUB}, {NUM, 7}, {NUMOP, MUL}, {VAR, "y"}, {NUMOP, POW}, {NUM, 3}, {NUMOP, ADD}, {FUNC, "min"},
				{LEFT, nil}, {VAR, "pi"}, {NUMOP, MUL}, {VAR, "e"}, {NUMOP, MUL}, {VAR, "a"}, {NUMOP, MUL}, {NUM, 9.3}, {NUMOP, MUL},
				{VAR, "t"}, {SEP, nil}, {VAR, "y"}, {NUMOP, MUL}, {NUM, 6.82}, {NUMOP, MUL}, {VAR, "t"}, {NUMOP, MUL}, {VAR, "x"},
				{RIGHT, nil}}},
	}
	for _, test := range tests {
		res, err := Tokens(test.expr, ctx)
		if fmt.Sprint(res) != fmt.Sprint(test.want) {
			t.Errorf("Tokens(%v) \nShould be: \n%v \nNot: \n%v with error: %v", test.expr, test.want, res, err)
		}
	}
	fmt.Printf("%d tokens tests passed \n", len(tests))
}
