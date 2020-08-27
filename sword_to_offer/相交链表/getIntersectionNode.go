package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tmp1, tmp2 := headA, headB
	for {
		tmp2 = headB
		for {
			if tmp1 == tmp2 {
				return tmp1
			}
			if tmp2.Next == nil {
				break
			}
			tmp2 = tmp2.Next
		}
		if tmp1.Next == nil {
			break
		}
		tmp1 = tmp1.Next
	}
	return nil
}
