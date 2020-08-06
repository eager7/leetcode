package numWays

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  int
		Out int
	}
	var list = []Te{
		{In: 0, Out: 1},
		{In: 2, Out: 2},
		{In: 7, Out: 21},
		{In: 92, Out: 720754435},
	}
	for _, te := range list {
		if out := my(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		}
	}
}
