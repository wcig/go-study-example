package avltree

import (
	"fmt"
	"go-app/algorithm/utils"
)

// AVL树
const (
	allowedImbalance = 1
)

type AVLTree struct {
	root       *Node
	comparator utils.Comparator
	size       int
}

type Node struct {
	value  interface{}
	left   *Node
	right  *Node
	height int
}

func height(n *Node) int {
	if n != nil {
		return n.height
	}
	return -1
}

func New(c utils.Comparator) *AVLTree {
	return &AVLTree{
		root:       nil,
		comparator: c,
		size:       0,
	}
}

func (at *AVLTree) Size() int {
	return at.size
}

func (at *AVLTree) IsEmpty() bool {
	return at.size == 0
}

func (at *AVLTree) Clear() {
	at.root = nil
	at.size = 0
}

func (at *AVLTree) Contains(e interface{}) bool {
	return at.containsElement(at.root, e)
}

func (at *AVLTree) containsElement(t *Node, v interface{}) bool {
	if t == nil {
		return false
	}

	compareResult := at.comparator(v, t.value)
	if compareResult < 0 {
		return at.containsElement(t.left, v)
	} else if compareResult > 0 {
		return at.containsElement(t.right, v)
	} else {
		return true
	}
}

func (at *AVLTree) FindMin() (interface{}, bool) {
	if node, exist := at.findMinNode(at.root); exist {
		return node.value, true
	}
	return nil, false
}

func (at *AVLTree) findMinNode(e *Node) (*Node, bool) {
	if e == nil {
		return nil, false
	}
	if e.left == nil {
		return e, true
	}
	return at.findMinNode(e.left)
}

func (at *AVLTree) FindMax() (interface{}, bool) {
	if node, exist := at.findMaxNode(at.root); exist {
		return node.value, true
	}
	return nil, false
}

func (at *AVLTree) findMaxNode(e *Node) (*Node, bool) {
	if e == nil {
		return nil, false
	}
	if e.right == nil {
		return e, true
	}
	return at.findMaxNode(e.right)
}

func (at *AVLTree) Add(v interface{}) {
	at.root = at.addElement(at.root, v)
}

func (at *AVLTree) addElement(t *Node, v interface{}) *Node {
	if t == nil {
		at.size++
		return &Node{v, nil, nil, 0}
	}

	cmp := at.comparator(v, t.value)
	if cmp < 0 {
		t.left = at.addElement(t.left, v)
	} else if cmp > 0 {
		t.right = at.addElement(t.right, v)
	}
	return at.balance(t)
}

func (at *AVLTree) Remove(v interface{}) {
	if !at.IsEmpty() {
		at.root = at.removeElement(at.root, v)
	}
}

func (at *AVLTree) removeElement(t *Node, v interface{}) *Node {
	if t == nil {
		return t
	}

	cmp := at.comparator(v, t.value)
	if cmp < 0 {
		t.left = at.removeElement(t.left, v)
	} else if cmp > 0 {
		t.right = at.removeElement(t.right, v)
	} else if t.left != nil && t.right != nil {
		minNode, _ := at.findMinNode(t.right)
		min := minNode.value
		t.value = min
		t.right = at.removeElement(t.right, min)
	} else {
		at.size--
		if t.left != nil {
			t = t.left
		} else if t.right != nil {
			t = t.right
		} else {
			t = nil
		}
	}
	return at.balance(t)
}

func (at *AVLTree) balance(t *Node) *Node {
	if t == nil {
		return t
	}

	imbalance := height(t.left) - height(t.right)
	if imbalance > allowedImbalance {
		if height(t.left.left) >= height(t.left.right) {
			t = at.rightRotation(t)
		} else {
			t = at.leftRightRotation(t)
		}
	} else if -imbalance > allowedImbalance {
		if height(t.right.right) >= height(t.right.left) {
			t = at.leftRotation(t)
		} else {
			t = at.rightLeftRotation(t)
		}
	}

	t.height = utils.MaxInt(height(t.left), height(t.right)) + 1
	return t
}

func (at *AVLTree) rightRotation(k *Node) *Node {
	kl := k.left
	k.left = kl.right
	kl.right = k
	k.height = utils.MaxInt(height(k.left), height(k.right)) + 1
	kl.height = utils.MaxInt(height(kl.left), k.height) + 1
	return kl
}

func (at *AVLTree) leftRotation(k *Node) *Node {
	kr := k.right
	k.right = kr.left
	kr.left = k
	k.height = utils.MaxInt(height(k.left), height(k.right)) + 1
	kr.height = utils.MaxInt(k.height, height(kr.right)) + 1
	return kr
}

func (at *AVLTree) leftRightRotation(k3 *Node) *Node {
	k3.left = at.leftRotation(k3.left)
	return at.rightRotation(k3)
}

func (at *AVLTree) rightLeftRotation(k3 *Node) *Node {
	k3.right = at.rightRotation(k3.right)
	return at.leftRotation(k3)
}

// ------------------------------------------------------------------------ //

func (at *AVLTree) PreOrderTraverse(printFunc utils.Printer) {
	fmt.Print("preorder: ")
	at.preOrderTraverseNode(at.root, printFunc)
	fmt.Println()
}

func (at *AVLTree) preOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		printFunc(t.value)
		at.preOrderTraverseNode(t.left, printFunc)
		at.preOrderTraverseNode(t.right, printFunc)
	}
}

func (at *AVLTree) InOrderTraverse(printFunc utils.Printer) {
	fmt.Print("inorder: ")
	at.inOrderTraverseNode(at.root, printFunc)
	fmt.Println()
}

func (at *AVLTree) inOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		at.inOrderTraverseNode(t.left, printFunc)
		printFunc(t.value)
		at.inOrderTraverseNode(t.right, printFunc)
	}
}

func (at *AVLTree) PostOrderTraverse(printFunc utils.Printer) {
	fmt.Print("postorder: ")
	at.postOrderTraverseNode(at.root, printFunc)
	fmt.Println()
}

func (at *AVLTree) postOrderTraverseNode(t *Node, printFunc utils.Printer) {
	if t != nil {
		at.postOrderTraverseNode(t.left, printFunc)
		at.postOrderTraverseNode(t.right, printFunc)
		printFunc(t.value)
	}
}
