package main

/*
	Version: 0.2.0
	Author: Alan Maldonado

	== OpenSource Project ==
*/

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

//Print ...
func Print(data interface{}) {
	test := map[string]interface{}{}
	if reflect.ValueOf(data).Kind() == reflect.Struct {
		mapData := structs.Map(data)
		depth := 1
		printData(mapData, &depth)
	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(test) {
		depth := 1
		printData(data.(map[string]interface{}), &depth)
	} else {
		fmt.Printf("%+v", data)
	}

}

func printData(mapData map[string]interface{}, depth *int) {
	depthStr := ""
	tempDepth := *depth
	for i := 0; i < *depth; i++ {
		depthStr += "   "
	}

	if *depth == 1 {
		fmt.Printf("{\n")
	}

	for key, value := range mapData {
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

	if tempDepth > 1 {
		fmt.Printf(strings.Repeat("   ", tempDepth-1) + "}\n")
	} else {
		fmt.Printf("}\n")
	}
}

func main() {
	example := map[string]interface{}{"get": map[string]interface{}{"some": "brother"}}
	Print(example)
}
