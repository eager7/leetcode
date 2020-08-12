package min_array

import "testing"

func TestName(t *testing.T) {
	type Te struct {
		in  []int
		out int
	}
	tList := []Te{
		{in: []int{1, 3, 3, 3}, out: 1},
		{in: []int{3, 3, 1, 3}, out: 1},
		{in: []int{3, 1, 1}, out: 1},
		{in: []int{1, 3, 5}, out: 1},
		{in: []int{2, 2, 2, 0, 1}, out: 0},
		{in: []int{3, 4, 5, 1, 2}, out: 1},
		{in: []int{3, 1, 3}, out: 1},
	}
	for _, te := range tList {
		if out := best(te.in); out != te.out {
			t.Fatal(out, te.out)
		}
	}
}
