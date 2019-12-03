package constant_test

import "testing"

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)

	a := 1
	//a := 7
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a& Executeable == Executeable)
}
