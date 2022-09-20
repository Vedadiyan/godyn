package godyn

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"strconv"
)

func evalBinary(context *Context, expr ast.Expr) (token.Token, string, error) {
	switch t := expr.(type) {
	case *ast.BinaryExpr:
		{
			return readBinaryExpression(context, t)
		}
	case *ast.CallExpr:
		{
			value, err := context.eval(t)
			if err != nil {
				return 0, "", err
			}
			return token.OR, fmt.Sprintf("%v", value), nil
		}
	case *ast.ParenExpr:
		{
			return evalBinary(context, t.X)
		}
	case *ast.BasicLit:
		{
			return t.Kind, t.Value, nil
		}
	case *ast.Ident:
		{
			return token.IDENT, t.Name, nil
		}
	}
	return 0, "", errors.New("invalid expression")
}

func readBinaryExpression(context *Context, expr *ast.BinaryExpr) (token.Token, string, error) {
	left_token, left_value, err := evalBinary(context, expr.X)
	if err != nil {
		return 0, "", err
	}
	right_token, right_value, err := evalBinary(context, expr.Y)
	if err != nil {
		return 0, "", err
	}
	if isLikeNumber(left_token, left_value) && isLikeNumber(right_token, right_value) {
		left := getValue(left_value)
		right := getValue(right_value)
		switch expr.Op {
		case token.ADD:
			{
				return token.FLOAT, fmt.Sprintf("%f", left+right), nil
			}
		case token.SUB:
			{
				return token.FLOAT, fmt.Sprintf("%f", left-right), nil
			}
		case token.MUL:
			{
				return token.FLOAT, fmt.Sprintf("%f", left*right), nil
			}
		case token.QUO:
			{
				return token.FLOAT, fmt.Sprintf("%f", left/right), nil
			}
		case token.REM:
			{
				return token.FLOAT, fmt.Sprintf("%f", math.Remainder(left, right)), nil
			}
		case token.LAND:
			fallthrough
		case token.AND:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)&int(right)), nil
			}
		case token.LOR:
			fallthrough
		case token.OR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)|int(right)), nil
			}
		case token.XOR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)^int(right)), nil
			}
		case token.SHL:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)<<int(right)), nil
			}
		case token.SHR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)>>int(right)), nil
			}
		case token.AND_NOT:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(left)&^int(right)), nil
			}
		case token.EQL:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left == right)), nil
			}
		case token.NEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left != right)), nil
			}
		case token.GTR:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left > right)), nil
			}
		case token.GEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left >= right)), nil
			}
		case token.LSS:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(left < right)), nil
			}
		case token.LEQ:
			{
				return token.IDENT, fmt.Sprintf("%d", getBoolean(left < right)), nil
			}
		}
		return 0, "", errors.New("invalid data")
	}
	if left_token == token.STRING && right_token == token.STRING {
		switch expr.Op {
		case token.EQL:
			{
				return token.IDENT, fmt.Sprintf("%t", left_value == right_value), nil
			}
		case token.NEQ:
			{
				return token.IDENT, fmt.Sprintf("%t", left_value != right_value), nil
			}
		case token.ADD:
			{
				return token.STRING, left_value + right_value, nil
			}
		}
	}
	return 0, "", errors.New("invalid operation")
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
