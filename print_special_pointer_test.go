package pprnt

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Contains(t, output, test)
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
		assert.Contains(t, output, test)
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
		assert.Contains(t, output, test)
	}
}
