package main

import (
	"fmt"
	"testing"
)

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

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		*s = append(*s, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// 层序遍历,获取每层数组
func TestLevelOrderSlice(t *testing.T) {
	root := getRoot()
	result := [][]int{}
	levelOrderSlice(root, &result, 0)
	fmt.Println(">> levelOrderSlice:", result) // >> levelOrderSlice: [[1] [2 3] [4 5 6 7]]
}

func levelOrderSlice(node *TreeNode, s *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*s) < level+1 {
		*s = append(*s, []int{node.Val})
	} else {
		(*s)[level] = append((*s)[level], node.Val)
	}
	levelOrderSlice(node.Left, s, level+1)
	levelOrderSlice(node.Right, s, level+1)
}
