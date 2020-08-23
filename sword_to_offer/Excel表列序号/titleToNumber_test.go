package titleToNumber

import (
	"testing"
)

func TestName(t *testing.T) {
	type Te struct {
		in  string
		out int
	}
	tList := []Te{
		{in: "AB", out: 28},
		{in: "A", out: 1},
		{in: "AA", out: 27},
		{in: "ZY", out: 701},
	}
	for _, te := range tList {
		if out := titleToNumber(te.in); out != te.out {
			t.Fatal(te.out, out)
		}
	}
}
