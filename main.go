package pprnt

/*
	Version: 2.0.0
	Author: Alan Maldonado

	== OpenSource Project ==
*/

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/DrN3MESiS/pprnt/cleaner"
	"github.com/DrN3MESiS/pprnt/helpers"
	"github.com/DrN3MESiS/pprnt/printer"
)

var (
	detailMode   bool = false
	initialDepth int  = 0
)

//DetailMode ...
/*
Code: 1 = Enables Detail Mode
Code: 0 = Disables Detail Mode
*/
func DetailMode(code int) {
	switch code {
	case 1:
		detailMode = true
		break
	case 0:
		detailMode = false
		break
	}
}

//IdentLength ...
/*
Defines the length of each identation level
*/
func IdentLength(length int) {
	printer.IdentString = strings.Repeat(" ", length)
}

//Print ...
/*
Print single object or value
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
		err := printer.PrintNormal(arg, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	default:
		if printer.NoColor {
			fmt.Println("{{Type not supported by PPRNT}}")
		} else {
			fmt.Println(printer.ColorRed, "{{Type not supported by PPRNT}}")
		}
		break
	}

}

//SuperDepthMapCleaning ...
func SuperDepthMapCleaning(curMap map[string]interface{}) (map[string]interface{}, error) {
	return cleaner.CleanMap(curMap), nil
}
