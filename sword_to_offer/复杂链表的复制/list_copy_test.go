package list_copy

import "testing"

func TestName(t *testing.T) {
	a := &Node{Val: -1}
	a.Random = a
	copyRandomList(a)
}
