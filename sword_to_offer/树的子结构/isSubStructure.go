package isSubStructure

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//用中序遍历查找B在A中的位置
func Search(A, B *TreeNode) bool {
	if A.Left != nil {
		if A.Left.Val == B.Val {
			if Equal(A, B) {
				return true
			}
		}
		if Search(A.Left, B) {
			return true
		}
	}
	if A.Val == B.Val {
		if Equal(A, B) {
			return true
		}
	}
	if A.Right != nil {
		if A.Right.Val == B.Val {
			if Equal(A, B) {
				return true
			}
		}
		if Search(A.Right, B) {
			return true
		}
	}
	return false
}

//判断B是不是A的子树
func Equal(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if B.Left != nil { //B有节点，A没有节点则不相同
		if !Equal(A.Left, B.Left) {
			return false
		}
	}
	if B.Right != nil { //B有节点，A没有节点则不相同
		if !Equal(A.Right, B.Right) {
			return false
		}
	}
	return A.Val == B.Val //比较当前节点是否相同
}

/*
** 根据树的左右遍历顺序判断
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return Search(A, B)
}
