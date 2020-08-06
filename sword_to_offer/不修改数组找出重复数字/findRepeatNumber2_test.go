package findRepeatNumber

import "testing"

func TestExample(t *testing.T) {
	type Te struct {
		In  []int
		Out int
	}
	tList := []Te{
		{In: []int{2, 3, 5, 4, 3, 2, 6, 7}, Out: 3},
		{In: []int{5, 5, 5, 5, 5, 5, 5}, Out: 5},
		{In: []int{3, 2, 1, 3, 4}, Out: 3},
		{In: []int{1, 1}, Out: 1},
	}

	for _, te := range tList {
		if out := my(te.In); out != te.Out {
			t.Fatal(te, out)
		}
	}
}
