package cuttingRope

import "testing"

func TestName(t *testing.T) {
	type Te struct {
		in  int
		out int
	}
	tList := []Te{
		{in: 2, out: 1},
		{in: 10, out: 36},
	}
	for _, te := range tList {
		if out := cuttingRope(te.in); out != te.out {
			t.Fatal(te.out, out)
		}
	}
}
