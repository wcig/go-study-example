package main

import (
	"fmt"
	"testing"
)

// 中序遍历 (递归实现,DFS深度优先遍历) 左根右
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

// 中序遍历 (辅助栈实现)
func TestInorderTraversalWithStack(t *testing.T) {
	root := getRoot()
	var s []int
	inorderWithStack(root, &s)
	fmt.Println(">> inorderWithStack:", s) // >> inorderWithStack: [4 2 5 1 6 3 7]
}

func inorderWithStack(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*s = append(*s, cur.Val)
			cur = cur.Right
		}
	}
}

// 中序遍历 (辅助栈实现-统一迭代法)
func TestInorderTraversalWithStack2(t *testing.T) {
	root := getRoot()
	var s []int
	inorderWithStack2(root, &s)
	fmt.Println(">> inorderWithStack2:", s) // >> inorderWithStack2: [4 2 5 1 6 3 7]
}

func inorderWithStack2(root *TreeNode, s *[]int) {
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
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
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
