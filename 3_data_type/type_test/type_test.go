package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	//var a int = 1
	var a int32 = 1
	var b int64

	b = int64(a)
	var c MyInt
	//c = b // false
	c = MyInt(b) // true

	// when a==int -> false
	// when a==int32 -> false
	//b = a // false

	t.Log(a, b, c)

}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr += 1 // error
	t.Log(a, aPtr, *aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
