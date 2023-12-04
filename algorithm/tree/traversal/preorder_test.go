package main

import (
	"fmt"
	"testing"
)

// 前序遍历 (递归实现,DFS深度优先遍历) 根左右
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

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		*s = append(*s, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

// 前序遍历 (辅助栈实现-统一迭代法)
func TestPreorderTraversalWithStack2(t *testing.T) {
	root := getRoot()
	var s []int
	preorderWithStack2(root, &s)
	fmt.Println(">> preorderWithStack2:", s) // >> preorderWithStack: [1 2 4 5 3 6 7]
}

func preorderWithStack2(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node != nil {
			// pop
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			// pop
			stack = stack[:len(stack)-1]
			// pop
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*s = append(*s, node.Val)
		}
	}
}
