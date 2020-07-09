package isMatch

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  string
		In2 string
		Out bool
	}
	var list = []Te{
		{In: "aaa", In2: "a*a", Out: true},
		{In: "mississippi", In2: "mis*is*ip*.", Out: true},
		{In: "aa", In2: "a", Out: false},
		{In: "aa", In2: "a*", Out: true},
		{In: "ab", In2: ".**", Out: true},
		{In: "aab", In2: "c*a*b", Out: true},
		{In: "mississippi", In2: "mis*is*p*.", Out: false},
	}
	for _, te := range list {
		if out := my(te.In, te.In2); out != te.Out {
			t.Fatal(fmt.Sprintf("in:%v,in2:%v, out:%v, err:%v", te.In, te.In2, te.Out, out))
		}
		//if out := best(te.In, te.In2); out != te.Out {
		//	t.Fatal(fmt.Sprintf("in:%v, out:%v, err:%v", te.In, te.Out, out))
		//}
	}
}
