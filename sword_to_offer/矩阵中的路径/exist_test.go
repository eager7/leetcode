package exist

import "testing"

func TestName(t *testing.T) {
	type Te struct {
		in  string
		out bool
	}

	ar := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}

	tList := []Te{
		{in: "fce", out: true},
		{in: "afe", out: false},
		{in: "fca", out: false},
	}
	for _, te := range tList {
		if out := exist(ar, te.in); out != te.out {
			t.Fatal(out, te.out)
		}
	}
}
