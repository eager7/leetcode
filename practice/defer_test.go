package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	time.Sleep(time.Second)
}
