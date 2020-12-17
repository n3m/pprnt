package printer

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/DrN3MESiS/pprnt/helpers"
)

var (
	//IdentString ...
	IdentString string = "  "
)

func PrintData(mapData map[string]interface{}, depth *int) error {
	errMessage := "[PrintData()] > "
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
			mapData, err := helpers.StructToMap(value)
			if err != nil {
				log.Printf("%+v", errMessage+err.Error())
				return errors.New(errMessage + err.Error())
			}
			PrintData(mapData, depth)
		} else {
			switch value.(type) {
			case map[string]interface{}:
				fmt.Printf("%+v\"%+v\": {\n", depthStr, key)
				*depth++
				PrintData(value.(map[string]interface{}), depth)
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

//PrintMap ...
func PrintMap(MAP map[string]interface{}, depth int, detailMode bool) error {
	stringToPrint := _CreateDepthString(depth)

	fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "{"))

	//INIT
	newIdentDepth := depth + 1
	newDepthString := _CreateDepthString(newIdentDepth)
	for key, value := range MAP {

		fmt.Print(fmt.Sprintf("%s'%+v': ", newDepthString, key))

		newMAPIdentDepth := depth + 1
		switch reflect.ValueOf(value).Kind() {
		//Case of Nested Maps or Structs
		case reflect.Map, reflect.Struct:
			//Value Print
			newMAP, err := helpers.ValueToMap(value)
			if err != nil {
				return err
			}
			newMAPIdentDepth := depth + 1
			PrintMap(newMAP, newMAPIdentDepth, detailMode)
			break
		//Case of Nested Array
		case reflect.Array, reflect.Slice:
			arrVal := value.([]interface{})
			PrintArray(arrVal, newMAPIdentDepth, detailMode)
			break
		//Case of Regular Value
		default:
			if detailMode {
				fmt.Printf("%#v,\n", value)
			} else {
				fmt.Printf("%+v,\n", value)
			}
			break
		}

	}
	//FINISH

	if depth > 0 {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "},"))
	} else {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "}"))
	}
	return nil
}

//PrintArray ...
func PrintArray(ARR []interface{}, depth int, detailMode bool) error {
	stringToPrint := _CreateDepthString(depth)
	fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "["))

	//INIT
	newIdentDepth := depth + 1
	newDepthString := _CreateDepthString(newIdentDepth)
	for i, value := range ARR {
		fmt.Print(fmt.Sprintf("%s[%d]: ", newDepthString, i))
		newMAPIdentDepth := depth + 1
		switch reflect.ValueOf(value).Kind() {
		//Case of Nested Maps or Structs
		case reflect.Map, reflect.Struct:
			//Value Print
			newMAP, err := helpers.ValueToMap(value)
			if err != nil {
				return err
			}
			newMAPIdentDepth := depth + 1
			PrintMap(newMAP, newMAPIdentDepth, detailMode)
			break
		//Case of Nested Array
		case reflect.Array, reflect.Slice:
			arrVal := value.([]interface{})
			PrintArray(arrVal, newMAPIdentDepth, detailMode)
			break
		//Case of Regular Value
		default:
			if detailMode {
				fmt.Printf("%#v,\n", value)
			} else {
				fmt.Printf("%+v,\n", value)
			}
			break
		}

	}
	//FINISH

	if depth > 0 {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "],"))
	} else {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, "]"))
	}
	return nil
}

//PrintNormal ...
func PrintNormal(value interface{}, depth int, detailMode bool) error {
	stringToPrint := _CreateDepthString(depth)
	if detailMode {
		fmt.Println(fmt.Sprintf("%s%#v", stringToPrint, value))
	} else {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, value))
	}
	return nil
}

func _CreateDepthString(depth int) string {
	depthString := ""
	for i := 0; i < depth; i++ {
		depthString += IdentString
	}

	return depthString
}
