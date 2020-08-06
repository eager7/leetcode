package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 14,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 16,
			},
		},
	}
	fmt.Println(trie.PreOrderList())
	fmt.Println(trie.InOrderList())
	fmt.Println(trie.BackOrderList())
	fmt.Println(trie.List())
}
