package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	ctx := DefaultContext()
	ctx.Set("x", 2.0)
	ctx.Set("a", 3.0)
	ctx.Set("c", 4.0)
	ctx.Set("y", 5.0)
	ctx.Set("time", 6.0)
	ctx.Set("seven", 7.0)
	ctx.Set("be", -97.4397)
	ctx.Set("t", true)
	ctx.Set("f", false)
	type evalTest struct {
		expr string
		want any
	}
	tests := []evalTest{
		// numbers and numerical operators without variables
		{expr: "7", want: 7.0},
		{expr: "-8.453", want: -8.453},
		{expr: "1+2", want: 3.0},
		{expr: "9-5", want: 4.0},
		{expr: "3*4", want: 12.0},
		{expr: "8/6", want: 8.0 / 6.0},
		{expr: "3-(-4)", want: 7.0},
		{expr: "2(3+5)", want: 16.0},
		{expr: "8/(9-4)", want: 8.0 / 5.0},
		{expr: "-((4%2.5)^2)", want: -(9.0 / 4.0)},
		{expr: "1+2*3.5", want: 8.0},
		{expr: "2.0(3+4)", want: 14.0},
		{expr: "3/4.0/2", want: 0.375},
		{expr: "8*9.2*4*3.1", want: 912.64},
		{expr: "1000*1000*1000*1000/100", want: 10000000000.0},
		{expr: "2^20", want: 1048576.0},
		{expr: "3^-6", want: 1.0 / 729.0},
		{expr: "+3%(125^(1/3))", want: 3.0},
		{expr: "2/3*4", want: 8.0 / 3.0},
		{expr: "1-2+3", want: 2.0},
		// bools and logical operators without variables
		{expr: "true", want: true},
		{expr: "false", want: false},
		{expr: "!false", want: true},
		{expr: "!true", want: false},
		{expr: "false|true", want: true},
		{expr: "true||false", want: true},
		{expr: "false&&true", want: false},
		{expr: "true&true", want: true},
		{expr: "(true&false)||true", want: true},
		{expr: "(false&&true)|(true&false)", want: false},
		{expr: "!(((false&false)|(true|true))&false)", want: true},
		{expr: "(((false||true)&&(!true|false)&&true)|(((true||false)&false)||(!false|false))", want: true},
		// comparison operators without variables
		{expr: "1>3", want: false},
		{expr: "5>4", want: true},
		{expr: "7>7", want: false},
		{expr: "7<7", want: false},
		{expr: "8<11", want: true},
		{expr: "-43.98<-32", want: true},
		{expr: "-4>=-4", want: true},
		{expr: "2<=2", want: true},
		{expr: "4.9999>=5", want: false},
		{expr: "-82≥42", want: false},
		{expr: "-3.492≤-1.287", want: true},
		{expr: "1=2", want: false},
		{expr: "4==4", want: true},
		{expr: "(3+-57.2)==-54.2", want: true},
		{expr: "(3.999^2)=16", want: false},
		{expr: "-3!=+3", want: true},
		// numbers and numerical operators with variables
		{expr: "x", want: 2.0},
		{expr: "3y-c", want: 11.0},
		{expr: "a+time", want: 9.0},
		{expr: "y/a", want: 5.0 / 3.0},
		{expr: "c*a", want: 12.0},
		{expr: "a%x", want: 1.0},
		{expr: "bey", want: -487.1985},
		{expr: "yx^2", want: 20.0},
		{expr: "2c%y", want: 3.0},
		{expr: "bexactime", want: -14031.3168},
		{expr: "pixe/y", want: 3.4158936890694265},
		{expr: "3timetimetime", want: 648.0},
		{expr: "4.5000x-acy/8.2", want: 1.6829268292682924},
		{expr: "3xpiepiepiepie/timetimetime", want: 191460.82200720004},
	}
	for _, test := range tests {
		expr := New(test.expr)
		err := expr.Compile(ctx)
		if err != nil {
			t.Errorf("Eval Test of expression %v failed with a compile error: %v", test.expr, err)
		}
		res, err := expr.Eval(ctx)
		if err != nil {
			t.Errorf("Eval Test of expression %v failed with an evaluation error: %v", test.expr, err)
		}
		if res != test.want {
			t.Errorf("Eval Test of expression: %v failed with expected result: %v and actual result %v", test.expr, test.want, res)
		}
	}
	fmt.Printf("%d evaluation tests passed \n", len(tests))
}
