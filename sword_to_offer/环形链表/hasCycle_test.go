package hasCycle

import "testing"

func TestName(t *testing.T) {
	type Te struct {
		in  *ListNode
		out bool
	}
	b := ListNode{Val: 2}
	c := ListNode{Val: 0}
	d := ListNode{Val: -4}
	head := ListNode{Val: 3, Next: &b}
	b.Next = &c
	c.Next = &d
	d.Next = &b
	if out := hasCycle(&head); out != true {
		t.Fatal(out)
	}
}
