package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeToSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	root := SliceToTree(s)
	PrintTree(root)
}

func TestSliceToTree(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	root := SliceToTree(s)

	s2 := TreeToSlice(root)
	assert.Equal(t, s, s2)
}
