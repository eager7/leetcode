package search

import "testing"

type Te struct {
	In1 []int
	In2 int
	Out bool
}

var TeList []Te

func TestName(t *testing.T) {
	for _, te := range TeList {
		if BinarySearch(te.In1, te.In2) != te.Out {
			t.Fatal(te)
		}
	}
}

func init() {
	t1 := Te{
		In1: []int{1, 2, 3, 4, 5},
		In2: 1, Out: true}
	TeList = append(TeList, t1)
	t2 := Te{
		In1: []int{1, 2, 3, 4, 5},
		In2: 7, Out: false}
	TeList = append(TeList, t2)
}
