package isSymmetric

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeLeftString(root *TreeNode) (str string) {
	if root == nil {
		return ""
	}
	if root.Left != nil {
		str += TreeLeftString(root.Left)
	} else {
		str += "#"
	}
	if root.Right != nil {
		str += TreeLeftString(root.Right)
	} else {
		str += "#"//空节点认为是相同节点
	}
	str += fmt.Sprintf("%d", root.Val)
	return str
}
func TreeRightString(root *TreeNode) (str string) {
	if root == nil {
		return ""
	}
	if root.Right != nil {
		str += TreeRightString(root.Right)
	} else {
		str += "#"
	}
	if root.Left != nil {
		str += TreeRightString(root.Left)
	} else {
		str += "#"
	}
	str += fmt.Sprintf("%d", root.Val)
	return str
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return TreeLeftString(root.Left) == TreeRightString(root.Right)
}
