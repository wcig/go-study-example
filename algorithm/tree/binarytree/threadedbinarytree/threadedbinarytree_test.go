package threadedbinarytree

import (
	"testing"
)

//             A
//         B           C
//     D           E       F
// G       H           I
func genBinaryTree() *ThreadedBinaryTree {
	tree := &ThreadedBinaryTree{}

	// left child tree
	g := &Node{"G", nil, nil, 0, 0}
	h := &Node{"H", nil, nil, 0, 0}
	d := &Node{"D", g, h, 0, 0}
	b := &Node{"B", d, nil, 0, 0}

	// right child tree
	i := &Node{"I", nil, nil, 0, 0}
	e := &Node{"E", nil, i, 0, 0}
	f := &Node{"F", nil, nil, 0, 0}
	c := &Node{"C", e, f, 0, 0}

	a := &Node{"A", b, c, 0, 0}
	tree.root = a
	return tree
}

func TestCreatePreThread(t *testing.T) {
	tree := genBinaryTree()
	// tree.PreOrderTraverse() // ABDGHCEIF

	tree.CreatePreThread()
}

func TestCreateInThread(t *testing.T) {
	tree := genBinaryTree()
	tree.InOrderTraverse() // GDHBAEICF

	tree.CreateInThread()
	tree.InOrder()    // GDHBAEICF
	tree.RevInOrder() // FCIEABHDG
}

func TestCreatePostThread(t *testing.T) {
	tree := genBinaryTree()
	tree.PostOrderTraverse() // GHDBIEFCA

	tree.CreatePostThread()
}
