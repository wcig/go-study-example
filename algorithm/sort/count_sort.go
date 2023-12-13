package sort

// 计数排序
// 1.平均时间复杂度：O(n+m) (最坏时间复杂度：O(n+m)，最好时间复杂度：O(n+m)), m为取值范围.
// 2.空间复杂度：O(n+m)
// 3.稳定性: 稳定排序
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
