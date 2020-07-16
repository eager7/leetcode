package minPathSum

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  [][]int
		Out int
	}
	var in = [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	var list = []Te{
		{In: in, Out: 7},
	}
	for _, te := range list {
		if out := my(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		}
	}
}