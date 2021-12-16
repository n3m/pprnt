package pprnt

import (
	"fmt"
	"log"
)

const formatString string = "\"%+v\""
const formatInt string = "%d"
const formatBoolean string = "%t"
const formatRune string = "'%+v'"
const formatFloat string = "%+v"

func _PrintString(elem string) string {
	log.Printf(formatString, elem)
	return fmt.Sprintf(formatString, elem)
}

func _PrintInt(elem interface{}) string {
	log.Printf(formatInt, elem)
	return fmt.Sprintf(formatInt, elem)
}

func _PrintBoolean(elem bool) string {
	log.Printf(formatBoolean, elem)
	return fmt.Sprintf(formatBoolean, elem)
}

func _PrintRune(elem rune) string {
	log.Printf(formatRune, elem)
	return fmt.Sprintf(formatRune, elem)
}

func _PrintFloat(elem interface{}) string {
	log.Printf(formatFloat, elem)
	return fmt.Sprintf(formatFloat, elem)
}
