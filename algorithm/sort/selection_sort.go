package sort

// 选择排序
// 1.最差时间复杂度和平均时间复杂度: O(n^2)
// 2.空间复杂度: O(1)
// 3.稳定性: 不稳定排序
func SelectionSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for i := 0; i < size-1; i++ {
		minPos := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}
		arr[i], arr[minPos] = arr[minPos], arr[i]
	}
}
