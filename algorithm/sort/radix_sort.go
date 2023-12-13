package sort

import "math"

// 基数排序
// 1.平均时间复杂度：O(n*k) (最坏时间复杂度：O(n*k)，最好时间复杂度：O(n*k)), k为最大数的位数.
// 2.空间复杂度：O(n+k)
// 3.稳定性: 稳定排序
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
