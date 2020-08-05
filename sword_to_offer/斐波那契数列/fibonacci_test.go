package fibonacci

import (
	"testing"
)

func TestName(t *testing.T) {
	if fib(2) != 1 {
		t.Fatal(2)
	}
	if fib(5) != 5 {
		t.Fatal(5)
	}
	if fib(45) != 134903163 {
		t.Fatal(fib(45))
	}
}
