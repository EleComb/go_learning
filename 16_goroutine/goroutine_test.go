package _6_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func (i int) {
			fmt.Println(i)
		}(i)
		// error usage
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 50)
}
