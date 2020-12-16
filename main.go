package pprnt

/*
	Version: 2.0.0
	Author: Alan Maldonado

	== OpenSource Project ==
*/

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/DrN3MESiS/pprnt/cleaner"
	"github.com/DrN3MESiS/pprnt/helpers"
	"github.com/DrN3MESiS/pprnt/printer"
)

//Print ...
/**/
func Print(data interface{}) error {
	errMessage := "[pprnt][Print()] > "

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		mapData, err := helpers.StructToMap(data)
		if err != nil {
			log.Printf("%+v", errMessage+err.Error())
			return errors.New(errMessage + err.Error())
		}

		depth := 1
		printer.PrintData(mapData, &depth)

	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]interface{}{}) {
		depth := 1
		printer.PrintData(data.(map[string]interface{}), &depth)

	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]string{}) {
		depth := 1
		toSend := map[string]interface{}{}
		for key, value := range data.(map[string]string) {
			toSend[key] = value
		}
		printer.PrintData(toSend, &depth)

	} else {
		fmt.Printf("%+v", data)
	}

	return nil
}

//SuperDepthMapCleaning ...
func SuperDepthMapCleaning(curMap map[string]interface{}) (map[string]interface{}, error) {
	return cleaner.CleanMap(curMap), nil
}
