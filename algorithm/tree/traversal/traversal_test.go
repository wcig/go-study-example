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

// 层序遍历 (辅助队列实现,BFS深度优先遍历)
func TestLevelOrderTraversal(t *testing.T) {
	root := getRoot()
	var s []int
	levelOrder(root, &s)
	fmt.Println(">> levelOrder:", s) // >> levelOrder: [1 2 3 4 5 6 7]
}

func levelOrder(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		node := front.Value.(*TreeNode)
		*s = append(*s, node.Val)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
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
