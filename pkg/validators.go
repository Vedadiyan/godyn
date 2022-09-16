package godyn

import (
	"errors"
	"reflect"
)

type Type string

const (
	STRING         Type = "string"
	INT            Type = "int"
	INT64          Type = "int64"
	UINT           Type = "uint"
	UINT64         Type = "uint64"
	SHORT          Type = "short"
	USHORT         Type = "ushort"
	BYTE           Type = "byte"
	BOOL           Type = "bool"
	FLOAT          Type = "float"
	DOUBLE         Type = "double"
	POINTER_STRING Type = "pointer string"
	POINTER_INT    Type = "pointer int"
	POINTER_INT64  Type = "pointer int64"
	POINTER_UINT   Type = "pointer uint"
	POINTER_UINT64 Type = "pointer uint64"
	POINTER_SHORT  Type = "pointer short"
	POINTER_USHORT Type = "pointer ushort"
	POINTER_BYTE   Type = "pointer byte"
	POINTER_BOOL   Type = "pointer bool"
	POINTER_FLOAT  Type = "pointer float"
	POINTER_DOUBLE Type = "pointer double"
	ANY            Type = "any"
)

func ValidateArguments(signature []Type, args []any) error {
	signatureLen := len(signature)
	argsLen := len(args)
	if signatureLen != argsLen {
		return InvalidNumberOfArgumentsError(signatureLen, argsLen)
	}
	for index, value := range signature {
		argValue := args[index]
		var ok bool
		switch value {
		case STRING:
			{
				_, ok = argValue.(string)
			}
		case INT:
			{
				_, ok = argValue.(int)
			}
		case INT64:
			{
				_, ok = argValue.(int64)
			}
		case UINT:
			{
				_, ok = argValue.(uint)
			}
		case UINT64:
			{
				_, ok = argValue.(uint64)
			}
		case SHORT:
			{
				_, ok = argValue.(int16)
			}
		case USHORT:
			{
				_, ok = argValue.(uint16)
			}
		case BYTE:
			{
				_, ok = argValue.(byte)
			}
		case BOOL:
			{
				_, ok = argValue.(bool)
			}
		case FLOAT:
			{
				_, ok = argValue.(float32)
			}
		case DOUBLE:
			{
				_, ok = argValue.(float64)
			}
		case POINTER_STRING:
			{
				_, ok = argValue.(*string)
			}
		case POINTER_INT:
			{
				_, ok = argValue.(*int)
			}
		case POINTER_INT64:
			{
				_, ok = argValue.(*int64)
			}
		case POINTER_UINT:
			{
				_, ok = argValue.(*uint)
			}
		case POINTER_UINT64:
			{
				_, ok = argValue.(*uint64)
			}
		case POINTER_SHORT:
			{
				_, ok = argValue.(*int16)
			}
		case POINTER_USHORT:
			{
				_, ok = argValue.(*uint16)
			}
		case POINTER_BYTE:
			{
				_, ok = argValue.(*byte)
			}
		case POINTER_BOOL:
			{
				_, ok = argValue.(*bool)
			}
		case POINTER_FLOAT:
			{
				_, ok = argValue.(*float32)
			}
		case POINTER_DOUBLE:
			{
				_, ok = argValue.(*float64)
			}
		case ANY:
			{
				ok = true
			}
		default:
			{
				return errors.New("unknown type")
			}
		}
		if !ok {
			return ArgumentMismatchError(index, string(value), reflect.TypeOf(argValue).Elem().Name())
		}
	}
	return nil
}
