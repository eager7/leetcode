package trie

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历，左根右
func (t *TreeNode) InOrderList() (list []int) {
	if t.Left != nil {
		list = append(list, t.Left.InOrderList()...)
	}
	list = append(list, t.Val)
	if t.Right != nil {
		list = append(list, t.Right.InOrderList()...)
	}
	return list
}

//前序遍历，根左右
func (t *TreeNode) PreOrderList() (list []int) {
	list = append(list, t.Val)
	if t.Left != nil {
		list = append(list, t.Left.PreOrderList()...)
	}
	if t.Right != nil {
		list = append(list, t.Right.PreOrderList()...)
	}
	return list
}

//后序遍历，左右根
func (t *TreeNode) BackOrderList() (list []int) {
	if t.Left != nil {
		list = append(list, t.Left.InOrderList()...)
	}
	if t.Right != nil {
		list = append(list, t.Right.InOrderList()...)
	}
	list = append(list, t.Val)
	return list
}

//深度优先遍历，自上而下
func (t *TreeNode) List() (list []int) {
	var queue []*TreeNode
	queue = append(queue, t)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		list = append(list, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return list
}
