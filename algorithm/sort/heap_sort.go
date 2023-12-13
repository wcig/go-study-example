package sort

// 堆排序
// 1.平均时间复杂度：O(n*logn)
// 2.空间复杂度：O(1)
// 3.稳定性: 不稳定排序
func HeapSort(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, 0, i)
	}
}

func siftDown(arr []int, i, n int) {
	for {
		max := i
		l, r := 2*i+1, 2*i+2
		if l < n && arr[l] > arr[max] {
			max = l
		}
		if r < n && arr[r] > arr[max] {
			max = r
		}
		if i == max {
			break
		}
		arr[i], arr[max] = arr[max], arr[i]
		i = max
	}
}
