package addTwoNumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}
	ln := AddTwoNumbers(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{0, 1}) {
		t.Fatal("must out [0,1]")
	}

	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 8}}
	l2 = &ListNode{Val: 0}
	ln = AddTwoNumbers(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{1, 8}) {
		t.Fatal("must out [1,8]")
	}

	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	ln = AddTwoNumbers(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{2, 4, 6}) {
		t.Fatal("must out [2,4,6]")
	}
}

func TestBest(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}
	ln := Best(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{0, 1}) {
		t.Fatal("must out [0,1]")
	}

	l1 = &ListNode{Val: 9, Next: &ListNode{Val: 8}}
	l2 = &ListNode{Val: 1}
	ln = Best(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{0, 9}) {
		t.Fatal("must out [0,9]")
	}

	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 8}}
	l2 = &ListNode{Val: 0}
	ln = Best(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{1, 8}) {
		t.Fatal("must out [1,8]")
	}

	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	ln = Best(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{2, 4, 6}) {
		t.Fatal("must out [2,4,6]")
	}

	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	ln = Best(l1, l2)
	fmt.Println(ln.Out())
	if !reflect.DeepEqual(ln.Out(), []int{7, 0, 8}) {
		t.Fatal("must out [7,0,8]")
	}
}

func TestIter(t *testing.T) {
	l :=  &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}

	fmt.Println(l.Out())
	fmt.Println(l.Out())
}