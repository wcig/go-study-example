// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_backtracking

package question

import (
	"fmt"
	"testing"
)

func TestPreorderTraversalIIITemplate(t *testing.T) {
	/* 初始化二叉树 */
	// root := SliceToTree([]any{1, 7, 3, 4, 5, 6, 7})
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 7}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	n2.Left = &TreeNode{Val: 4}
	n2.Right = &TreeNode{Val: 5}
	n3.Left = &TreeNode{Val: 6}
	n3.Right = &TreeNode{Val: 7}
	root := n1
	fmt.Println("\n初始化二叉树")

	// 回溯算法
	res := make([][]*TreeNode, 0)
	state := make([]*TreeNode, 0)
	choices := make([]*TreeNode, 0)
	choices = append(choices, root)
	backtrackIII(&state, &choices, &res)

	fmt.Println("\n输出所有根节点到节点 7 的路径，路径中不包含值为 3 的节点")
	for _, path := range res {
		for _, node := range path {
			fmt.Printf("%v ", node.Val)
		}
		fmt.Println()
	}
}
