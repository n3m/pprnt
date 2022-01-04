package pprnt

import (
	"fmt"
)

const formatString string = "\"%+v\""
const formatInt string = "%d"
const formatBoolean string = "%t"
const formatNil string = "%s"
const formatFloat string = "%+v"

func _PrintString(elem interface{}, key, endChar *string) string {
	return _DoPrintPrimitive(formatString, key, endChar, elem)
}

func _PrintInt(elem interface{}, key, endChar *string) string {
	return _DoPrintPrimitive(formatInt, key, endChar, elem)
}

func _PrintBoolean(elem interface{}, key, endChar *string) string {
	return _DoPrintPrimitive(formatBoolean, key, endChar, elem)
}

func _PrintNil(_ interface{}, key, endChar *string) string {
	return _DoPrintPrimitive(formatNil, key, endChar, "<nil>")
}

func _PrintFloat(elem interface{}, key, endChar *string) string {
	return _DoPrintPrimitive(formatFloat, key, endChar, elem)
}

func _DoPrintPrimitive(format string, key, endChar *string, elem interface{}) string {
	lvlStr := _CreateLevelString()

	strToPrint := ""

	if key != nil {
		strToPrint += lvlStr + "\"" + *key + "\"" + ": "
		key = nil
	} else {
		strToPrint += lvlStr
	}

	strToPrint += fmt.Sprintf(format, elem)

	if _state.Level > 0 {
		strToPrint += ","
	}

	if endChar != nil {
		strToPrint += *endChar
	} else {
		strToPrint += "\n"
	}

	fmt.Print(strToPrint)
	return fmt.Sprint(strToPrint)
}
