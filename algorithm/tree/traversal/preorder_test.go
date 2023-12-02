package main

import (
	"container/list"
	"fmt"
	"testing"
)

// 前序遍历 (递归实现,DFS深度优先遍历)
func TestPreorderTraversal(t *testing.T) {
	root := getRoot()
	var s []int
	preorder(root, &s)
	fmt.Println(">> preorder:", s) // >> preorder: [1 2 4 5 3 6 7]
}

func preorder(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	*s = append(*s, node.Val)
	preorder(node.Left, s)
	preorder(node.Right, s)
}

// 前序遍历 (辅助栈实现,DFS深度优先遍历)
func TestPreorderTraversalWithStack(t *testing.T) {
	root := getRoot()
	var s []int
	preorderWithStack(root, &s)
	fmt.Println(">> preorderWithStack:", s) // >> preorderWithStack: [1 2 4 5 3 6 7]
}

func preorderWithStack(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		back := stack.Back()
		stack.Remove(back)

		node := back.Value.(*TreeNode)
		*s = append(*s, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
}
