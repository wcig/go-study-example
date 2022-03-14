package binarytreetraverse

import (
	"fmt"
	"testing"
)

// 二叉树3种遍历
type Tree struct {
	root *Node
}

type Node struct {
	value      interface{}
	leftChild  *Node
	rightChild *Node
}

// 先序遍历
func (t *Tree) PreOrderTraverse() {
	preOrderTraverse(t.root)
}

func preOrderTraverse(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value)
	preOrderTraverse(n.leftChild)
	preOrderTraverse(n.rightChild)
}

// 中序遍历
func (t *Tree) InOrderTraverse() {
	inOrderTraverse(t.root)
}

func inOrderTraverse(n *Node) {
	if n == nil {
		return
	}
	inOrderTraverse(n.leftChild)
	fmt.Print(n.value)
	inOrderTraverse(n.rightChild)
}

// 后序遍历
func (t *Tree) PostOrderTraverse() {
	postOrderTraverse(t.root)
}

func postOrderTraverse(n *Node) {
	if n == nil {
		return
	}
	postOrderTraverse(n.leftChild)
	postOrderTraverse(n.rightChild)
	fmt.Print(n.value)
}

func TestTraverse(t *testing.T) {
	fmt.Println("先序遍历")
	tree := genBinaryTree()
	tree.PreOrderTraverse()
	fmt.Println()

	fmt.Println("中序遍历")
	tree.InOrderTraverse()
	fmt.Println()

	fmt.Println("后序遍历")
	tree.PostOrderTraverse()
	fmt.Println()
	// Output:
	// 先序遍历
	// ABDGHCEIF
	// 中序遍历
	// GDHBAEICF
	// 后序遍历
	// GHDBIEFCA
}

//             A
//         B           C
//     D           E       F
// G       H           I
func genBinaryTree() *Tree {
	tree := &Tree{}

	// left child tree
	g := &Node{"G", nil, nil}
	h := &Node{"H", nil, nil}
	d := &Node{"D", g, h}
	b := &Node{"B", d, nil}

	// right child tree
	i := &Node{"I", nil, nil}
	e := &Node{"E", nil, i}
	f := &Node{"F", nil, nil}
	c := &Node{"C", e, f}

	a := &Node{"A", b, c}
	tree.root = a
	return tree
}
