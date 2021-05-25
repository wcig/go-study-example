package ch7_slice

import (
	"errors"
	"fmt"
	"testing"
)

func printIntSlice(s []int) {
	fmt.Println(s, len(s), cap(s))
}

// 切片声明
func TestSliceDeclare(t *testing.T) {
	var s1 []int
	printIntSlice(s1)

	s2 := make([]int, 3, 10)
	// 等价: var s2 []int = make([]int, 3, 10)
	// 等价: var s2 := make([]int, 3, 10)
	printIntSlice(s2)

	s3 := make([]int, 3)
	printIntSlice(s3)
}

// output:
// [] 0 0
// [0 0 0] 3 10
// [0 0 0] 3 3

// 切片初始化
func TestSliceInit(t *testing.T) {
	// 直接初始化
	var s1 = []int{1, 2, 3}
	printIntSlice(s1)

	// 声明后再添加值
	var s2 []int
	s2 = append(s2, 1)
	printIntSlice(s2)

	// 声明后再赋值
	s3 := make([]int, 3)
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	printIntSlice(s3)

	// 通过数组初始化
	arr := [5]int{1, 2, 3, 4, 5}
	s4 := arr[1:4]
	printIntSlice(s4)

	// 通过切片初始化
	s := []int{1, 2, 3, 4, 5}
	s5 := s[1:4]
	printIntSlice(s5)
}

// 从数组初始化
func TestSliceInitFormArray(t *testing.T) {
	arr := [5]int{0, 1, 2, 3, 4}

	s1 := arr[:]
	printIntSlice(s1)

	s2 := arr[1:]
	printIntSlice(s2)

	s3 := arr[:4]
	printIntSlice(s3)

	s4 := arr[1:4]
	printIntSlice(s4)
}

// output:
// [0 1 2 3 4] 5 5
// [1 2 3 4] 4 4
// [0 1 2 3] 4 5
// [1 2 3] 3 4

// slice修改会影响源数组或源切片
func TestOutputSliceModify(t *testing.T) {
	arr := [5]int{0, 1, 2, 3, 4}
	s1 := arr[1:3]
	s1[0] = 100
	fmt.Println(arr)

	s := []int{0, 1, 2, 3, 4}
	s2 := s[1:3]
	s2[0] = 100
	fmt.Println(s)
}

// output:
// [0 100 2 3 4]
// [0 100 2 3 4]

// 从一个切片初始化另衣蛾切片
func TestSliceInitFormSlice(t *testing.T) {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[1:5]
	printIntSlice(s1)

	s2 := s1[1:3]
	// s2 := s1[1:10] // error：索引不能超过原切片cap容量
	printIntSlice(s2)

	s2[0] = 100
	fmt.Println("--- after modify ---")
	fmt.Println(arr)
	printIntSlice(s1)
	printIntSlice(s2)

	// s2[3] = 10 // error：索引不能超出切片大小
}

// output:
// [1 2 3 4] 4 9
// [2 3] 2 8
// --- after modify ---
// [0 1 100 3 4 5 6 7 8 9]
// [1 100 3 4] 4 9
// [100 3] 2 8

// 添加元素 (append添加元素后切片地址会变，所以需要将添加元素后的切片赋值给源切片)
func TestAppend(t *testing.T) {
	s1 := make([]int, 3, 5)
	printIntSlice(s1)
	fmt.Printf("s1=%p\n", s1)

	s1 = append(s1, 1)
	printIntSlice(s1)
	fmt.Printf("s1=%p\n", s1)

	s1 = append(s1, 2, 3)
	printIntSlice(s1)
	fmt.Printf("s1=%p\n", s1)

	var s2 []int
	printIntSlice(s2)
	s2 = append(s2, 1)
	printIntSlice(s2)
}

// output:
// [0 0 0] 3 5
// s1=0xc000082030
// [0 0 0 1] 4 5
// s1=0xc000082030
// [0 0 0 1 2 3] 6 10
// s1=0xc0000c8000
// [] 0 0
// [1] 1 1

