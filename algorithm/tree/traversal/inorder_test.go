package main

import (
	"fmt"
	"testing"
)

// 中序遍历 (递归实现,DFS深度优先遍历)
func TestInorderTraversal(t *testing.T) {
	root := getRoot()
	var s []int
	inorder(root, &s)
	fmt.Println(">> preorder:", s) // >> inorder: [4 2 5 1 6 3 7]
}

func inorder(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, s)
	*s = append(*s, node.Val)
	inorder(node.Right, s)
}
