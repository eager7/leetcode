package isSubStructure

import "testing"

func TestName(t *testing.T) {
	A := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	B := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: nil,
	}
	C := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	if isSubStructure(A, B) {
		t.Fatal(false)
	}
	if !isSubStructure(A, C) {
		t.Fatal(true)
	}
}
