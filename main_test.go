package pprnt

import (
	"fmt"
	"testing"
)

type StructToTest struct {
	Name    string
	ID      int32
	Enabled bool
	Server  *InnerStructForTesting
}

type InnerStructForTesting struct {
	ID string
}

func TestSuperDepthMapCleaning(t *testing.T) {
	curMap := map[string]interface{}{
		"test1": nil,
		"test2": 1,
		"test3": map[string]interface{}{
			"test3.1": "lol",
			"test3.2": nil,
			"test3.3": []interface{}{
				nil,
				1,
				2,
			},
		},
	}

	Print(curMap)
	curMap, err := SuperDepthMapCleaning(curMap)
	if err != nil {
		t.Fatal(err)
	}
	Print(curMap)
}

func TestPrintNormal(t *testing.T) {
	tests := map[string]interface{}{
		"int":    1,
		"float":  5.2,
		"bool":   false,
		"string": "some string",
	}

	for key, test := range tests {
		fmt.Printf("> Test: %+v\n", key)
		Print(test)
	}

	fmt.Printf("= = =\n")

	SetDetailMode(1)
	for key, test := range tests {
		fmt.Printf("> Test: %+v\n", key)
		Print(test)
	}
}

func TestPrintMapStruct(t *testing.T) {
	tests := map[string]interface{}{
		"map": map[string]interface{}{
			"test1": nil,
			"test2": 1,
			"test3": map[string]interface{}{
				"test3.1": "lol",
				"test3.2": nil,
				"test3.3": []interface{}{
					nil,
					1,
					2,
				},
			},
		},
		"struct": StructToTest{
			Name:    "Alan",
			ID:      654,
			Enabled: true,
			Server: &InnerStructForTesting{
				ID: "897465",
			},
		},
	}

	for key, test := range tests {
		fmt.Printf("\n\n\n> Test: %+v\n", key)
		Print(test)
	}
}
func TestPrintArray(t *testing.T) {
	tests := map[string][]interface{}{
		"arr1": {
			1,
			map[string]interface{}{
				"test1": nil,
				"test2": 1,
				"test3": map[string]interface{}{
					"test3.1": "lol",
					"test3.2": nil,
					"test3.3": []interface{}{
						nil,
						1,
						2,
					},
				},
			},
		},
		"arr2": {
			3,
			nil,
			StructToTest{
				Name:    "Alan",
				ID:      654,
				Enabled: true,
				Server: &InnerStructForTesting{
					ID: "897465",
				},
			},
		},
		"arr3": {
			"data",
			"test",
		},
	}

	for key, test := range tests {
		fmt.Printf("\n\n\n> Test: %+v\n", key)
		Print(test)
	}
}
