package pprnt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintString(t *testing.T) {
	cases := []string{
		"Howdy",
		"Hello world",
		"God\"s eye",
		"Mamma mia",
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("\"%+v\"\n", c), output)
	}

}

func Test_PrintInt(t *testing.T) {
	cases := []int{
		111, 69, 420, 000, 646545465,
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", c), output)
	}
}

func Test_PrintBoolean(t *testing.T) {
	cases := []bool{
		true, false, false, true,
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", c), output)
	}
}

func Test_PrintFloat64(t *testing.T) {
	cases := []float64{
		1.1, 1.2, 1.3, 1.4, 1.5,
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", c), output)
	}
}
