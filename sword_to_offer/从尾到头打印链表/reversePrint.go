package reversePrint

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var list []int
	var temp = head
	for temp != nil {
		list = append([]int{temp.Val}, list...)
		temp = temp.Next
	}
	return list
}
