package pprnt

/*
	Version: 1.10.3
	Author: Alan Maldonado
*/

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/n3m/pprnt/cleaner"
	"github.com/n3m/pprnt/helpers"
	"github.com/n3m/pprnt/printer"
)

var (
	detailMode   bool = false
	initialDepth int  = 0
)

//SetDetailMode ...
/*
This function allows the developer to change the printing "Detail Mode", which in summary,
prints the type of the value that it has read.

Param: "code"
Expects an integer: '1' to enable "Detail Mode" or '0' to disable "Detail Mode"

TL;DR: '1' to enable, '0' to disable
*/
func SetDetailMode(code int) {
	switch code {
	case 1:
		detailMode = true
		break
	case 0:
		detailMode = false
		break
	}
}

//SetIdentLength ...
/*
This function allows the developer to change the number of spaces defined as "Identantion"
while printing a value.

Param: "length"
Expects a positive integer.

TL;DR: The given value in the parameter will be the number of spaces in the "Identation"
*/
func SetIdentLength(length int) {
	if length < 0 {
		panic("SetIdentLength() :: Function expects a positive length value")
	}
	printer.IdentString = strings.Repeat(" ", length)
}

//Print ...
/*
This function allows the developer to print any type of variable with a prettier printing
structure

Param: "arg"
Expects any type of variable of any type of value.

TL;DR: Prints a single parameter
*/
func Print(arg interface{}) {
	errMessage := "[PPRNT]"
	switch reflect.ValueOf(arg).Kind() {
	case reflect.Map, reflect.Struct:
		MTP, err := helpers.ValueToMap(arg)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}

		err = printer.PrintMap(MTP, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	case reflect.Array, reflect.Slice:
		tempArray := arg.([]interface{})
		err := printer.PrintArray(tempArray, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Bool,
		reflect.Float32, reflect.Float64, reflect.String:
		printer.PrintValue(arg, detailMode)
		break
	default:
		err := printer.PrintNormal(arg, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}

		break
	}

	fmt.Println(printer.ColorReset)

}

//MPrint ...
/*
This function allows the developer to print an 'N' amount of variables with a prettier printing
structure

Param: "args"
Expects an 'N' amount of paramters of any type

TL;DR: Prints an 'N' amount of given parameters
*/
func MPrint(args ...interface{}) {
	for _, arg := range args {
		Print(arg)
	}

}

//SuperDepthMapCleaning ...
/*
This function allows the developer to "Clean" a map[string]interface{} structure from any 'nil' fields.
Basically it removes the fields with value of type 'nil'.

Param: "aMap"
Expects a map[string]interface{}

TL;DR: Cleans a map from all nil key:values
*/
func SuperDepthMapCleaning(aMap map[string]interface{}) (map[string]interface{}, error) {
	return cleaner.CleanMap(aMap), nil
}
