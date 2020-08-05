package sort

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	type Te struct {
		In  []int
		Out []int
	}
	tList := []Te{
		{In: []int{2, 4, 3, 6, 5, 8, 9, 7, 1}, Out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{In: []int{2, 4, 3, 3, 4, 1}, Out: []int{1, 2, 3, 3, 4, 4}},
	}
	for _, te := range tList {
		if out := QuickSort(te.In); !reflect.DeepEqual(te.Out, out) {
			t.Fatal(out)
		}
	}
}
