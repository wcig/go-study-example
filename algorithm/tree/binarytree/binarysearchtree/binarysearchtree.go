package binarysearchtree

import (
	"fmt"
	"go-app/algorithm/utils"
)

// 二叉搜索树
type BinarySearchTree struct {
	root       *Node
	comparator utils.Comparator
	size       int
}

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func New(c utils.Comparator) *BinarySearchTree {
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
	return bst.size == 0
}

func (bst *BinarySearchTree) Clear() {
	bst.root = nil
	bst.size = 0
}

func (bst *BinarySearchTree) Contains(e interface{}) bool {
	return bst.containsElement(bst.root, e)
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

func (bst *BinarySearchTree) FindMin() (interface{}, bool) {
	if node, exist := bst.findMinNode(bst.root); exist {
		return node.value, true
	}
	return nil, false
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

func (bst *BinarySearchTree) FindMax() (interface{}, bool) {
	if node, exist := bst.findMaxNode(bst.root); exist {
		return node.value, true
	}
	return nil, false
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

func (bst *BinarySearchTree) Add(v interface{}) {
	bst.root = bst.addElement(bst.root, v)
}

func (bst *BinarySearchTree) addElement(t *Node, v interface{}) *Node {
	if t == nil {
		bst.size++
		return &Node{v, nil, nil}
	}

	comp := bst.comparator(v, t.value)
	if comp < 0 {
		t.left = bst.addElement(t.left, v)
	} else if comp > 0 {
		t.right = bst.addElement(t.right, v)
	}
	return t
}

func (bst *BinarySearchTree) Remove(e interface{}) {
	if !bst.IsEmpty() {
		bst.root = bst.removeElement(bst.root, e)
	}
}

func (bst *BinarySearchTree) removeElement(t *Node, v interface{}) *Node {
	if t == nil {
		return nil
	}

	comp := bst.comparator(v, t.value)
	if comp < 0 {
		t.left = bst.removeElement(t.left, v)
	} else if comp > 0 {
		t.right = bst.removeElement(t.right, v)
	} else if t.left != nil && t.right != nil {
		minNode, _ := bst.findMinNode(t.right)
		min := minNode.value
		t.value = min
		t.right = bst.removeElement(t.right, min)
	} else {
		bst.size--
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

func (bst *BinarySearchTree) Depth() int {
	return bst.getDepth(bst.root)
}

func (bst *BinarySearchTree) getDepth(t *Node) int {
	if t == nil {
		return 0
	}
	l := bst.getDepth(t.left)
	r := bst.getDepth(t.right)
	return utils.MaxInt(l, r) + 1
}

// ------------------------------------------------------------------------ //

func (bst *BinarySearchTree) PreOrderTraverse(printFunc utils.Printer) {
	fmt.Print("preorder: ")
	bst.preOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) preOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		printFunc(t.value)
		bst.preOrderTraverseNode(t.left, printFunc)
		bst.preOrderTraverseNode(t.right, printFunc)
	}
}

func (bst *BinarySearchTree) InOrderTraverse(printFunc utils.Printer) {
	fmt.Print("inorder: ")
	bst.inOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) inOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		bst.inOrderTraverseNode(t.left, printFunc)
		printFunc(t.value)
		bst.inOrderTraverseNode(t.right, printFunc)
	}
}

func (bst *BinarySearchTree) PostOrderTraverse(printFunc utils.Printer) {
	fmt.Print("postorder: ")
	bst.postOrderTraverseNode(bst.root, printFunc)
	fmt.Println()
}

func (bst *BinarySearchTree) postOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		bst.postOrderTraverseNode(t.left, printFunc)
		bst.postOrderTraverseNode(t.right, printFunc)
		printFunc(t.value)
	}
}
