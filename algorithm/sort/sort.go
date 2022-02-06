package sort

import "fmt"

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

	fmt.Println(arr)
	mergeSortMerge(arr, 0, 5, 9)
}

func mergeSortSub(arr []int) {
	//
}

func mergeSortMerge(arr []int, leftPtr, rightPtr, rightBound int) {
	//tmpArr := make([]int, rightBound-leftPtr+1)
	//
}
