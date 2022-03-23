package threadedbinarytree

import "fmt"

// 线索二叉树
type ThreadedBinaryTree struct {
	root *Node
}

type Node struct {
	value      interface{} // 结点元素值
	leftChild  *Node       // 左孩子或前驱
	rightChild *Node       // 右孩子或后继
	leftTag    int         // 值为0表示左孩子,1表示前驱
	rightTag   int         // 值为0表示右孩子,1表示后继
}

// 前序线索化二叉树
func (t *ThreadedBinaryTree) CreatePreThread() {
	createPreThread(t.root)
}

func createPreThread(e *Node) {
	pre := new(*Node)
	if e != nil {
		// 前序线索化
		preThread(e, pre)
		// 处理最后一个遍历结点
		if (*pre).rightChild == nil {
			(*pre).rightTag = 1
		}
	}
}

func preThread(e *Node, pre **Node) {
	if e != nil {
		if e.leftChild == nil {
			e.leftChild = *pre
			e.leftTag = 1
		}
		if pre != nil && *pre != nil && (*pre).rightChild == nil {
			(*pre).rightChild = e
			(*pre).rightTag = 1
		}
		*pre = e
		if e.leftTag == 0 {
			preThread(e.leftChild, pre)
		}
		if e.rightTag == 0 {
			preThread(e.rightChild, pre)
		}
	}
}

// 中序线索化二叉树
func (t *ThreadedBinaryTree) CreateInThread() {
	createInThread(t.root)
}

func createInThread(e *Node) {
	pre := new(*Node)
	if e != nil {
		// 中序线索化
		inThread(e, pre)
		// 处理最后一个遍历结点
		(*pre).rightChild = nil
		(*pre).rightTag = 1
	}
}

// 中序遍历方式线索化二叉树
func inThread(e *Node, pre **Node) {
	if e != nil {
		inThread(e.leftChild, pre)
		// 若结点左孩子为空,则其前驱为前一个遍历结点
		if e.leftChild == nil {
			if pre != nil {
				e.leftChild = *pre
			} else {
				e.leftChild = nil
			}
			e.leftTag = 1
		}
		// 若前一个遍历结点不为空且其右孩子为空,则其后继为之后遍历的当前结点
		if pre != nil && *pre != nil && (*pre).rightChild == nil {
			(*pre).rightChild = e
			(*pre).rightTag = 1
		}
		*pre = e
		inThread(e.rightChild, pre)
	}
}

// 后序线索化二叉树
func (t *ThreadedBinaryTree) CreatePostThread() {
	createPostThread(t.root)
}

func createPostThread(e *Node) {
	pre := new(*Node)
	if e != nil {
		// 后序线索化
		postThread(e, pre)
		// 处理最后一个遍历结点
		if (*pre).rightChild == nil {
			(*pre).rightTag = 1
		}
	}
}

func postThread(e *Node, pre **Node) {
	if e != nil {
		postThread(e.leftChild, pre)
		postThread(e.rightChild, pre)
		if e.leftChild == nil {
			e.leftChild = *pre
			e.leftTag = 1
		}
		if pre != nil && *pre != nil && (*pre).rightChild == nil {
			(*pre).rightChild = e
			(*pre).rightTag = 1
		}
		*pre = e
	}
}

// 对中序线索二叉树进行中序遍历 (前驱方式)
func (t *ThreadedBinaryTree) RevInOrder() {
	for e := inOrderLastNode(t.root); e != nil; e = inOrderPrevNode(e) {
		fmt.Print(e.value)
	}
	fmt.Println()
}

func inOrderLastNode(e *Node) *Node {
	if e == nil {
		return nil
	}

	for e.rightTag == 0 {
		e = e.rightChild
	}
	return e
}

func inOrderPrevNode(e *Node) *Node {
	if e == nil {
		return nil
	}

	if e.leftTag == 0 {
		return inOrderLastNode(e.leftChild)
	}
	return e.leftChild
}

// 对中序线索二叉树进行中序遍历 (后继方式)
func (t *ThreadedBinaryTree) InOrder() {
	for e := inOrderFirstNode(t.root); e != nil; e = inOrderNextNode(e) {
		fmt.Print(e.value)
	}
	fmt.Println()
}

// 中序遍历第一个结点
func inOrderFirstNode(e *Node) *Node {
	if e == nil {
		return nil
	}

	for e.leftTag == 0 {
		e = e.leftChild
	}
	return e
}

// 中序遍历后继结点
func inOrderNextNode(e *Node) *Node {
	if e == nil {
		return nil
	}

	if e.rightTag == 0 {
		return inOrderFirstNode(e.rightChild)
	}
	return e.rightChild
}

// ------------------------------------------------------------- //

// 先序遍历
func (t *ThreadedBinaryTree) PreOrderTraverse() {
	preOrderTraverse(t.root)
	fmt.Println()
}

func preOrderTraverse(e *Node) {
	if e == nil {
		return
	}
	fmt.Print(e.value)
	preOrderTraverse(e.leftChild)
	preOrderTraverse(e.rightChild)
}

// 中序遍历
func (t *ThreadedBinaryTree) InOrderTraverse() {
	inOrderTraverse(t.root)
	fmt.Println()
}

func inOrderTraverse(e *Node) {
	if e == nil {
		return
	}
	inOrderTraverse(e.leftChild)
	fmt.Print(e.value)
	inOrderTraverse(e.rightChild)
}

// 后序遍历
func (t *ThreadedBinaryTree) PostOrderTraverse() {
	postOrderTraverse(t.root)
	fmt.Println()
}

func postOrderTraverse(e *Node) {
	if e == nil {
		return
	}
	postOrderTraverse(e.leftChild)
	postOrderTraverse(e.rightChild)
	fmt.Print(e.value)
}
