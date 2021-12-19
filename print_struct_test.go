package pprnt

import (
	"testing"
)

type _Test struct {
	A string
	B int
	C bool
	D float64
	F struct {
		G string
	}
	Z []interface{}
}

func Test_PrintStruct(t *testing.T) {
	cases := []_Test{
		{
			A: "Howdy",
			B: 2,
			C: true,
			D: 1.5000000004,
			F: struct {
				G string
			}{
				G: "God's eye",
			},
			Z: []interface{}{
				"Hello world",
				"Mamma mia",
				true,
				false,
				1,
				2.5,
				&_Test{
					A: "Howdy 2",
				},
				nil,
			},
		},
	}

	for _, each := range cases {
		output := Print(each)
		t.Fatal(output)
	}

}
