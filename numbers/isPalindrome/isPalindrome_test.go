package isPalindrome

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  int
		Out bool
	}
	var list = []Te{
		{In: 121, Out: true},
		{In: -121, Out: false},
		{In: 10, Out: false},
	}
	for _, te := range list {
		if out := my(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		}
		if out := best(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		}
	}
}
