package pprnt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_PrintMapPointerString(t *testing.T) {
	caseT := &map[string]string{
		"At1": "t1",
		"At2": "t2",
	}
	output := Print(caseT)

	tests := []string{
		"\t\t\"At1\": \"t1\",\n",
		"\t\t\"At2\": \"t2\",\n",
	}

	for _, test := range tests {
		if !assert.Contains(t, output, test) {
			t.Fatalf("Case test failed > " + test)
		}
	}
}

func Test_PrintMapPointerInterface(t *testing.T) {
	caseT := map[string]interface{}{
		"At1": GetStringAdddress("t1"),
		"At2": GetStringAdddress("t2"),
	}
	output := Print(caseT)

	tests := []string{
		"\t\t\"At1\": \"t1\",\n",
		"\t\t\"At2\": \"t2\",\n",
	}

	for _, test := range tests {
		if !assert.Contains(t, output, test) {
			t.Fatalf("Case test failed > " + test)
		}
	}
}

func Test_PrintMapPointerValueString(t *testing.T) {
	caseT := &_Test{
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

func Test_PrintHyperAnnidated(t *testing.T) {

	caseTest := []interface{}{
		&_Test{
			A: "Howdy",
			B: 2,
			C: true,
			D: 1.5000000004,
			ARR: []string{
				"TEST",
			},
		},
		map[string]interface{}{
			"At1": "t1",
			"At2": GetStringAdddress("t2"),
			"Stct1": &_Test{
				A: "Cowboy",
				B: 3,
				C: false,
				D: 1,
			},
			"Data": primitive.NilObjectID,
		},
	}

	output := Print(caseTest)

	tests := []string{
		"\t\t{",
		"\t\t\t\t\"A\": \"Howdy\",",
		"\t\t\t\t\"B\": 2,",
		"\t\t\t\t\"C\": true,",
		"\t\t\t\t\"D\": 1.5000000004,",
		"\t\t\t\t\"ARR\": [",
		"\t\t\t\t\t\t\"TEST\",",
		"\t\t\t\t],",
		"\t\t},",
		"\t\t{",
		"\t\t\t\t\"At1\": \"t1\",",
		"\t\t\t\t\"At2\": \"t2\",",
		"\t\t\t\t\"Stct1\": {",
		"\t\t\t\t\t\t\"A\": \"Cowboy\",",
		"\t\t\t\t\t\t\"B\": 3,",
		"\t\t\t\t\t\t\"C\": false,",
		"\t\t\t\t\t\t\"D\": 1,",
		"\t\t\t\t\t\t\"ARR\": [",
		"\t\t\t\t\t\t],",
		"\t\t}",
	}

	for _, test := range tests {
		if !assert.Contains(t, output, test) {
			t.Fatalf("Case test failed > " + test)
		}
	}
}
