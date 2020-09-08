package list_copy

/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		if head.Random == nil {
			return &Node{Val: head.Val}
		} else {
			n := &Node{Val: head.Val}
			n.Random = n
			return n
		}
	}
	temp := head

	//复制节点为双节点
	for temp != nil {
		temp2 := &Node{
			Val:    temp.Val,
			Next:   temp.Next,
			Random: temp.Random,
		}
		temp, temp.Next = temp2.Next, temp2
	}
	for i, temp := 0, head; temp != nil; i, temp = i+1, temp.Next {
		if i%2 == 1 && temp.Random != nil {
			temp.Random = temp.Random.Next //将随机指针指向复制节点
		}
	}
	//拆分链表
	newHead := head.Next
	for i, temp := 0, head; temp != nil; i = i + 1 {
		next := temp.Next

		if i%2 == 0 && temp.Next != nil {
			temp.Next = temp.Next.Next
		}

		if i%2 == 1 && temp.Next != nil {
			temp.Next = temp.Next.Next
		}
		temp = next
	}
	return newHead
}
