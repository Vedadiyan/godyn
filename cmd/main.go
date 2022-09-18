package main

import (
	"go/ast"
	"go/parser"

	godyn "github.com/vedadiyan/godyn/pkg"
)

const TEST = `fn((1+2)*3*(8*(8+3*(2+2)))+10)`

// const TEST = `fn(value_a && false)`

func main() {
	expr, err := parser.ParseExpr(TEST)
	if err != nil {
		panic(err)
	}
	val, e := godyn.ReadExpression(expr.(*ast.CallExpr).Args[0])
	_ = val
	_ = e
	_ = expr
}
