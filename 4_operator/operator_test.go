package __operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1,2,3,4}

	t.Log(a == b)
	//t.Log(a == c) // error: lenth not same
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable   // 0110
	a = a &^ Executeable  // 0010
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a& Executeable == Executeable)
}
