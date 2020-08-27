package hasCycle

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	a, b := head, head
	for a.Next != nil && b.Next != nil && b.Next.Next != nil {
		if a == b.Next.Next {
			return true
		}
		a, b = a.Next, b.Next.Next //a走一步，b走两步
	}
	return false
}
