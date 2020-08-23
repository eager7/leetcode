package myPow

import (
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	type Te struct {
		in1 float64
		in2 int
		out float64
	}
	tList := []Te{
		{in1: 0.00001, in2: 2147483647, out: 0},
		{in1: -2.00000, in2: 2, out: 4},
		{in1: 2.00000, in2: -2, out: 0.25000},
		{in1: 2.00000, in2: 10, out: 1024.00000},
		{in1: 2.10000, in2: 3, out: 9.26100},
	}
	for _, te := range tList {
		if out := myPow(te.in1, te.in2); !equal(out, te.out, 5) {
			t.Fatal(te.out, out)
		}
	}
}

func equal(x, y float64, decimals int) bool {
	min := math.Pow(10, -float64(decimals))
	fmt.Println(x, y, min, math.Dim(math.Max(x, y), math.Min(x, y)))
	return math.Dim(math.Max(x, y), math.Min(x, y)) <= min //差值比精度小说明相等
}

func TestEqual(t *testing.T) {
	fmt.Println(equal(0.9, 1.0, 1))
}
