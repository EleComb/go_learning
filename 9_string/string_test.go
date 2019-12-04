package __string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 默认初始化为0
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' // error: string是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) // 是byte数

	c := []rune(s) // 取出rune的切片
	t.Log(len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
