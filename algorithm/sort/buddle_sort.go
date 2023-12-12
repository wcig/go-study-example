package sort

// 冒泡排序:
// 1.最差时间复杂度和平均时间复杂度: O(n^2)
// 2.空间复杂度: O(1)
// 3.稳定性: 稳定排序
func BubbleSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 冒泡排序 (标志优化):
// 1.最差时间复杂度和平均时间复杂度: O(n^2) (优化后当输入数组完全有序时间复杂度O(n))
// 2.空间复杂度: O(1)
// 3.稳定性: 稳定排序
func BubbleSortWithFlag(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for i := size - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
