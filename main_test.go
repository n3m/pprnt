package pprnt

import (
	"log"
	"testing"
)

type Server struct {
	Name    string
	ID      int32
	Enabled bool
	Server  *Server2
}

type Server2 struct {
	ID string
}

func TestPrint(t *testing.T) {

	example2 := Server{Name: "John", ID: 255, Enabled: true, Server: &Server2{ID: "test"}}

	// example := map[string]interface{}{"get": map[string]interface{}{"some": "brother"}}
	err := Print(example2)
	if err != nil {
		log.Printf("> %+v", err)
		t.Fatal(err)
	}
}
