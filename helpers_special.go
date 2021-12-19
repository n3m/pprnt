package pprnt

import (
	"fmt"
	"reflect"
)

func _PrintSpecialInitSetup(str, lvlStr string, key *string, initSetupChar string) (string, *string) {
	if key != nil {
		str += lvlStr + "\"" + *key + "\"" + ": "
		key = nil
	} else {
		str += lvlStr
	}

	str += initSetupChar + "\n"
	fmt.Print(str)

	return str, key
}

func _PrintSpecialEndSetup(str, lvlStr string, endChar *string, endSetupChar string) string {
	finalStrToPrint := lvlStr + endSetupChar
	if _state.Level > 0 {
		finalStrToPrint += ","
	}

	if endChar != nil {
		finalStrToPrint += *endChar
	} else {
		finalStrToPrint += "\n"
	}

	return finalStrToPrint
}

func _PrintSliceOrArray(arg interface{}, key *string, endChar *string) string {
	str := ""
	lvlStr := _CreateLevelString()

	// Initial Setup
	str, key = _PrintSpecialInitSetup(str, lvlStr, key, "[")
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
	finalStrToPrint := _PrintSpecialEndSetup(str, lvlStr, endChar, "]")
	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}

func _PrintMap(arg interface{}, key *string, endChar *string) string {
	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	str, _ = _PrintSpecialInitSetup(str, lvlStr, key, "{")
	_state.Level++

	// Processing each element
	mapObj := reflect.ValueOf(arg)
	for _, key := range mapObj.MapKeys() {
		value := mapObj.MapIndex(key).Interface()

		str += _ProcessArchitecture(value,
			GetStringAdddress(key.String()),
			endChar)
	}

	// Final Setup
	_state.Level--
	finalStrToPrint := _PrintSpecialEndSetup(str, lvlStr, endChar, "}")
	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}

func _PrintStruct(arg interface{}, key *string, endChar *string) string {
	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	str, _ = _PrintSpecialInitSetup(str, lvlStr, key, "{")
	_state.Level++

	// Processing each element
	structValues := reflect.ValueOf(arg)
	structValues = reflect.Indirect(structValues)

	for j := 0; j < structValues.NumField(); j++ {
		str += _ProcessArchitecture(structValues.Field(j).Interface(),
			GetStringAdddress(structValues.Type().Field(j).Name),
			endChar)
	}

	// Final Setup
	_state.Level--
	finalStrToPrint := _PrintSpecialEndSetup(str, lvlStr, endChar, "}")
	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}
