package sort

// 快速排序
// 1.时间复杂度: O(n*logn)
// 2.空间复杂度: O(n)
// 3.稳定性: 不稳定排序
func QuickSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	quickSortSub(arr, 0, len(arr)-1)
}

func quickSortSub(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := partition(arr, left, right)
	quickSortSub(arr, left, mid-1)
	quickSortSub(arr, mid+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := left
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= arr[pivot] {
			j--
		}
		for i < j && arr[i] <= arr[pivot] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}

// 快速排序 (基准数优化):
// 1.时间复杂度: O(n*logn)
// 2.空间复杂度: O(n)
// 3.稳定性: 不稳定排序
func QuickMedianSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	quickMedianSortSub(arr, 0, len(arr)-1)
}

func quickMedianSortSub(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := partitionMedian(arr, left, right)
	quickSortSub(arr, left, mid-1)
	quickSortSub(arr, mid+1, right)
}

func partitionMedian(arr []int, left, right int) int {
	med := medianThree(left, (left+right)/2, right)
	arr[left], arr[med] = arr[med], arr[left]
	pivot := left
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= arr[pivot] {
			j--
		}
		for i < j && arr[i] <= arr[pivot] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}

func medianThree(left, mid, right int) int {
	if (left < mid) != (left < right) {
		return left
	} else if (mid < left) != (mid < right) {
		return mid
	} else {
		return right
	}
}

// 快速排序 (尾递归优化):
// 1.时间复杂度: O(n*logn)
// 2.空间复杂度: O(logn)
// 3.稳定性: 不稳定排序
func QuickTailSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	quickTailSortSub(arr, 0, len(arr)-1)
}

func quickTailSortSub(arr []int, left, right int) {
	for left < right {
		pivot := partition(arr, left, right)
		if pivot-left < right-pivot {
			quickTailSortSub(arr, left, pivot-1)
			left = pivot + 1
		} else {
			quickTailSortSub(arr, pivot+1, right)
			right = pivot - 1
		}
	}
}
