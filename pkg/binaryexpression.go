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
	// if left_token != right_token {
	// 	panic("inavlid operation")
	// }
	if isLikeNumber(left_token, left_value) && isLikeNumber(right_token, right_value) {
		left := getValue(left_value)
		right := getValue(right_value)
		switch expr.Op {
		case token.ADD:
			{
				return token.FLOAT, fmt.Sprintf("%f", left+right)
			}
		case token.SUB:
			{
				return token.FLOAT, fmt.Sprintf("%f", left-right)
			}
		case token.MUL:
			{
				return token.FLOAT, fmt.Sprintf("%f", left*right)
			}
		case token.QUO:
			{
				return token.FLOAT, fmt.Sprintf("%f", left/right)
			}
		case token.REM:
			{
				return token.FLOAT, fmt.Sprintf("%f", math.Remainder(left, right))
			}
		case token.LAND:
			fallthrough
		case token.AND:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)&int(right))
			}
		case token.LOR:
			fallthrough
		case token.OR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)|int(right))
			}
		case token.XOR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)^int(right))
			}
		case token.SHL:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)<<int(right))
			}
		case token.SHR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)>>int(right))
			}
		case token.AND_NOT:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)&^int(right))
			}
		case token.EQL:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left == right))
			}
		case token.NEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left != right))
			}
		case token.GTR:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left > right))
			}
		case token.GEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left >= right))
			}
		case token.LSS:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left < right))
			}
		case token.LEQ:
			{
				return token.IDENT, fmt.Sprintf("%d", getBoolean(left < right))
			}
		}
		panic("Invalid Data")
	}
	// if left_token == token.IDENT {
	// 	// if isBoolean(left_value) && isBoolean(right_value) {
	// 	// 	left, err := strconv.ParseBool(left_value)
	// 	// 	if err != nil {
	// 	// 		panic("")
	// 	// 	}
	// 	// 	right, err := strconv.ParseBool(right_value)
	// 	// 	if err != nil {
	// 	// 		panic("")
	// 	// 	}
	// 	// 	switch expr.Op {
	// 	// 	case token.LAND:
	// 	// 		{
	// 	// 			return token.FLOAT, fmt.Sprintf("%d", getBoolean(left && right))
	// 	// 		}
	// 	// 	case token.LOR:
	// 	// 		{
	// 	// 			return token.FLOAT, fmt.Sprintf("%d", getBoolean(left || right))
	// 	// 		}
	// 	// 	}
	// 	// }
	// 	// v := 1
	// 	// _ = v
	// 	// panic("")
	// }
	if left_token == token.STRING {
		switch expr.Op {
		case token.EQL:
			{
				return token.IDENT, fmt.Sprintf("%t", left_value == right_value)
			}
		case token.NEQ:
			{
				return token.IDENT, fmt.Sprintf("%t", left_value != right_value)
			}
		case token.ADD:
			{
				return token.STRING, left_value + right_value
			}
		}
	}
	_ = left_token
	_ = left_value
	_ = right_token
	_ = right_value
	return expr.Op, "XX"
}

func isLikeNumber(kind token.Token, value string) bool {
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
	return isBoolean(value)
}
func isBoolean(value string) bool {
	switch value {
	case "1":
		fallthrough
	case "2":
		fallthrough
	case "false":
		fallthrough
	case "true":
		{
			return true
		}
	}
	return false
}

func getBoolean(b bool) int {
	if b {
		return 1
	}
	return 0
}

func getValue(value string) float64 {
	if isBoolean(value) {
		v, err := strconv.ParseBool(value)
		if err != nil {
			panic(err)
		}
		return float64(getBoolean(v))
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return v
}
