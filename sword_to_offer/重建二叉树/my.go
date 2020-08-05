package buildTree

import "reflect"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) List() (list []int) {
	list = append(list, t.Val)
	if t.Left != nil {
		list = append(list, t.Left.List()...)
	}
	if t.Right != nil {
		list = append(list, t.Right.List()...)
	}
	return list
}
func (t *TreeNode) Equal(in *TreeNode) bool {
	tList := t.List()
	inList := in.List()
	return reflect.DeepEqual(tList, inList)
}

func my(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := preorder[0]                 //找到根节点，在中序遍历里，根节点处于数组中间位置
	var leftInOrder, rightInOrder []int //中序遍历根据根节点进行分组
	for index, val := range inorder {
		if val == root { //中序遍历中，根节点处于中间，两边分别是左右树
			leftInOrder, rightInOrder = inorder[:index], inorder[index+1:]
			break
		}
	}
	//找出分组后的前序遍历数组，进行迭代
	leftPreOrder, rightPreOrder := preorder[1:1+len(leftInOrder)], preorder[1+len(leftInOrder):]

	return &TreeNode{
		Val:   root,
		Left:  my(leftPreOrder, leftInOrder),
		Right: my(rightPreOrder, rightInOrder),
	}
}
