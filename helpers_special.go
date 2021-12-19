package pprnt

import (
	"fmt"
	"reflect"
)

func _PrintSliceOrArray(arg interface{}, key *string, endChar *string) string {
	str := ""
	lvlStr := _CreateLevelString()

	// Initial Setup
	if key != nil {
		str += lvlStr + "\"" + *key + "\"" + ": "
	} else {
		str += lvlStr
	}

	str += "[\n"
	fmt.Print(str)
	_state.Level++

	// Processing each element
	items := reflect.ValueOf(arg)
	items = reflect.Indirect(items)
	for i := 0; i < items.Len(); i++ {
		item := items.Index(i)
		item = reflect.Indirect(item)

		str += _ProcessArchitecture(item.Interface(), key, endChar)
	}

	// Final Setup
	_state.Level--
	finalStrToPrint := lvlStr + "]"
	if _state.Level > 0 {
		finalStrToPrint += ","
	}

	if endChar != nil {
		finalStrToPrint += *endChar
	} else {
		finalStrToPrint += "\n"
	}

	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}

func _PrintMap(arg interface{}, key *string, endChar *string) string {
	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	if key != nil {
		str += lvlStr + "\"" + *key + "\"" + ": "
	} else {
		str += lvlStr
	}

	str += "{\n"
	fmt.Print(str)
	_state.Level++

	// Processing each element

	// Final Setup
	_state.Level--
	finalStrToPrint := lvlStr + "}"
	if _state.Level > 0 {
		finalStrToPrint += ","
	}

	if endChar != nil {
		finalStrToPrint += *endChar
	} else {
		finalStrToPrint += "\n"
	}

	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}

func _PrintStruct(arg interface{}, key *string, endChar *string) string {
	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	if key != nil {
		str += lvlStr + "\"" + *key + "\"" + ": "
	} else {
		str += lvlStr
	}

	str += "{\n"
	fmt.Print(str)
	_state.Level++

	// Processing each element
	structValues := reflect.ValueOf(arg)
	structValues = reflect.Indirect(structValues)

	for j := 0; j < structValues.NumField(); j++ {
		fieldItem := structValues.Field(j)
		str += _ProcessArchitecture(fieldItem.Interface(), key, endChar)
	}

	// Final Setup
	_state.Level--
	finalStrToPrint := lvlStr + "}"
	if _state.Level > 0 {
		finalStrToPrint += ","
	}

	if endChar != nil {
		finalStrToPrint += *endChar
	} else {
		finalStrToPrint += "\n"
	}

	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}
