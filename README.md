# eval

[![Go Report Card](https://goreportcard.com/badge/github.com/eqgo/eval)](https://goreportcard.com/report/github.com/eqgo/eval)
[![Go Reference](https://pkg.go.dev/badge/github.com/eqgo/eval)](https://pkg.go.dev/github.com/eqgo/eval)

Package eval provides numerical and boolean expression evaluation in Go. It is still a work in progress. Eval is based off of [Knetic/govaluate](https://github.com/Knetic/govaluate) and it has some improvements like supporting expressions that lack multiplication symbols and parentheses.

## Example Usage

    package main

    import (
        "fmt"

        "github.com/eqgo/eval"
    )

    func main() {
        ctx := eval.DefaultContext()
        ctx.Set("x", 3.0)

        expr := eval.New("xcospi")
        err := expr.Compile(ctx)
        if err != nil {
            panic(err)
        }
        res, err := expr.Eval(ctx)
        if err != nil {
            panic(err)
        }
        fmt.Println(res) // -3
    }