// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_backtracking

package question

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSubsetSumINaive(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	res := subsetSumINaive(nums, target)

	fmt.Printf("target = " + strconv.Itoa(target) + ", 输入数组 nums = ")
	fmt.Println(nums)

	fmt.Println("所有和等于 " + strconv.Itoa(target) + " 的子集 res = ")
	for i := range res {
		fmt.Println(res[i])
	}
	fmt.Println("请注意，该方法输出的结果包含重复集合")

	// Output:
	// target = 9, 输入数组 nums = [3 4 5]
	// 所有和等于 9 的子集 res =
	// [3 3 3]
	// [4 5]
	// [5 4]
	// 请注意，该方法输出的结果包含重复集合
}

func TestSubsetSumI(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	res := subsetSumI(nums, target)

	fmt.Printf("target = " + strconv.Itoa(target) + ", 输入数组 nums = ")
	fmt.Println(nums)

	fmt.Println("所有和等于 " + strconv.Itoa(target) + " 的子集 res = ")
	for i := range res {
		fmt.Println(res[i])
	}

	// Output:
	// target = 9, 输入数组 nums = [3 4 5]
	// 所有和等于 9 的子集 res =
	// [3 3 3]
	// [4 5]
}

func TestSubsetSumII(t *testing.T) {
	nums := []int{4, 4, 5}
	target := 9
	res := subsetSumII(nums, target)

	fmt.Printf("target = " + strconv.Itoa(target) + ", 输入数组 nums = ")
	fmt.Println(nums)

	fmt.Println("所有和等于 " + strconv.Itoa(target) + " 的子集 res = ")
	for i := range res {
		fmt.Println(res[i])
	}

	// Output:
	// target = 9, 输入数组 nums = [4 4 5]
	// 所有和等于 9 的子集 res =
	// [4 5]
}
