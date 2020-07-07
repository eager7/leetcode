package nreverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	type Te struct {
		In  int
		Out int
	}
	var list = []Te{
		{In: 123, Out: 321},
		{In: -123, Out: -321},
		{In: 1534236469, Out: 0},
		{In: -2147483412, Out: -2143847412},
	}
	for _, te := range list {
		if out := my(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("out:%d, err:%d", te.Out, out))
		}
		if out := best(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("out:%d, err:%d", te.Out, out))
		}
	}
}
