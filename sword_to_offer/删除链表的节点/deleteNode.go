package deleteNode

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) List() {
	tmp := l
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func deleteNode(head *ListNode, val int) *ListNode {
	var pre, tmp = head, head
	for tmp != nil {
		if tmp.Val == val {
			if tmp == pre { //头结点则返回第二节点
				return tmp.Next
			}
			pre.Next = tmp.Next //跳过持有值的节点
			break
		}
		pre, tmp = tmp, tmp.Next //遍历
	}
	return head
}
