package eval

import (
	"fmt"
	"testing"
)

func TestStages(t *testing.T) {
	ctx := NewContextFrom(DefaultContext())
	ctx.Set("x", 2.0)
	ctx.Set("a", 3.0)
	ctx.Set("t", 4.0)
	ctx.Set("y", 5.0)
	type stagesTest struct {
		expr []Token
		want *stage
	}
	testTokens := make([]([]Token), 1)
	testTokens[0], _ = Tokens("sinx", ctx)
	tests := []stagesTest{
		{testTokens[0], &stage{
			evalFunc: funcStage("sin"),
			right: &stage{
				evalFunc: varStage("x"),
			},
		}},
	}
	for i, test := range tests {
		res, err := stages(test.expr)
		if res.String() != test.want.String() {
			t.Errorf("Stages(%v) \nShould be: \n%v \nNot: \n%v with error: %v", test.expr, test.want, res, err)
		} else {
			fmt.Printf("Stages test %d passed with error: %v \n", i, err)
		}
	}
}
