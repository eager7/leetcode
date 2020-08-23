package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var filter = make(map[int]struct{})
	del := func() {
		var pre, tmp = head, head
		for tmp != nil {
			if _, ok := filter[tmp.Val]; ok {
				if pre == tmp {
					pre, head, tmp = tmp.Next, tmp.Next, tmp.Next //删除头节点
				} else {
					pre.Next = tmp.Next //删除tmp节点
					tmp = tmp.Next
				}
			} else {
				filter[tmp.Val] = struct{}{}
				pre, tmp = tmp, tmp.Next
			}
		}
	}
	del()
	del()
	return head
}
