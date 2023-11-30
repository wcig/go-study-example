package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	    1
//	  /   \
//	 2     3
//	/ \   / \
//
// 4   5 6   7
func getRoot() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	return node1
}
