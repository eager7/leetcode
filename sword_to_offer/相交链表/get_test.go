package getIntersectionNode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	head1 := ListNode{Val: 3}
	head2 := ListNode{Val: 2, Next: &head1}
	fmt.Println(getIntersectionNode(&head1, &head2))
}
