package pprnt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintPointerString(t *testing.T) {
	cases := []*string{
		GetStringAdddress("t1"),
		GetStringAdddress("t2"),
		GetStringAdddress("t3"),
		GetStringAdddress("t4"),
		GetStringAdddress("t5"),
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("\"%+v\"\n", *c), output)
	}
}

func Test_PrintPointerInt(t *testing.T) {
	cases := []*int{
		new(int), new(int), new(int), new(int), new(int),
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", *c), output)
	}
}

func Test_PrintPointerBoolean(t *testing.T) {
	cases := []*bool{
		new(bool), new(bool), new(bool), new(bool), new(bool),
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", *c), output)
	}
}

func Test_PrintPointerFloat64(t *testing.T) {
	cases := []*float64{
		new(float64), new(float64), new(float64), new(float64), new(float64),
	}

	for _, c := range cases {
		output := Print(c)

		assert.Equal(t, fmt.Sprintf("%+v\n", *c), output)
	}
}
