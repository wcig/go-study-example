package heap

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 给定一个长度为的无序数组nums, 请返回数组中前大的元素 (不要求顺序, 以下解答未考虑边界情况)
func TestTopK(t *testing.T) {
	nums := []int{1, 7, 6, 3, 2}
	k := 3
	expect := []int{3, 6, 7}

	out1 := TopKWithHeap(nums, k)
	sort.Ints(out1)
	assert.Equal(t, expect, out1)

	out2 := TopKWithSort(nums, k)
	sort.Ints(out2)
	assert.Equal(t, expect, out2)

	out3 := TopKWithTraverse(nums, k)
	sort.Ints(out3)
	assert.Equal(t, expect, out3)
}

func TopKWithHeap(nums []int, k int) []int {
	h := NewMinHeap()
	for i := 0; i < k; i++ {
		h.Push(nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Top() {
			h.Pop()
			h.Push(nums[i])
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = h.Pop()
	}
	return result
}

func TopKWithSort(nums []int, k int) []int {
	sort.Ints(nums)
	return nums[len(nums)-k:]
}

func TopKWithTraverse(nums []int, k int) []int {
	result := make([]int, 0, k)
	last := math.MinInt
	for i := 0; i < k; i++ {
		max := math.MinInt
		for j := 0; j < len(nums); j++ {
			if nums[j] > max && ((nums[j] < last) || (i == 0)) {
				max = nums[j]
			}
		}
		last = max
		result = append(result, max)
	}
	return result
}
