package pprnt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type _Test struct {
	A   string
	B   int
	C   bool
	D   float64
	ARR []string
}

func Test_PrintStruct(t *testing.T) {
	caseT := _Test{
		A: "Howdy",
		B: 2,
		C: true,
		D: 1.5000000004,
		ARR: []string{
			"TEST",
		},
	}

	output := Print(caseT)

	tests := []string{
		"\t\t\"A\": \"Howdy\",\n",
		"\t\t\"B\": 2,\n",
		"\t\t\"C\": true,\n",
		"\t\t\"D\": 1.5000000004,\n",
		"\t\t\"ARR\": [\n",
		"\t\t\t\t\"TEST\",\n",
		"\t\t],\n",
	}

	for _, test := range tests {
		if !assert.Contains(t, output, test) {
			t.Fatalf("Case test failed > " + test)
		}
	}
}

func Test_PrintMap(t *testing.T) {
	caseT := map[string]interface{}{
		"A": "Howdy",
		"B": 2,
		"C": true,
		"D": 1.5000000004,
	}

	output := Print(caseT)

	tests := []string{
		"\t\t\"A\": \"Howdy\",\n",
		"\t\t\"B\": 2,\n",
		"\t\t\"C\": true,\n",
		"\t\t\"D\": 1.5000000004,\n",
	}

	for _, test := range tests {
		if !assert.Contains(t, output, test) {
			t.Fatalf("Case test failed > " + test)
		}
	}

	// 	expected :=
	// 		`{
	// 		"A": "Howdy",
	// 		"B": 2,
	// 		"C": true,
	// 		"D": 1.5000000004,
	// }
	// `
}
