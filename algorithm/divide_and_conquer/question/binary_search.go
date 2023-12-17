// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package question

// 二分查找: 存在则返回元素索引, 不存在返回-1 (要求输入数组是有序的)
func binarySearch(nums []int, target int) int {
	return binarySearchDfs(nums, target, 0, len(nums)-1)
}

func binarySearchDfs(nums []int, target, left, right int) int {
	// 区间为空表示没有搜索到元素即终止
	if left > right {
		return -1
	}
	// 计算中间节点
	mid := left + (right-left)/2
	if nums[mid] == target {
		// 找到元素返回索引
		return mid
	} else if nums[mid] > target {
		// 中间节点大于目标值, 递归左半部分f(left, mid-1)
		return binarySearchDfs(nums, target, left, mid-1)
	} else {
		// 中间节点小于目标值, 递归右半部分f(mid+1, right)
		return binarySearchDfs(nums, target, mid+1, right)
	}
}
