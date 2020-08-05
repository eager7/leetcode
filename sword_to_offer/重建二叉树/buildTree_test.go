package buildTree

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	fmt.Println("test example")
	type Te struct {
		In1 []int     //前序遍历
		In2 []int     //中序遍历
		Out *TreeNode //输出树
	}
	trie1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(trie1.List())
	t1 := Te{
		In1: []int{3, 9, 20, 15, 7},
		In2: []int{9, 3, 15, 20, 7},
		Out: trie1,
	}

	for _, te := range []Te{t1} {
		if out := my(te.In1, te.In2); !out.Equal(te.Out) {
			t.Fatal(te)
		}
	}
}
