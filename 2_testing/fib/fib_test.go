package fib

import (
	"testing"
)

var a int
func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b     = 1
	//)

	a = 1
	//a := 1
	b := 1

	//fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
	//fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	a, b = b, a
	//tmp := a
	//a = b
	//b = tmp

	t.Log(a, b)
}





