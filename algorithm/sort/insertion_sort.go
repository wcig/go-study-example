package sort

// 插入排序
// 1.最差时间复杂度和平均时间复杂度: O(n^2) (当输入数组完全有序时间复杂度O(n))
// 2.空间复杂度: O(1)
// 3.稳定性: 稳定排序
func InsertionSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for i := 1; i < size; i++ {
		base := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > base {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = base
	}
}
