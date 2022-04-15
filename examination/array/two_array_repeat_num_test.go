package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 给定2个有序数组, 每个数组内部元素不重复, 计算2个数组值相等元素的交集个数
func Test(t *testing.T) {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{0, 2, 5, 6, 9}
	n := calRepeat(arr1, arr2)
	assert.Equal(t, 2, n)
}

func calRepeat(arr1 []int, arr2 []int) int {
	var i, j, n int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			n++
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return n
}
