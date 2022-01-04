package pprnt

import (
	"fmt"
	"reflect"
	"strings"
)

var indentation int = 2

var lvl string = strings.Repeat("\t", indentation)

func _CreateLevelString() string {
	return strings.Repeat(lvl, _state.Level)
}

func _ProcessArchitecture(arg interface{}, key, endChar *string) string {
	value := reflect.ValueOf(arg)

	if arg == nil {
		return _PrintNil(arg, key, endChar)
	}

	switch reflect.TypeOf(arg).Kind() {

	case reflect.Slice, reflect.Array:
		return _PrintSliceOrArray(reflect.Indirect(value).Interface(), key, endChar)
	case reflect.Map:
		return _PrintMap(reflect.Indirect(value).Interface(), key, endChar)
	case reflect.Struct:
		return _PrintStruct(reflect.Indirect(value).Interface(), key, endChar)
	// case reflect.Chan:
	// case reflect.Interface:
	case reflect.Ptr:
		newValue := value.Elem()
		newValue = reflect.Indirect(newValue)
		return _ProcessArchitecture(newValue.Interface(), key, endChar)
		// case reflect.UnsafePointer:
		// case reflect.Uintptr:
	}

	return _PrintPrimitive(reflect.Indirect(value).Interface(), key, endChar)
}

func _PrintPrimitive(arg interface{}, key, endChar *string) string {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.String:
		return _PrintString(arg, key, endChar)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return _PrintInt(arg, key, endChar)
	case reflect.Bool:
		return _PrintBoolean(arg, key, endChar)
	case reflect.Float32, reflect.Float64:
		return _PrintFloat(arg, key, endChar)
	default:

		fmt.Printf("[!!] %+v\n", arg)
		return fmt.Sprintf("%+v\n", arg)
	}
}
