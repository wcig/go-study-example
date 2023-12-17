package question

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 问题: 给定一棵二叉树的前序遍历 preorder 和中序遍历 inorder，请从中构建二叉树，返回二叉树的根节点。假设二叉树中没有值重复的节点。
func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 2, 1, 7}
	inorder := []int{9, 3, 1, 2, 7}
	fmt.Print("前序遍历 = ")
	fmt.Println(preorder)
	fmt.Print("中序遍历 = ")
	fmt.Println(inorder)

	root := buildTree(preorder, inorder)
	var preorderSlice, inorderSlice []int
	preorderTraversal(root, &preorderSlice)
	inorderTraversal(root, &inorderSlice)
	assert.Equal(t, preorder, preorderSlice)
	assert.Equal(t, inorder, inorderSlice)
}

func preorderTraversal(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	*s = append(*s, root.Val)
	preorderTraversal(root.Left, s)
	preorderTraversal(root.Right, s)
}

func inorderTraversal(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, s)
	*s = append(*s, root.Val)
	inorderTraversal(root.Right, s)
}
