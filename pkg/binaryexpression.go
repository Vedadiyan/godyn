package godyn

import (
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"strconv"

	"github.com/vedadiyan/gocollections/pkg/queue"
	"github.com/vedadiyan/gocollections/pkg/stack"
)

func ToPostfix(expr *ast.BinaryExpr) {
	params := queue.New[string]()
	_ = params
	tokens := stack.New[token.Token]()
	tokens.Push(token.EOF)
	isOk := 0
	_ = isOk

}

func evalBinary(context *Context, expr ast.Expr) (token.Token, string) {
	switch t := expr.(type) {
	case *ast.BinaryExpr:
		{
			return readBinaryExpression(context, t)
		}
	case *ast.CallExpr:
		{
			value, err := context.eval(t)
			if err != nil {
				panic("")
			}
			return token.OR, fmt.Sprintf("%v", value)
		}
	case *ast.ParenExpr:
		{
			return evalBinary(context, t.X)
		}
	case *ast.BasicLit:
		{
			return t.Kind, t.Value
		}
	case *ast.Ident:
		{
			return token.IDENT, t.Name
		}
	}
	return token.EOF, ""
}

func readBinaryExpression(context *Context, expr *ast.BinaryExpr) (token.Token, string) {
	left_token, left_value := evalBinary(context, expr.X)
	right_token, right_value := evalBinary(context, expr.Y)
	if left_token != right_token {
		panic("inavlid operation")
	}
	if isNumber(left_token) {
		left, err := strconv.ParseFloat(left_value, 64)
		if err != nil {
			panic("")
		}
		right, err := strconv.ParseFloat(right_value, 64)
		if err != nil {
			panic("")
		}
		switch expr.Op {
		case token.ADD:
			{
				return left_token, fmt.Sprintf("%f", left+right)
			}
		case token.SUB:
			{
				return left_token, fmt.Sprintf("%f", left-right)
			}
		case token.MUL:
			{
				return left_token, fmt.Sprintf("%f", left*right)
			}
		case token.QUO:
			{
				return left_token, fmt.Sprintf("%f", left/right)
			}
		case token.REM:
			{
				return left_token, fmt.Sprintf("%f", math.Remainder(left, right))
			}
		case token.AND:
			{
				return left_token, fmt.Sprintf("%d", int(left)&int(right))
			}
		case token.OR:
			{
				return left_token, fmt.Sprintf("%d", int(left)|int(right))
			}
		case token.XOR:
			{
				return left_token, fmt.Sprintf("%d", int(left)^int(right))
			}
		case token.SHL:
			{
				return left_token, fmt.Sprintf("%d", int(left)<<int(right))
			}
		case token.SHR:
			{
				return left_token, fmt.Sprintf("%d", int(left)>>int(right))
			}
		case token.AND_NOT:
			{
				return left_token, fmt.Sprintf("%d", int(left)&^int(right))
			}
		}
	}
	_ = left_token
	_ = left_value
	_ = right_token
	_ = right_value
	return expr.Op, "XX"
}

func isNumber(kind token.Token) bool {
	switch kind {
	case token.INT:
		fallthrough
	case token.FLOAT:
		fallthrough
	case token.IMAG:
		{
			return true
		}
	}
	return false
}
func isBoolean(value string) bool {
	switch value {
	case "false":
		fallthrough
	case "true":
		{
			return true
		}
	}
	return false
}
