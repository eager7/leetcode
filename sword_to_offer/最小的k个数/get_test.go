package getLeastNumbers

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	type Te struct {
		in  []int
		out int
	}
	tList := []Te{
		{in: []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, out: 8},
		{in: []int{3, 2, 1}, out: 2},
		{in: []int{0, 0, 0, 1, 2, 2, 3, 7, 6, 1}, out: 3},
		{in: []int{0, 0, 1, 3, 4, 5, 0, 7, 6, 7}, out: 9},
	}
	for _, te := range tList {
		ol := getLeastNumbers(te.in, te.out)
		if out := len(ol); out != te.out {
			fmt.Println(te.in)
			t.Fatal(out)
		} else {
			fmt.Println(ol)
		}
	}
}
