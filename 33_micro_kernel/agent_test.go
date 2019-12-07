package _3_micro_kernel

import (
	"fmt"
	"testing"
	"time"
)

func TestAgent(t *testing.T) {
	agentObj := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	_ = agentObj.RegisterCollector("c1", c1)
	_ = agentObj.RegisterCollector("c2", c2)
	if err := agentObj.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	fmt.Println("agentStart:", agentObj.Start())
	time.Sleep(time.Second * 1)
	_ = agentObj.Stop()
	_ = agentObj.Destory()
}
