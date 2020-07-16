package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	t1 := time.NewTicker(time.Second)
	t2 := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-t1.C:
			fmt.Println("t1", time.Now())
			t1.Stop()
			time.Sleep(time.Second * 6)
		case <-t2.C://没有被选中执行也不会丢失
			fmt.Println("t2", time.Now())
		}
	}
}

func TestSendChannel(t *testing.T) {
	var channel chan string
	channel = make(chan string)
	fmt.Println(<-channel)
	channel <- "pct"
}