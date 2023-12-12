package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func genRandNum(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func genRandArr(len, min, max int) []int {
	if len < 0 {
		panic("invalid array len")
	}
	if len == 0 {
		return []int{}
	}

	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = genRandNum(min, max)
	}
	return arr
}

func copyArr(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func compareArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func sortCheck(t *testing.T, sortFunc func(arr []int)) {
	testNum := 10000
	testArrLen := 1000
	for i := 0; i < testNum; i++ {
		arr := genRandArr(testArrLen, -testArrLen, testArrLen)
		arr1 := copyArr(arr)
		arr2 := copyArr(arr)

		sortFunc(arr1)
		sort.Ints(arr2)
		if !compareArr(arr1, arr2) {
			fmt.Println(arr)
			fmt.Println(arr1)
			fmt.Println(arr2)
			t.Fatal("array sort check failed")
		}
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	SelectionSort(arr)
	fmt.Println(arr)

	sortCheck(t, SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	BubbleSort(arr)
	fmt.Println(arr)

	sortCheck(t, BubbleSort)
}

func TestBubbleSortWithFlag(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	BubbleSortWithFlag(arr)
	fmt.Println(arr)

	sortCheck(t, BubbleSortWithFlag)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	InsertionSort(arr)
	fmt.Println(arr)

	sortCheck(t, InsertionSort)
}

func TestShellSortSimple(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	ShellSortSimple(arr)
	fmt.Println(arr)

	sortCheck(t, ShellSortSimple)
}

func TestShellSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	ShellSort(arr)
	fmt.Println(arr)

	sortCheck(t, ShellSort)
}

func TestMergeSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	MergeSort(arr)
	fmt.Println(arr)

	sortCheck(t, MergeSort)
}

func TestQuickSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 8, 9, 0, 7, 6, 5}
	QuickSort(arr)
	fmt.Println(arr)

	sortCheck(t, QuickSort)
}

func TestQuickMedianSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 8, 9, 0, 7, 6, 5}
	quickMedianSort(arr)
	fmt.Println(arr)

	sortCheck(t, quickMedianSort)
}

func TestQuickTailSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 8, 9, 0, 7, 6, 5}
	quickTailSort(arr)
	fmt.Println(arr)

	sortCheck(t, quickTailSort)
}

func TestCountSort(t *testing.T) {
	arr := []int{2, 4, 3, 1, 8, 9, 0, 7, 6, 5, -1}
	CountSort(arr)
	fmt.Println(arr)

	sortCheck(t, CountSort)
}

func TestRadixSort(t *testing.T) {
	// arr := []int{2, 4, 3, 1, 8, 9, 0, 7, 6, 5, -1}
	arr := []int{421, 240, 115, 532, 305, 430, 124, 0, 12, -123}
	RadixSort(arr)
	fmt.Println(arr)

	sortCheck(t, RadixSort)
}

func Test(t *testing.T) {
	//for i := 0; i < 100; i++ {
	//	fmt.Println(genRandArr(10, -9, 9))
	//}

	//s := make([]int, 0)
	//fmt.Println(s)
	//fmt.Println(copyArr(s))
}
