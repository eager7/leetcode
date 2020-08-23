package hammingWeight

import "testing"

func TestName(t *testing.T) {
	type Te struct {
		in  uint32
		out int
	}
	tList := []Te{
		{in: 1, out: 1},
		{in: 9, out: 2},
	}
	for _, te := range tList {
		if out := hammingWeight(te.in); out != te.out {
			t.Fatal(out, te.out)
		}
	}
}
