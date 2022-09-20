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
	lt, lv, err := evalBinary(context, expr.X)
	if err != nil {
		return 0, "", err
	}
	rt, rv, err := evalBinary(context, expr.Y)
	if err != nil {
		return 0, "", err
	}
	if isLikeNumber(lt, lv) && isLikeNumber(rt, rv) {
		l := getValue(lv)
		r := getValue(rv)
		switch expr.Op {
		case token.ADD:
			{
				return token.FLOAT, fmt.Sprintf("%f", l+r), nil
			}
		case token.SUB:
			{
				return token.FLOAT, fmt.Sprintf("%f", l-r), nil
			}
		case token.MUL:
			{
				return token.FLOAT, fmt.Sprintf("%f", l*r), nil
			}
		case token.QUO:
			{
				return token.FLOAT, fmt.Sprintf("%f", l/r), nil
			}
		case token.REM:
			{
				return token.FLOAT, fmt.Sprintf("%f", math.Remainder(l, r)), nil
			}
		case token.LAND:
			fallthrough
		case token.AND:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)&int(r)), nil
			}
		case token.LOR:
			fallthrough
		case token.OR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)|int(r)), nil
			}
		case token.XOR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)^int(r)), nil
			}
		case token.SHL:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)<<int(r)), nil
			}
		case token.SHR:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)>>int(r)), nil
			}
		case token.AND_NOT:
			{
				return token.FLOAT, fmt.Sprintf("%d", int(l)&^int(r)), nil
			}
		case token.EQL:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(l == r)), nil
			}
		case token.NEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(l != r)), nil
			}
		case token.GTR:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(l > r)), nil
			}
		case token.GEQ:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(l >= r)), nil
			}
		case token.LSS:
			{
				return token.FLOAT, fmt.Sprintf("%d", getBoolean(l < r)), nil
			}
		case token.LEQ:
			{
				return token.IDENT, fmt.Sprintf("%d", getBoolean(l < r)), nil
			}
		}
		return 0, "", errors.New("invalid data")
	}
	if lt == token.STRING && rt == token.STRING {
		switch expr.Op {
		case token.EQL:
			{
				return token.IDENT, fmt.Sprintf("%t", lv == rv), nil
			}
		case token.NEQ:
			{
				return token.IDENT, fmt.Sprintf("%t", lv != rv), nil
			}
		case token.ADD:
			{
				return token.STRING, lv + rv, nil
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
	case "0":
		fallthrough
	case "1":
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
