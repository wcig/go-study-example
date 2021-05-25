package ch6_array

import (
	"fmt"
	"testing"
)

// 数组初始化
func TestArrayInit(t *testing.T) {
	// 先声明再赋值
	var arr1 [5]int
	fmt.Println("arr1：", arr1)

	arr1[3] = 100
	fmt.Println("arr1：", arr1)

	// 声明的同时赋值
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)
}

// 数组多种初始化方式
func TestArraySpecialInit(t *testing.T) {
	var arr1 [3]int
	fmt.Println(arr1) // [0 0 0]

	var arr2 = [3]int{}
	fmt.Println(arr2) // [0 0 0]

	arr3 := [3]int{0: 0, 1: 1, 2: 2}
	fmt.Println(arr3) // [0 1 2]

	arr4 := [3]int{1: 1}
	fmt.Println(arr4) // [0 1 0]

	var arr5 = [...]int{0, 1, 2}
	fmt.Println(arr5) // [0 1 2]

	arr6 := [...]int{2: 2}
	fmt.Println(arr6) // [0 0 2]
}

// 数组比较
func TestCompareArray(t *testing.T) {
	// 数组大小
	var arr7 = [...]int{0, 1, 2}
	var size = len(arr7)
	fmt.Println("size=", size)

	// 数组比较 ==, !=
	arr8 := [3]int{0, 1, 2}
	// arr8 := [2]int{0, 1}	// 不同大小不能比较
	// arr8 := [3]string{}	// 不同类型不能比较
	fmt.Println(arr7 == arr8)
	fmt.Println(arr7 != arr8)
}

// 指针数组、数组指针
func TestArrayPointer(t *testing.T) {
	// 指针数组
	a, b, c := 1, 2, 3
	arr1 := [3]*int{&a, &b, &c}
	fmt.Println("arr1=", arr1)

	// 数组指针初始化3种方式
	// 1.指向其他数组
	arr2 := [3]int{0, 1, 2}
	arr3 := &arr2
	fmt.Println("arr3=", arr3)

	// 2.先声明再初始化
	var arr4 *[3]int
	arr4 = &[3]int{1, 2, 3}
	fmt.Println("arr4=", arr4)

	// 3.直接初始化
	arr5 := &[3]int{1, 2, 3}
	fmt.Println("arr5=", arr5)
}

// output:
// arr1= [0xc000058080 0xc000058088 0xc000058090]
// arr3= &[0 1 2]
// arr4= &[1 2 3]
// arr5= &[1 2 3]

/**
 * 数组声明注意：
 * 1.数组声明：自动赋默认值
 * 2.指针数组声明：自动赋值nil
 * 3.数组指针声明：默认nil
 */
func TestArrayInitDefaultValue(t *testing.T) {
	var arr1 [3]int
	fmt.Println("arr1=", arr1)
	fmt.Println("size1=", len(arr1))

	var arr2 [3]*int
	fmt.Println("arr2=", arr2)
	fmt.Println("size2=", len(arr2))

	var arr3 *[3]int
	fmt.Println("arr3=", arr3)
	fmt.Println("size3=", len(arr3))
}

// output:
// arr1= [0 0 0]
// size1= 3
// arr2= [<nil> <nil> <nil>]
// size2= 3
// arr3= <nil>
// size3= 3

// 函数返回值为数组、指针数组、数组指针
func TestArrayAsReturnValue(t *testing.T) {
	arr1 := t1()
	fmt.Println("arr1=", arr1)
	fmt.Println("size1=", len(arr1))

	arr2 := t2()
	fmt.Println("arr2=", arr2)
	fmt.Println("size2=", len(arr2))

	arr3 := t3()
	fmt.Println("arr3=", arr3)
	fmt.Println("size3=", len(arr3))
}

func t1() [3]int {
	return [3]int{}
}
func t2() [3]*int {
	return [3]*int{}
}
func t3() *[3]int {
	return nil
}

// 数组为值类型：作为参数传递，值不会改变
func TestArrayAsFuncParam1(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Println("before modify, arr=", arr)
	modify1(arr)
	fmt.Println("after  modify, arr=", arr)
}

func modify1(arr [3]int) {
	arr[0] = 0
}

// output：
// before modify, arr= [1 2 3]
// after  modify, arr= [1 2 3]

// 指针数组为值类型：作为参数传递，值不会改变
func TestArrayAsFuncParam2(t *testing.T) {
	a, b, c := 1, 2, 3
	arr := [3]*int{&a, &b, &c}
	fmt.Println("before modify, arr=", arr)
	for _, value := range arr {
		fmt.Println(*value)
	}

	modify2(arr)
	fmt.Println("after  modify, arr=", arr)
	for _, value := range arr {
		fmt.Println(*value)
	}
}

func modify2(arr [3]*int) {
	d := 0
	arr[0] = &d
}

// output：
// before modify, arr= [0xc00000a0b8 0xc00000a0d0 0xc00000a0d8]
// 1
// 2
// 3
// after  modify, arr= [0xc00000a0b8 0xc00000a0d0 0xc00000a0d8]
// 1
// 2
// 3

// 数组指针为引用类型：作为参数传递，值会改变
func TestArrayAsFuncParam3(t *testing.T) {
	arr := &[3]int{1, 2, 3}
	fmt.Println("before modify, arr=", arr)
	modify3(arr)
	fmt.Println("after  modify, arr=", arr)
}

func modify3(arr *[3]int) {
	arr[0] = 0 // 等价于：(*arr)[0] = 0
}

// output：
// before modify, arr= &[1 2 3]
// after  modify, arr= &[0 2 3]

// new创建数组 (此时为数组指针)
func TestNewArray(t *testing.T) {
	arr := new([3]int)       // 等价于：var arr *[3]int
	fmt.Println("arr=", arr) // &[0 0 0]
}

// output：
// arr= &[0 0 0]

// 多维数组
func TestMultidimensionalArray(t *testing.T) {
	arr1 := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println("arr1=", arr1)

	arr2 := [...][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println("arr2=", arr2)

	// error
	// arr3 := [2][...]int{
	// 	{0, 1, 2},
	// 	{3, 4, 5},
	// }
	// fmt.Println("arr3=", arr3)
}

// output:
// arr1= [[0 1 2] [3 4 5]]
// arr2= [[0 1 2] [3 4 5]]
