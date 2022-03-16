package binarysearchtree

import "fmt"

// 二叉搜索树
type Comparator func(a, b interface{}) int

type PrintFunc func(v interface{})

type BinarySearchTree struct {
	root       *Node
	comparator Comparator
	size       int
}

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func New(c Comparator) *BinarySearchTree {
	return &BinarySearchTree{
		root:       nil,
		comparator: c,
		size:       0,
	}
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.root == nil || bst.size == 0
}

func (bst *BinarySearchTree) Clear() {
	bst.root = nil
	bst.size = 0
}

func (bst *BinarySearchTree) Contains(e interface{}) bool {
	return bst.containsElement(bst.root, e)
}

func (bst *BinarySearchTree) FindMin() (interface{}, bool) {
	if node, exist := bst.findMinNode(bst.root); exist {
		return node.value, true
	}
	return nil, false
}

func (bst *BinarySearchTree) FindMax() (interface{}, bool) {
	if node, exist := bst.findMaxNode(bst.root); exist {
		return node.value, true
	}
	return nil, false
}

func (bst *BinarySearchTree) Add(e interface{}) {
	newNode := &Node{e, nil, nil}
	if bst.IsEmpty() {
		bst.root = newNode
	} else {
		bst.addNode(bst.root, newNode)
	}
	bst.size++
}

func (bst *BinarySearchTree) Remove(e interface{}) {
	if !bst.IsEmpty() {
		bst.root = bst.removeElement(bst.root, e)
		bst.size--
	}
}

func (bst *BinarySearchTree) containsElement(t *Node, e interface{}) bool {
	if t == nil {
		return false
	}

	compareResult := bst.comparator(e, t.value)
	if compareResult < 0 {
		return bst.containsElement(t.left, e)
	} else if compareResult > 0 {
		return bst.containsElement(t.right, e)
	} else {
		return true
	}
}

func (bst *BinarySearchTree) findMinNode(e *Node) (*Node, bool) {
	if e == nil {
		return nil, false
	}
	if e.left == nil {
		return e, true
	}
	return bst.findMinNode(e.left)
}

func (bst *BinarySearchTree) findMaxNode(e *Node) (*Node, bool) {
	if e == nil {
		return nil, false
	}
	if e.right == nil {
		return e, true
	}
	return bst.findMaxNode(e.right)
}

func (bst *BinarySearchTree) addNode(t *Node, e *Node) {
	compareResult := bst.comparator(e.value, t.value)
	if compareResult < 0 {
		if t.left == nil {
			t.left = e
		} else {
			bst.addNode(t.left, e)
		}
	} else if compareResult > 0 {
		if t.right == nil {
			t.right = e
		} else {
			bst.addNode(t.right, e)
		}
	}
}

func (bst *BinarySearchTree) removeElement(t *Node, e interface{}) *Node {
	compareResult := bst.comparator(e, t.value)
	if compareResult < 0 {
		t.left = bst.removeElement(t.left, e)
	} else if compareResult > 0 {
		t.right = bst.removeElement(t.right, e)
	} else if t.left != nil && t.right != nil {
		minNode, _ := bst.findMinNode(t.right)
		min := minNode.value
		t.value = min
		t.right = bst.removeElement(t.right, min)
	} else {
		if t.left != nil {
			t = t.left
		} else if t.right != nil {
			t = t.right
		} else {
			t = nil
		}
	}
	return t
}

// ------------------------------------------------------------------------ //

func (bst *BinarySearchTree) PreOrderTraverse(printFunc PrintFunc) {
	fmt.Print("preorder: ")
	bst.preOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) preOrderTraverseNode(t *Node, printFunc PrintFunc) {
	if t != nil {
		printFunc(t.value)
		bst.preOrderTraverseNode(t.left, printFunc)
		bst.preOrderTraverseNode(t.right, printFunc)
	}
}

func (bst *BinarySearchTree) InOrderTraverse(printFunc PrintFunc) {
	fmt.Print("inorder: ")
	bst.inOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) inOrderTraverseNode(t *Node, printFunc PrintFunc) {
	if t != nil {
		bst.inOrderTraverseNode(t.left, printFunc)
		printFunc(t.value)
		bst.inOrderTraverseNode(t.right, printFunc)
	}
}

func (bst *BinarySearchTree) PostOrderTraverse(printFunc PrintFunc) {
	fmt.Print("postorder: ")
	bst.postOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) postOrderTraverseNode(t *Node, printFunc PrintFunc) {
	if t != nil {
		bst.postOrderTraverseNode(t.left, printFunc)
		bst.postOrderTraverseNode(t.right, printFunc)
		printFunc(t.value)
	}
}
