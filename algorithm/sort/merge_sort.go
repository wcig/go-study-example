package sort

// 归并排序
// 1.时间复杂度: O(n*logn)
// 2.空间复杂度: O(n)
// 3.稳定性: 稳定排序
func MergeSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = arr[j]
		j++
		k++
	}
	for m := 0; m < len(tmp); m++ {
		arr[m+left] = tmp[m]
	}
}
