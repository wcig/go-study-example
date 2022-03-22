package binarytreetraverse

import (
	"fmt"
	"go-app/algorithm/linearlist/queue/simplequeue"
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
	fmt.Println()
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
	fmt.Println()
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
	fmt.Println()
}

func postOrderTraverse(n *Node) {
	if n == nil {
		return
	}
	postOrderTraverse(n.leftChild)
	postOrderTraverse(n.rightChild)
	fmt.Print(n.value)
}

// 层序遍历
func (t *Tree) LevelOrderTraverse() {
	levelOrderTraverse(t.root)
	fmt.Println()
}

func levelOrderTraverse(n *Node) {
	if n == nil {
		return
	}

	queue := simplequeue.New()
	queue.Push(n)
	for {
		val, ok := queue.Pop()
		if !ok {
			break
		}
		e := val.(*Node)
		fmt.Print(e.value)
		if e.leftChild != nil {
			queue.Push(e.leftChild)
		}
		if e.rightChild != nil {
			queue.Push(e.rightChild)
		}
	}
}

func TestTraverse(t *testing.T) {
	fmt.Println("先序遍历")
	tree := genBinaryTree()
	tree.PreOrderTraverse()

	fmt.Println("中序遍历")
	tree.InOrderTraverse()

	fmt.Println("后序遍历")
	tree.PostOrderTraverse()

	fmt.Println("层序遍历")
	tree.LevelOrderTraverse()
	// Output:
	// 先序遍历
	// ABDGHCEIF
	// 中序遍历
	// GDHBAEICF
	// 后序遍历
	// GHDBIEFCA
	// 层序遍历
	// ABCDEFGHI
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
