package pprnt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintSlice(t *testing.T) {
	cases := []interface{}{
		"Howdy",
		nil,
	}

	output := Print(cases)

	expected := "[\n\t\t\"Howdy\",\n\t\t<nil>,\n]\n"
	assert.Equal(t, expected, output)
}
