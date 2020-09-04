package mirrorTree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func PrintTree(tree *TreeNode)(list []int) {
	queue := []*TreeNode{tree}//初始化队列，用队列按层级打印
	for len(queue)!=0{
		temp, queue:=queue[0],queue[1:]
		if temp.Left!=nil{
			queue=append(queue,temp.Left)
		}
		if temp.Right!=nil{
			queue=append(queue,temp.Right)
		}
		list = append(list, temp.Val)
	}
	return list
}

func MirrorTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Left!= nil {
		MirrorTree(tree.Left)
	}
	if tree.Right!=nil{
		MirrorTree(tree.Right)
	}
	tree.Left,tree.Right=tree.Right,tree.Left
}

func mirrorTree(root *TreeNode) *TreeNode {
	MirrorTree(root)
	return root
}