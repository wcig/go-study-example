package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 将一个数组的所有元素向右移动若干单位，并把数组右侧溢出的元素填补
// 在数组左侧的空缺中，这种经操作称为数组的循环平移。
//
// 给你一个不小于 3 个元素的数组 a，已知 a 是从一个有序且不包含
// 重复元素的数组平移 k(k 大于等于 0 且小于数组长度)个单位而来；
// 请写一个函数，输入 int 类型数组 a，返回 k 的值。
//
// 例如，对于数组 a = {5, 1, 2, 3, 4}，它由有序数组 {1, 2, 3, 4, 5} 循环平移 1 个单位而来，因此 k = 1。
func TestStepMove(t *testing.T) {
	testCases := []struct {
		in     []int
		ou     []int
		except int
	}{
		{
			in:     []int{1, 2, 3, 4, 5},
			ou:     []int{1, 2, 3, 4, 5},
			except: 0,
		},
		{
			in:     []int{1, 2, 3, 4, 5},
			ou:     []int{5, 1, 2, 3, 4},
			except: 1,
		},
		{
			in:     []int{1, 2, 3, 4, 5},
			ou:     []int{4, 5, 1, 2, 3},
			except: 2,
		},
		{
			in:     []int{1, 2, 3, 4, 5},
			ou:     []int{3, 4, 5, 1, 2},
			except: 3,
		},
		{
			in:     []int{1, 2, 3, 4, 5},
			ou:     []int{2, 3, 4, 5, 1},
			except: 4,
		},
	}

	for _, v := range testCases {
		step := calMoveStep(v.in, v.ou)
		assert.Equal(t, v.except, step)
	}
}

func calMoveStep(in []int, ou []int) int {
	if len(in) != len(ou) {
		return -1
	}

	size := len(in)
	result := -1
	for step := 0; step < size; step++ {
		match := true
		for i := 0; i < size; i++ {
			index := (i + step) % size
			if in[i] != ou[index] {
				match = false
			}
		}
		if match {
			result = step
			break
		}
	}
	return result
}
