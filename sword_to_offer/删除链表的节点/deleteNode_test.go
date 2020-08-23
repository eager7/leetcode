package deleteNode

import "testing"

func TestName(t *testing.T) {
	list := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	resp := deleteNode(list, 1)
	resp.List()
}
