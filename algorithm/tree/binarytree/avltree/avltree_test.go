package avltree

import (
	"go-app/algorithm/utils"
	"testing"
)

func genTree() *AVLTree {
	tree := New(utils.IntComparator)

	// left
	n2 := &Node{2, nil, nil, 0}
	n1 := &Node{1, nil, n2, 0}
	n4 := &Node{4, nil, nil, 0}
	n3 := &Node{3, n1, n4, 0}

	// right
	n6 := &Node{6, nil, nil, 0}
	n7 := &Node{7, n6, nil, 0}
	n9 := &Node{9, nil, nil, 0}
	n8 := &Node{8, n7, n9, 0}

	// root
	n5 := &Node{5, n3, n8, 0}
	tree.root = n5
	tree.size = 9
	return tree
}

func Test(t *testing.T) {
	tree := New(utils.IntComparator)

	// root
	tree.Add(5)

	// left
	tree.Add(3)
	tree.Add(4)
	tree.Add(1)
	tree.Add(2)

	// right
	tree.Add(8)
	tree.Add(9)
	tree.Add(7)
	tree.InOrderTraverse(utils.IntPrinter)

	tree.Add(6)

	tree.InOrderTraverse(utils.IntPrinter)
}
