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
	//arr := []int{2, 4, 3, 1, 5, 9, 0, 7, 6, 8}
	arr := []int{1, 2, 6, 7, 9, 0, 3, 4, 5, 8}
	MergeSort(arr)
	fmt.Println(arr)

	//sortCheck(t, MergeSort)
}

func Test(t *testing.T) {
	//for i := 0; i < 100; i++ {
	//	fmt.Println(genRandArr(10, -9, 9))
	//}

	//s := make([]int, 0)
	//fmt.Println(s)
	//fmt.Println(copyArr(s))
}
