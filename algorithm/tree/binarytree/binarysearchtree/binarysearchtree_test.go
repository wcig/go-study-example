package binarysearchtree

import (
	"go-app/algorithm/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genTree() *BinarySearchTree {
	tree := New(utils.IntComparator)

	// left
	n2 := &Node{2, nil, nil}
	n1 := &Node{1, nil, n2}
	n4 := &Node{4, nil, nil}
	n3 := &Node{3, n1, n4}

	// right
	n6 := &Node{6, nil, nil}
	n7 := &Node{7, n6, nil}
	n9 := &Node{9, nil, nil}
	n8 := &Node{8, n7, n9}

	// root
	n5 := &Node{5, n3, n8}
	tree.root = n5
	tree.size = 9
	return tree
}

func TestFindMinMax(t *testing.T) {
	tree := New(utils.IntComparator)
	tree.InOrderTraverse(utils.IntPrinter)
	val, ok := tree.FindMin()
	assert.False(t, ok)
	assert.Nil(t, val)
	val, ok = tree.FindMax()
	assert.False(t, ok)
	assert.Nil(t, val)

	tree = genTree()
	tree.InOrderTraverse(utils.IntPrinter)
	val, ok = tree.FindMin()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = tree.FindMax()
	assert.True(t, ok)
	assert.Equal(t, 9, val)
}

func TestContains(t *testing.T) {
	var arr []int
	tree := New(utils.IntComparator)
	for i := -16; i <= 16; i++ {
		if tree.Contains(i) {
			arr = append(arr, i)
		}
	}
	assert.Equal(t, 0, len(arr))

	tree = genTree()
	for i := -16; i <= 16; i++ {
		if tree.Contains(i) {
			arr = append(arr, i)
		}
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, arr)
}

func TestAdd(t *testing.T) {
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
	tree.Add(6)

	tree.InOrderTraverse(utils.IntPrinter)
}

func TestRemove(t *testing.T) {
	tree := genTree()
	tree.InOrderTraverse(utils.IntPrinter)

	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
	}
	for i := 0; i < 10000; i++ {
		for k := range m {
			tree.Remove(k)
		}
		assert.Equal(t, 0, tree.size)
	}
}

func TestTraverse(t *testing.T) {
	tree := New(utils.IntComparator)
	tree.PreOrderTraverse(utils.IntPrinter)
	tree.InOrderTraverse(utils.IntPrinter)
	tree.PostOrderTraverse(utils.IntPrinter)

	tree = genTree()
	tree.PreOrderTraverse(utils.IntPrinter)
	tree.InOrderTraverse(utils.IntPrinter)
	tree.PostOrderTraverse(utils.IntPrinter)
}
