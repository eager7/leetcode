package findRepeatNumber

import "testing"

func TestExample(t *testing.T) {
	type Te struct {
		In  []int
		Out int
	}
	tList := []Te{
		{In: []int{2, 3, 1, 0, 2, 5, 3}, Out: 2},
		{In: []int{1, 1}, Out: 1},
	}

	for _, te := range tList {
		if out := my(te.In); out != te.Out {
			t.Fatal(te, out)
		}
	}
}
