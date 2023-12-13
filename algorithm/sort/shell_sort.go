package sort

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
