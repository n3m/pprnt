package pprnt

/*
	Version: 1.0.0
	Author: Alan Maldonado

	== OpenSource Project ==
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

//Print ...
/**/
func Print(data interface{}) error {
	errMessage := "[pprnt][Print()] > "

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		mapData, err := structToMap(data)
		if err != nil {
			log.Printf("%+v", errMessage+err.Error())
			return errors.New(errMessage + err.Error())
		}

		depth := 1
		printData(mapData, &depth)

	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]interface{}{}) {
		depth := 1
		printData(data.(map[string]interface{}), &depth)

	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]string{}) {
		depth := 1
		toSend := map[string]interface{}{}
		for key, value := range data.(map[string]string) {
			toSend[key] = value
		}
		printData(toSend, &depth)

	} else {
		fmt.Printf("%+v", data)
	}

	return nil
}

func printData(mapData map[string]interface{}, depth *int) error {
	errMessage := "[printData()] > "
	depthStr := ""
	tempDepth := *depth

	for i := 0; i < *depth; i++ {
		depthStr += "   "
	}

	if *depth == 1 {
		fmt.Printf("{\n")
	}

	for key, value := range mapData {
		if reflect.ValueOf(value).Kind() == reflect.Struct {
			mapData, err := structToMap(value)
			if err != nil {
				log.Printf("%+v", errMessage+err.Error())
				return errors.New(errMessage + err.Error())
			}
			printData(mapData, depth)
		} else {
			switch value.(type) {
			case map[string]interface{}:
				fmt.Printf("%+v\"%+v\": {\n", depthStr, key)
				*depth++
				printData(value.(map[string]interface{}), depth)
				break
			default:
				fmt.Printf("%+v\"%+v\": %+v\n", depthStr, key, value)
				break
			}
		}

	}

	if tempDepth > 1 {
		fmt.Printf(strings.Repeat("   ", tempDepth-1) + "}\n")
	} else {
		fmt.Printf("}\n")
	}

	return nil
}

func structToMap(data interface{}) (map[string]interface{}, error) {
	errMessage := "[structToMap()] > "
	m := make(map[string]interface{})

	j, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(errMessage + err.Error() + " :: [001]")
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, errors.New(errMessage + err.Error() + " :: [002]")
	}

	return m, nil
}
