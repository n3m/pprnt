package pprnt

import (
	"reflect"
)

func Print(arg interface{}) string {
	switch arg := arg.(type) {
	case string:
		return _PrintString(arg)
	case int, int8, int16, int64, uint, uint8, uint16, uint32, uint64:
		return _PrintInt(arg)
	case bool:
		return _PrintBoolean(arg)
	case rune:
		return _PrintRune(arg)
	case float32, float64:
		return _PrintFloat(arg)
	default:

		switch reflect.TypeOf(arg).Kind() {
		case reflect.Slice:
		case reflect.Array:
		case reflect.Map:
		case reflect.Struct:
		case reflect.Chan:
		case reflect.Interface:
		case reflect.Ptr:
		case reflect.UnsafePointer:
		case reflect.Uintptr:
		default:

		}
	}

}
