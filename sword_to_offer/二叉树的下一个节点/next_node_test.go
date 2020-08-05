package next_node

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := &Trie{val: "a"}
	b := &Trie{val: "b", parent: a}
	c := &Trie{val: "c", parent: a}
	d := &Trie{val: "d", parent: b}
	e := &Trie{val: "e", parent: b}
	f := &Trie{val: "f", parent: c}
	g := &Trie{val: "g", parent: c}
	h := &Trie{val: "h", parent: e}
	i := &Trie{val: "i", parent: e}

	a.left, a.right = b, c
	b.left, b.right = d, e
	e.left, e.right = h, i
	c.left, c.right = f, g

	fmt.Println(a.InOrderList())

	type Te struct {
		In  *Trie
		Out string
	}
	tList := []Te{
		{In: d, Out: "b"},
		{In: i, Out: "a"},
		{In: c, Out: "g"},
		{In: g, Out: ""},
	}
	for _, te := range tList {
		if out := my(te.In); out != te.Out {
			t.Fatal(out)
		}
	}
}
