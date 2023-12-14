// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_backtracking

package question

import (
	"fmt"
	"testing"
)

func TestPermutationI(t *testing.T) {
	/* 全排列 I */
	nums := []int{1, 2, 3}
	fmt.Printf("输入数组 nums = ")
	fmt.Println(nums)

	res := permutationsI(nums)
	fmt.Printf("所有排列 res = ")
	fmt.Println(res)

	// Output:
	// 输入数组 nums = [1 2 3]
	// 所有排列 res = [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}

func TestPermutationII(t *testing.T) {
	nums := []int{1, 2, 2}
	fmt.Printf("输入数组 nums = ")
	fmt.Println(nums)

	res := permutationsII(nums)
	fmt.Printf("所有排列 res = ")
	fmt.Println(res)

	// Output:
	// 输入数组 nums = [1 2 2]
	// 所有排列 res = [[1 2 2] [2 1 2] [2 2 1]]
}
