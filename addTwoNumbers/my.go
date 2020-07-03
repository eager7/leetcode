package addTwoNumbers

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (li *ListNode) Out() (list []int) {
	temp := li
	list = append(list, temp.Val)
	for temp.Next != nil {
		list = append(list, temp.Next.Val)
		temp = temp.Next
	}
	return list
}

func Iter() {
	li :=  &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	fmt.Println(li.Out())
	fmt.Println(li.Out())
}

/*
### 解题思路
这个题的坑在于单向链表不能去直接遍历，否则会导致链表头找不到，而且l1和l2不一定长度相等，不能放到一个循环中遍历，因此l1和l2遍历后要构造出非链表格式的数据出来，这里用map记录最长的链表，用结构体存储数据，然后去拼接。
注意最后一个进位问题，在这里遇到一次错误。
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	type add struct { //创建对象存储两个链表对应节点数值
		l1 int
		l2 int
	}
	var mList = make(map[int]add)          //用map解决链表长度不一致问题
	mList[0] = add{l1: l1.Val, l2: l2.Val} //初始化第一个值

	var i int
	for l1.Next != nil {
		i++
		mList[i] = add{l1: l1.Next.Val} //加入l1的值
		l1 = l1.Next
	}
	i = 0
	for l2.Next != nil {
		i++
		if v, ok := mList[i]; ok { //l1存在的加入，不存在的创建
			mList[i] = add{l1: v.l1, l2: l2.Next.Val}
		} else {
			mList[i] = add{l2: l2.Next.Val}
		}
		l2 = l2.Next
	}
	fmt.Println(mList)
	//组装新链表
	ln := new(ListNode)
	var up int //是否需要进位

	//用匿名函数传递链表对象，可以拷贝指针地址，否则会修改原有链表
	iter := func(temp *ListNode) {
		for i := 0; i < len(mList); i++ {
			if mList[i].l1+mList[i].l2+up >= 10 {
				temp.Val = (mList[i].l1 + mList[i].l2 + up) % 10
				if up == 0 { //处理第一次进来的情况
					up = 1
				}
			} else {
				temp.Val = mList[i].l1 + mList[i].l2 + up
				if up != 0 {
					up = 0 //进位功能完成，需清空
				}
			}
			if i+1 >= len(mList) { //跳出循环避免添加一个无效的节点
				if up == 1 { //向前进位，需要额外添加一个节点
					temp.Next = &ListNode{Val: 1}
				}
				break
			}
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}
	iter(ln) //传参，操作拷贝地址

	return ln
}
