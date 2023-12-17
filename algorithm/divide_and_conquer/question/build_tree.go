// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package question

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构建二叉树: 分治
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 哈希表: 缓存inorder元素对应索引
	inorderMap := map[int]int{}
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return buildTreeDfs(preorder, inorder, inorderMap, 0, 0, len(preorder)-1)
}

// 构建二叉树
func buildTreeDfs(preorder []int, inorder []int, inorderMap map[int]int, i, l, r int) *TreeNode {
	// 区间为空即终止
	if l > r {
		return nil
	}
	// 初始化根节点
	rootVal := preorder[i]
	root := &TreeNode{Val: rootVal}
	// 定位根节点在inorder索引
	m := inorderMap[rootVal]
	// 子问题: 构建左子树
	root.Left = buildTreeDfs(preorder, inorder, inorderMap, i+1, l, m-1)
	// 子问题: 构建右子树
	root.Right = buildTreeDfs(preorder, inorder, inorderMap, i+1+m-l, m+1, r)
	// 返回根节点
	return root
}
