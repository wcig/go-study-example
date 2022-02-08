package sort

import (
	"math"
)

// 选择排序
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

// 冒泡排序
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

// 插入排序
func InsertionSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for i := 1; i < size; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// 希尔排序 (间隔为2h)
func ShellSortSimple(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	for gap := size / 2; gap > 0; gap /= 2 {
		for i := gap; i < size; i++ {
			for j := i; j > gap-1 && arr[j] < arr[j-gap]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
}

// 希尔排序 (间隔为3h+1)
func ShellSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	h := 1
	for h <= size/3 {
		h = h*3 + 1
	}

	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < size; i++ {
			for j := i; j > gap-1 && arr[j] < arr[j-gap]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
}

// 归并排序
func MergeSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	mergeSortSub(arr, 0, len(arr)-1)
}

func mergeSortSub(arr []int, left, right int) {
	if left == right {
		return
	}

	mid := left + (right-left)/2
	mergeSortSub(arr, left, mid)
	mergeSortSub(arr, mid+1, right)
	mergeSortMerge(arr, left, mid+1, right)
}

func mergeSortMerge(arr []int, leftPtr, rightPtr, rightBound int) {
	tmpSize := rightBound - leftPtr + 1
	tmpArr := make([]int, tmpSize)

	i := leftPtr
	j := rightPtr
	k := 0
	mid := rightPtr - 1

	for i <= mid && j <= rightBound {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			k++
			i++
		} else {
			tmpArr[k] = arr[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmpArr[k] = arr[i]
		k++
		i++
	}
	for j <= rightBound {
		tmpArr[k] = arr[j]
		k++
		j++
	}

	for m := 0; m < tmpSize; m++ {
		arr[leftPtr+m] = tmpArr[m]
	}
}

// 快速排序
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
	pivot := right
	i := left
	j := right - 1

	for i <= j {
		for i <= j && arr[i] <= arr[pivot] {
			i++
		}
		for i <= j && arr[j] > arr[pivot] {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}

// 计数排序
func CountSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	min := arr[0]
	max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	countSortWithMinMax(arr, min, max)
}

func countSortWithMinMax(arr []int, min, max int) {
	countArr := make([]int, max-min+1)
	for i := range arr {
		countArr[arr[i]-min]++
	}
	for i := 1; i < len(countArr); i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}

	tmpArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		pos := countArr[arr[i]-min] - 1
		tmpArr[pos] = arr[i]
		countArr[arr[i]-min]--
	}
	for i := range tmpArr {
		arr[i] = tmpArr[i]
	}
}

// 基数排序
func RadixSort(arr []int) {
	size := len(arr)
	if size == 0 {
		return
	}

	// 计算位数
	min := arr[0]
	max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	digits := 1
	if max > 0 {
		val := max
		if min < 0 {
			val -= min
		}
		digits = int(math.Log10(float64(val))) + 1
	}

	// 兼容数组元素负数
	if min < 0 {
		for i := range arr {
			arr[i] -= min
		}
	}

	// 排序
	tmpArr := make([]int, size)
	countArr := make([]int, 10)
	for i := 0; i < digits; i++ {
		division := int(math.Pow10(i))
		for j := 0; j < size; j++ {
			num := (arr[j] / division) % 10
			countArr[num]++
		}
		for j := 1; j < len(countArr); j++ {
			countArr[j] = countArr[j] + countArr[j-1]
		}
		for j := size - 1; j >= 0; j-- {
			num := (arr[j] / division) % 10
			pos := countArr[num] - 1
			tmpArr[pos] = arr[j]
			countArr[num]--
		}

		copy(arr, tmpArr)
		for j := range countArr {
			countArr[j] = 0
		}
	}

	// 兼容数组元素负数
	if min < 0 {
		for i := range arr {
			arr[i] += min
		}
	}
}
