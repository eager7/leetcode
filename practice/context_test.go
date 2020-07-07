package practice

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	c1 := context.Background()
	c2 := context.TODO()
	c3 := context.WithValue(c1, "key", "val")
	c4, cancel := context.WithCancel(c1)
	c5, cancel := context.WithTimeout(c1, time.Second)

	fmt.Println(c1, c2, c3, c4, c5, cancel)
}
