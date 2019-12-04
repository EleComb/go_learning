package _1_encapsulation_interface

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (goProgrammer *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
