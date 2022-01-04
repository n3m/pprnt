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

func _PrintSliceOrArray(arg interface{}, key, endChar *string) string {
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

		var value interface{}

		if item.CanInterface() {
			if item.IsZero() {
				value = nil
			} else {
				value = item.Interface()
			}
		} else {
			value = nil
		}

		str += _ProcessArchitecture(value, key, endChar)

	}

	// Final Setup
	_state.Level--
	finalStrToPrint := _PrintSpecialEndSetup(str, lvlStr, endChar, "]")
	str += finalStrToPrint
	fmt.Print(finalStrToPrint)
	return str
}

func _PrintMap(arg interface{}, key, endChar *string) string {
	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	str, _ = _PrintSpecialInitSetup(str, lvlStr, key, "{")
	_state.Level++

	// Processing each element
	mapObj := reflect.ValueOf(arg)
	for _, key := range mapObj.MapKeys() {

		var value interface{}

		if mapObj.MapIndex(key).CanInterface() {
			if mapObj.MapIndex(key).IsZero() {
				value = nil
			} else {
				value = mapObj.MapIndex(key).Interface()
			}
		} else {
			value = nil
		}

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

func _PrintStruct(arg interface{}, key, endChar *string) string {

	lvlStr := _CreateLevelString()
	str := ""

	// Initial Setup
	str, _ = _PrintSpecialInitSetup(str, lvlStr, key, "{")
	_state.Level++

	// Processing each element
	structValues := reflect.ValueOf(arg)
	structValues = reflect.Indirect(structValues)

	for j := 0; j < structValues.NumField(); j++ {
		field := structValues.Field(j)

		var value interface{}

		if field.CanInterface() {
			value = field.Interface()
		} else {
			value = nil
		}
		str += _ProcessArchitecture(value,
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
