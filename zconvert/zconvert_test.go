package zconvert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	type Te struct {
		In   string
		Rows int
		Out  string
	}

	list := []Te{
		{In: "AB", Rows: 1, Out: "AB"},
		{In: "LEETCODEISHIRING", Rows: 3, Out: "LCIRETOESIIGEDHN"},
		{In: "LEETCODEISHIRING", Rows: 4, Out: "LDREOEIIECIHNTSG"},
	}

	for _, l := range list {
		if out := best(l.In, l.Rows);out != l.Out {
			fmt.Println(l.Rows)
			t.Fatal(l.Out + "--" + out)
		}
	}
}
