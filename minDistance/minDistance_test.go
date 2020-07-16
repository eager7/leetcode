package minDistance

import "testing"

func TestExample(t *testing.T) {
	type Te struct {
		In1 string
		In2 string
		Out int
	}

	tList := []Te{
		{In1: "distance", In2: "springbok", Out: 9},
		{In1: "a", In2: "b", Out: 1},
		{In1: "", In2: "", Out: 0},
		{In1: "", In2: "pct", Out: 3},
		{In1: "horse", In2: "ros", Out: 3},
		{In1: "intention", In2: "execution", Out: 5},
	}
	for _, te := range tList {
		if out := my(te.In1, te.In2); out != te.Out {
			t.Fatal(te, out)
		}
	}
}
