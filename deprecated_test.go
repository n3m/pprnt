package pprnt

import "testing"

func Test_Deprecated_Print(_ *testing.T) {

	Deprecated.CleanMap(map[string]interface{}{})

	Deprecated.CleanArray([]interface{}{})

}
