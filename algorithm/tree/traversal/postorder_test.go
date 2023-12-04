package main

import (
	"fmt"
	"testing"
)

// 后序遍历 (递归实现,DFS深度优先遍历) 左右根
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

// 后序遍历 (辅助栈实现)
func TestPostorderTraversalWithStack(t *testing.T) {
	root := getRoot()
	var s []int
	postorderWithStack(root, &s)
	fmt.Println(">> postorderWithStack:", s) // >> postorderWithStack: [4 5 2 6 7 3 1]
}

func postorderWithStack(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*s = append(*s, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse(s)
}

func reverse(s *[]int) {
	n := len(*s)
	if n == 0 {
		return
	}
	l, r := 0, n-1
	for l < r {
		(*s)[l], (*s)[r] = (*s)[r], (*s)[l]
		l++
		r--
	}
}
