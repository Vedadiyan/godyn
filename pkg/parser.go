package godyn

import (
	"bytes"
	"errors"
	"go/ast"
	"go/parser"

	stack "github.com/vedadiyan/gocollections/pkg/stack"
)

type Expression[T any] func(ctx T, args []any) (any, error)

type Godyn[T any] struct {
	expressions map[string]Expression[T]
}

func New[T any](expressions map[string]Expression[T]) Godyn[T] {
	godyn := Godyn[T]{}
	godyn.expressions = expressions
	return godyn
}

func (context Godyn[T]) Invoke(ctx T, expr string) (any, error) {
	parsed, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, err
	}
	callExpr, ok := parsed.(*ast.CallExpr)
	if !ok {
		return nil, errors.New("expression is not a call expression")
	}
	return context.eval(ctx, callExpr)
}

func (context Godyn[T]) eval(ctx T, callExp *ast.CallExpr) (any, error) {
	ident := callExp.Fun.(*ast.Ident)
	args := make([]any, 0)
	for _, arg := range callExp.Args {
		switch t := arg.(type) {
		case *ast.BasicLit:
			{
				args = append(args, t.Value)
			}
		case *ast.CallExpr:
			{
				value, err := context.eval(ctx, t)
				if err != nil {
					return nil, err
				}
				args = append(args, value)
			}
		case *ast.BinaryExpr:
			{
				_, value, err := evalBinary(ctx, &context, t)
				if err != nil {
					return nil, err
				}
				args = append(args, value)
			}
		case *ast.Ident:
			{
				args = append(args, t.Name)
			}
		case *ast.SelectorExpr:
			{
				var buffer bytes.Buffer
				x := t.X
				q := stack.New[string]()
				for {
					sel, ok := x.(*ast.SelectorExpr)
					if !ok {
						ident := x.(*ast.Ident)
						q.Push(ident.Name)
						break
					}
					x = sel.X
					q.Push(sel.Sel.Name)
				}
				for !q.IsEmpty() {
					val, err := q.Pop()
					if err != nil {
						return nil, err
					}
					buffer.WriteString(val)
					buffer.WriteString(".")
				}
				buffer.WriteString(t.Sel.Name)
				args = append(args, buffer.String())
			}
		default:
			{
				return nil, errors.New("unsupported operation")
			}
		}
	}
	return context.expressions[ident.Name](ctx, args)
}
