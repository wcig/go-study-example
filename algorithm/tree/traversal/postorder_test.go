package main

import (
	"fmt"
	"testing"
)

// 后序遍历 (递归实现,DFS深度优先遍历)
func TestPostorderTraversal(t *testing.T) {
	root := getRoot()
	var s []int
	postorder(root, &s)
	fmt.Println(">> postorder:", s) // >> postorder: [4 5 2 6 7 3 1]
}

func postorder(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, s)
	postorder(node.Right, s)
	*s = append(*s, node.Val)
}
