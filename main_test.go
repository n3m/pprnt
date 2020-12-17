package pprnt

import (
	"log"
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
		log.Printf("> Test: %+v", key)
		Print(test)
	}

	log.Printf("= = =")

	DetailMode(1)
	for key, test := range tests {
		log.Printf("> Test: %+v", key)
		Print(test)
	}
}

func TestPrintMap(t *testing.T) {

}
