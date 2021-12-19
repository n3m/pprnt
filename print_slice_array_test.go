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

	expected :=
		`[
		"Howdy",
		<nil>,
]
`

	assert.Equal(t, expected, output)
}
