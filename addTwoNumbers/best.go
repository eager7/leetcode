package addTwoNumbers

import "fmt"

func Best(l1 *ListNode, l2 *ListNode) *ListNode {
	list := new(ListNode)
	temp := list //通过temp操作，保留原始list head用于返回
	var up int
	for {
		v := temp.Val + l1.Val + l2.Val
		if v >= 10 {
			v, up = v%10, 1 //需要向前进位
		}
		temp.Val = v

		if l1.Next == nil && l2.Next == nil { //结束后不要忘记检测是否还能进一位
			if up == 1 {
				temp.Next = &ListNode{Val: up} //向前进位
			}
			break
		}
		if l1.Next != nil { //迭代l1，注意，这里当两个链表长度不一致时，短链表值需置为0，否则会影响相加
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil { //迭代l2
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
		temp.Next = &ListNode{Val: up} //向前进位
		up, temp = 0, temp.Next        //将up清空
	}
	fmt.Printf("list:%p, temp:%p\n", list, temp)
	return list
}
