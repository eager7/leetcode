package uniquePaths

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  int
		In2 int
		Out int
	}
	var list = []Te{
		{In: 3, In2: 2, Out: 3},
		{In: 7, In2: 3, Out: 28},
	}
	for _, te := range list {
		if out := my(te.In, te.In2); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		}
	}
}
