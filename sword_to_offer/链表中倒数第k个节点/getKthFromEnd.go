package getKthFromEnd

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	temp1, temp2 := head, head

	step := 0
	for temp1 != nil {
		step++
		temp1 = temp1.Next
		if step == k {
			break
		}
	}
	if step < k { //链表长度小于k
		return nil
	}
	for temp1 != nil && temp2 != nil {
		temp1, temp2 = temp1.Next, temp2.Next
	}
	return temp2
}
