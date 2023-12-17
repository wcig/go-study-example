// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package question

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	target := 6
	noTarget := 99
	targetIndex := binarySearch(nums, target)
	fmt.Println("目标元素 6 的索引 = ", targetIndex)
	noTargetIndex := binarySearch(nums, noTarget)
	fmt.Println("不存在目标元素的索引 = ", noTargetIndex)
}
