package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	m := sync.Map{}
	m.Store("pct", "hello")
	fmt.Println(m.Load("pct"))
	fmt.Println(m.Load("pct2"))
}
