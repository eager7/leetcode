package mergeTwoLists

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var list = &ListNode{}
	temp, temp1, temp2 := list, l1, l2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = &ListNode{Val: temp1.Val, Next: nil}
			temp, temp1 = temp.Next, temp1.Next
		} else {
			temp.Next = &ListNode{Val: temp2.Val, Next: nil}
			temp, temp2 = temp.Next, temp2.Next
		}
	}
	for temp1 != nil {
		temp.Next = &ListNode{Val: temp1.Val, Next: nil}
		temp, temp1 = temp.Next, temp1.Next
	}
	for temp2 != nil {
		temp.Next = &ListNode{Val: temp2.Val, Next: nil}
		temp, temp2 = temp.Next, temp2.Next
	}
	return list.Next
}
