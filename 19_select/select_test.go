package _9_select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	//time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)  // channel
	retCh := make(chan string, 1) // buffered channel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	select {
	case ret := <- AsyncService():
		t.Log(ret)
	case <- time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
