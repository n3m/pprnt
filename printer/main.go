package printer

import (
	"fmt"
	"os"
	"reflect"

	"github.com/DrN3MESiS/pprnt/helpers"
	"github.com/mattn/go-isatty"
)

var (
	//IdentString ...
	IdentString string = "  "
	//NoColor ...
	NoColor bool = !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())
)

const (
	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m" //Not supported
	colorGreen  string = "\033[32m" //String
	colorYellow string = "\033[33m" //Keys
	colorBlue   string = "\033[34m" //Bool
	colorPurple string = "\033[35m" //Int floats
	colorCyan   string = "\033[36m" //Brackets or Corch
	colorGray   string = "\033[37m"
	colorWhite  string = "\033[97m"
)

//PrintMap ...
func PrintMap(MAP map[string]interface{}, depth int, detailMode bool) error {
	stringToPrint := _CreateDepthString(depth)

	_PrintSymbol("{", stringToPrint)

	//INIT
	newIdentDepth := depth + 1
	newDepthString := _CreateDepthString(newIdentDepth)
	for key, value := range MAP {

		_PrintMapKey(key, newDepthString)

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
		_PrintSymbol("},", stringToPrint)
	} else {
		_PrintSymbol("}", stringToPrint)
	}
	return nil
}

//PrintArray ...
func PrintArray(ARR []interface{}, depth int, detailMode bool) error {
	stringToPrint := _CreateDepthString(depth)
	_PrintSymbol("[", stringToPrint)

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
		_PrintSymbol("],", stringToPrint)
	} else {
		_PrintSymbol("]", stringToPrint)
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

func _PrintSymbol(symbol, stringToPrint string) {
	if NoColor {
		fmt.Println(fmt.Sprintf("%s%+v", stringToPrint, symbol))
	} else {
		fmt.Println(colorCyan, fmt.Sprintf("%s%+v", stringToPrint, symbol))
	}
}

func _PrintMapKey(key, newDepthString string) {
	if NoColor {
		fmt.Print(fmt.Sprintf("%s'%+v': ", newDepthString, key))
	} else {
		fmt.Print(colorYellow, fmt.Sprintf("%s'%+v': ", newDepthString, key))
	}
}

func _PrintValue(value interface{}, detailMode bool) {
	switch reflect.ValueOf(value).Kind() {
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:

		if detailMode {
			if NoColor {
				fmt.Println(fmt.Sprintf("%#v,", value))
			} else {
				fmt.Println(colorPurple, fmt.Sprintf("%#v,", value))
			}
		} else {
			if NoColor {
				fmt.Println(fmt.Sprintf("%+v,", value))
			} else {
				fmt.Println(colorPurple, fmt.Sprintf("%+v,", value))
			}
		}
		break
	case reflect.Bool:
		if detailMode {
			if NoColor {
				fmt.Println(fmt.Sprintf("%#v,", value))
			} else {
				fmt.Println(colorBlue, fmt.Sprintf("%#v,", value))
			}
		} else {
			if NoColor {
				fmt.Println(fmt.Sprintf("%+v,", value))
			} else {
				fmt.Println(colorBlue, fmt.Sprintf("%+v,", value))
			}
		}
		break
	case reflect.Float32, reflect.Float64:
		if detailMode {
			if NoColor {
				fmt.Println(fmt.Sprintf("%#v,", value))
			} else {
				fmt.Println(colorPurple, fmt.Sprintf("%#v,", value))
			}
		} else {
			if NoColor {
				fmt.Println(fmt.Sprintf("%+v,", value))
			} else {
				fmt.Println(colorPurple, fmt.Sprintf("%+v,", value))
			}
		}
		break
	case reflect.String:
		if detailMode {
			if NoColor {
				fmt.Println(fmt.Sprintf("%#v,", value))
			} else {
				fmt.Println(colorGreen, fmt.Sprintf("%#v,", value))
			}
		} else {
			if NoColor {
				fmt.Println(fmt.Sprintf("%+v,", value))
			} else {
				fmt.Println(colorGreen, fmt.Sprintf("%+v,", value))
			}
		}
		break
	default:

		if NoColor {
			fmt.Println("{{Type not supported by PPRNT}}")
		} else {
			fmt.Println(colorRed, "{{Type not supported by PPRNT}}")
		}

		break
	}

}