// 删除元素
func TestRemove(t *testing.T) {
	fmt.Println(removeIndexInOrder([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(removeIndexInOrder([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(removeIndexInOrder([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(removeIndexInOrder([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(removeIndexInOrder([]int{1, 2, 3, 4, 5}, -1))

	fmt.Println(removeIndexNoOrder([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(removeIndexNoOrder([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(removeIndexNoOrder([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(removeIndexNoOrder([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(removeIndexNoOrder([]int{1, 2, 3, 4, 5}, -1))

	fmt.Println(removeIndexInOrder(nil, 0))
	fmt.Println(removeIndexInOrder(nil, -1))

	fmt.Println(removeIndexNoOrder(nil, 0))
	fmt.Println(removeIndexNoOrder(nil, -1))
}

// 保持源切片顺序删除元素
func removeIndexInOrder(s []int, index int) (result []int, err error) {
	if index < 0 || index >= len(s) {
		return nil, errors.New("index out of bounds")
	}
	return append(s[:index], s[index+1:]...), nil
}

// 不保持源切片顺序删除元素
func removeIndexNoOrder(s []int, index int) (result []int, err error) {
	if index < 0 || index >= len(s) {
		return nil, errors.New("index out of bounds")
	}

	s[index] = s[len(s)-1] // 或者: s[index], s[len(s)-1] = s[len(s)-1], s[index]
	return s[:len(s)-1], nil
}

// copy
func TestCopy(t *testing.T) {
	// copy后将不会相互影响
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copy(s2, s1)

	s2[0] = 100
	printIntSlice(s1)
	printIntSlice(s2)

	// copy目标切片比源切片小
	s3 := []int{1, 2, 3, 4, 5}
	s4 := make([]int, 3)
	copy(s4, s3)
	printIntSlice(s4)

	// copy目标切片比源切片大
	s5 := []int{6, 7, 8}
	s6 := make([]int, 5)
	copy(s6, s5)
	printIntSlice(s6)
}

// output:
// [1 2 3 4 5] 5 5
// [100 2 3 4 5] 5 5
// [1 2 3] 3 3
// [6 7 8 0 0] 5 5

// 空切片
func TestEmptySlice(t *testing.T) {
	var s1 []int
	printIntSlice(s1)

	s2 := make([]int, 0)
	printIntSlice(s2)

	s3 := getNullSlice()
	printIntSlice(s3)

	s4 := getPointerNullSlice()
	fmt.Println(s4)
	// printIntSlice(*s4) // 错误：invalid memory address or nil pointer dereference
}

// output:
// [] 0 0
// [] 0 0
// [] 0 0
// <nil>

func getNullSlice() []int {
	return nil
}

func getPointerNullSlice() *[]int {
	return nil
}

// 切片清空
func TestClearSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	printIntSlice(s1)
	fmt.Printf("s1=%p\n", s1)

	s1 = s1[:0] // 等价于：s1 = s1[0:0]
	printIntSlice(s1)
	fmt.Printf("s1=%p\n", s1)
}

// output:
// [1 2 3 4 5] 5 5
// s1=0xc000082030
// [] 0 5
// s1=0xc000082030

// range遍历
func TestRange(t *testing.T) {
	// 遍历int数组
	s1 := []int{1, 2, 3}
	for index, value := range s1 {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// 简化index
	for _, value := range s1 {
		fmt.Printf("value: %d\n", value)
	}

	// 简化value
	for index := range s1 {
		fmt.Printf("index: %d, value: %d\n", index, s1[index])
	}
}

// range遍历基础类型切片和指针类型切片区别
func TestRangePointerSlice(t *testing.T) {
	n1, n2, n3 := 1, 2, 3

	// 遍历基础类型切片：只能通过index修改切片
	s1 := []int{n1, n2, n3}
	for index, value := range s1 {
		fmt.Printf("index: %d, value: %d, value-addr:%X, elem-addr:%X\n", index, value, &value, &s1[index])
	}

	for _, value := range s1 {
		value += 1
	}
	printIntSlice(s1)

	for index := range s1 {
		s1[index] += 1
	}
	printIntSlice(s1)

	// 遍历指针切片：即可通过index，也可通过val修改切片
	s2 := []*int{&n1, &n2, &n3}
	for index, value := range s2 {
		fmt.Printf("index: %d, value: %d, value-addr:%X, elem-addr:%X\n", index, *value, &value, &s1[index])
	}

	for _, value := range s2 {
		*value += 1
	}

	for index := range s2 {
		fmt.Printf("%d, ", *s2[index])
	}
	fmt.Println()

	for index := range s2 {
		*s2[index] += 1
	}

	for index := range s2 {
		fmt.Printf("%d, ", *s2[index])
	}
	fmt.Println()
}

// output:
// index: 0, value: 1, value-addr:C00008A240, elem-addr:C0000B20E0
// index: 1, value: 2, value-addr:C00008A240, elem-addr:C0000B20E8
// index: 2, value: 3, value-addr:C00008A240, elem-addr:C0000B20F0
// [1 2 3] 3 3
// [2 3 4] 3 3
// index: 0, value: 1, value-addr:C000076028, elem-addr:C0000B20E0
// index: 1, value: 2, value-addr:C000076028, elem-addr:C0000B20E8
// index: 2, value: 3, value-addr:C000076028, elem-addr:C0000B20F0
// 2, 3, 4,
// 3, 4, 5,

// 二维切片
func TestTwoDSlice(t *testing.T) {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]
}
