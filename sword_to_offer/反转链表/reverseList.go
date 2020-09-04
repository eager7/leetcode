package reverseList

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var latest, tail *ListNode
	for head != nil {
		tail = &ListNode{Val:  head.Val, Next: latest}//第一个节点在新链表尾部，是空
		latest = tail
		head = head.Next
	}
	return tail
}
