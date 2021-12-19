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
	expected :=
		`{
		"A": "Howdy",
		"B": 2,
		"C": true,
		"D": 1.5000000004,
		"ARR": [
				"TEST",
		],
}
`

	assert.Equal(t, expected, output)
}

func Test_PrintMap(t *testing.T) {
	caseT := map[string]interface{}{
		"A": "Howdy",
		"B": 2,
		"C": true,
		"D": 1.5000000004,
	}

	output := Print(caseT)
	expected :=
		`{
		"A": "Howdy",
		"B": 2,
		"C": true,
		"D": 1.5000000004,
}
`
	assert.Equal(t, expected, output)

}
