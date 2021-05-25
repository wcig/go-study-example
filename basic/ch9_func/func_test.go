package ch9_func

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 函数定义和使用
func TestFuncDefine(t *testing.T) {
	m := max(10, 20)
	fmt.Println("m=", m)
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// 多返回值参数
func TestFuncMultipleReturnValues(t *testing.T) {
	x1, y1 := vals1()
	fmt.Println(x1, y1) // 1 2

	// 返回值参数在函数中不赋值时将返回该类型默认零值
	x2, y2 := vals2()
	fmt.Println(x2, y2) // 3 0
}

func vals1() (int, int) {
	return 1, 2
}

func vals2() (x, y int) {
	x = 3
	return x, y
}

// 不定长变参
func TestFuncVariadic1(t *testing.T) {
	fmt.Println(sum(1, 2, 3))
}

func sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestFuncVariadic2(t *testing.T) {
	s := []int{1, 2, 3}
	prefix := "hi "
	printSlices(prefix, s...)
	fmt.Println(s)
}

func printSlices(prefix string, s ...int) {
	s[0] = 10
	s = append(s, 4)
	fmt.Println(prefix, s)
}

// output:
// hi  [10 2 3 4]
// [10 2 3]

// 函数传参: 值类型与引用类型
func TestFuncParamType(t *testing.T) {
	u := &user{"tom", 11}
	modify1(u)
	fmt.Println("user:", u)
	modify2(*u)
	fmt.Println("user:", u)
}

// output:
// u: &{jerry1 11}
// user: &{jerry1 11}
// u: {jerry2 11}
// user: &{jerry1 11}

type user struct {
	name string
	age  int
}

func modify1(u *user) {
	u.name = "jerry1"
	fmt.Println("u:", u)
}

func modify2(u user) {
	u.name = "jerry2"
	fmt.Println("u:", u)
}

// 匿名函数
func TestFuncAnonymous(t *testing.T) {
	printSlice := func(s []int) {
		fmt.Println(s)
	}

	s := []int{1, 2, 3}
	printSlice(s)
}

// 闭包
func TestFuncClosure(t *testing.T) {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(y int) int {
	fmt.Printf("%p\n", &x)

	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

// output:
// 0xc0000901d8
// 0xc0000901d8
// 11
// 0xc0000901d8
// 12

// defer
// defer执行顺序与定义顺序相反
func TestFuncDefer1(t *testing.T) {
	s := []int{1, 2, 3}
	for _, n := range s {
		defer fmt.Println(n)
	}
}

// output:
// 3
// 2
// 1

// 函数发生panic defer也会执行
func TestFuncDefer2(t *testing.T) {
	defer fmt.Println("defer...")
	panic("panic...")
}

// output:
// defer...
// --- FAIL: TestFuncDefer2 (0.00s)
// panic: panic... [recovered]
//	panic: panic...

// defer常用于文件关闭资源清理等
func TestFuncDefer3(t *testing.T) {
	file, err := os.Open("func_test.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// 后续文件操作
	fmt.Println("file operator...")
}

// defer函数带参数时, 此时为该参数的拷贝
func TestFuncDefer4(t *testing.T) {
	n := 0
	defer fmt.Println(n)
	n = 10

	u := &user{"tom", 11}
	defer fmt.Println(u)
	u.name = "jerry"
}

// output:
// &{jerry 11}
// 0

// defer函数与defer匿名函数区别
func TestFuncDefer5(t *testing.T) {
	n1 := 10
	defer fmt.Println("n1:", n1)
	n1 = 20

	n2 := 10
	defer func() {
		fmt.Println("n2:", n2)
	}()
	n2 = 20
}

// output:
// n2: 20
// n1: 10

// 使用defer统计函数耗时
func TestFuncDefer6(t *testing.T) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start)) // 1.001410423s
	}()
	time.Sleep(time.Second)
}

// 错误处理 panic/recover
func TestFuncPanicRecover1(t *testing.T) {
	fmt.Println("start...")
	panic("panic...")
	fmt.Println("end...")
}

// output:
// start...
// --- FAIL: TestFuncPanicRecover1 (0.00s)
// panic: panic... [recovered]
//	panic: panic...

func TestFuncPanicRecover2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover...")
		}
	}()

	fmt.Println("start...")
	gPanic()
	fmt.Println("end...")
}

func gPanic() {
	fmt.Println("before panic...")
	panic("panic...")
	fmt.Println("after panic...")
}

// output:
// start...
// before panic...
// recover...
