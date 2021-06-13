package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// fmt：类似于 C 的 printf 和 scanf 的函数实现格式化的 I/O（格式化样式参考fmt包的doc.go文件）

// 函数
// 1.错误
// func Errorf(format string, a ...interface{}) error // 返回格式化字符串错误

// 2.写入到w
// func Fprint(w io.Writer, a ...interface{}) (n int, err error) // 默认格式写入数据到w，当操作数都不是字符串则之间添加空格，返回写入的字节数和错误
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) // 以指定格式化方式写入数据到w
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error) // 以默认格式写入w，操作数之间总是添加空格并附加换行符

// 3.扫描读取
//  func Fscan(r io.Reader, a ...interface{}) (n int, err error) // 扫描从r读取的内容，以空格分割存入参数中，返回成功扫描的个数，如果小于传入参数数量将返回错误（换行符算空格）
// func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) // 以指定格式扫描
// func Fscanln(r io.Reader, a ...interface{}) (n int, err error) // 类似于Fscan，但在换行符停止扫描

// 4.标准输出打印
// func Print(a ...interface{}) (n int, err error) // 默认格式写入到标准输出，如果所有操作数都是非字符串则在之间添加空格，返回写入的字节数和错误
// func Printf(format string, a ...interface{}) (n int, err error) // 以指定格式输出
// func Println(a ...interface{}) (n int, err error) // 以默认格式输出，操作数之间添加空格并在最后添加换行符

// 5.扫描读取
// func Scan(a ...interface{}) (n int, err error) // 读取标准输入，通过空格分割存储值到参数中，返回成功扫描的个数，如果传入参数小于参数数量将返回错误（换行符算空格）
// func Scanf(format string, a ...interface{}) (n int, err error) // 以指定格式读取
// func Scanln(a ...interface{}) (n int, err error) // 默认格式读取，读取到换行符停止

// 6.组装
// func Sprint(a ...interface{}) string // 默认格式组装成字符串并返回，操作数都不是字符串在在之间添加空格
// func Sprintf(format string, a ...interface{}) string // 以指定格式组装
// func Sprintln(a ...interface{}) string // 默认格式组装，操作数之间添加空格并在末尾添加换行符

// 7.扫描
// func Sscan(str string, a ...interface{}) (n int, err error) // 从传入字符串读取，通过空格分割存储值到参数中，返回成功扫描的个数，如果传入参数小于参数数量将返回错误（换行符算空格）
// func Sscanf(str string, format string, a ...interface{}) (n int, err error) // 以指定格式读取（空格分割）
// func Sscanln(str string, a ...interface{}) (n int, err error) // 默认格式读取，读取到换行符停止

// %w或%v包装错误，但只能包装一个
func TestErrorf(t *testing.T) {
	err1 := fmt.Errorf("file %s not found", "tmp.txt")
	fmt.Println(err1)

	err2 := fmt.Errorf("w wrap %w", err1)
	fmt.Println(err2)

	err3 := fmt.Errorf("v wrap %v", err1)
	fmt.Println(err3)
	// output:
	// file tmp.txt not found
	// w wrap file tmp.txt not found
	// v wrap file tmp.txt not found
}

func TestFprint(t *testing.T) {
	const name, age = "tom", 22
	n, err := fmt.Fprint(os.Stdout, "name:", name, ", age:", age, "\n")
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
	fmt.Println(n)

	nn, err := fmt.Fprint(os.Stdout, n, age)
	fmt.Println()
	fmt.Println(nn, err)
	// output:
	// name:tom, age:22
	// 17
	// 17 22
	// 5 <nil>
}

func TestFprintf(t *testing.T) {
	const name, age = "tom", 22
	n, err := fmt.Fprintf(os.Stdout, "name:%s, age:%d\n", name, age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Println(n)
	// output:
	// name:tom, age:22
	// 17
}

func TestFprintln(t *testing.T) {
	const name, age = "tom", 22
	n, err := fmt.Fprintln(os.Stdout, "name:", name, ", age:", age)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(n)
	// output:
	// name: tom , age: 22
	// 20
}

func TestFscan(t *testing.T) {
	r := strings.NewReader("name tom")
	var s1, s2 string
	n, err := fmt.Fscan(r, &s1, &s2)
	fmt.Println(n, err)
	fmt.Println(s1)
	fmt.Println(s2)
	// output:
	// 2 <nil>
	// name
	// tom
}

func TestFscanf(t *testing.T) {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(i, b, s)
	// output:
	// 3
	// 5 true gophers
}

func TestFscanln(t *testing.T) {
	s := `dmr 1771 1.61803398875
    ken 271828 3.14159`
	r := strings.NewReader(s)
	var (
		a string
		b int
		c float64
	)
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(n, a, b, c)
	}
	// output:
	// 3 dmr 1771 1.61803398875
	// 3 ken 271828 3.14159
}

func TestPrint(t *testing.T) {
	const name, age = "tom", 20
	n, err := fmt.Print("name:", name, ", age:", age, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	// output:
	// name:tom, age:20
	// 17
}

func TestPrintf(t *testing.T) {
	const name, age = "tom", 22
	n, err := fmt.Printf("name:%s, age:%d\n", name, age)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	// output:
	// name:tom, age:22
	// 17
}

func TestPrintln(t *testing.T) {
	const name, age = "tom", 22
	n, err := fmt.Println("name:", name, ", age:", age)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	// output:
	// name: tom , age: 22
	// 20
}

func TestSprint(t *testing.T) {
	const name, age = "tom", 20
	val := fmt.Sprintf("name:%s, age:%d", name, age)
	fmt.Println(val) // name:tom, age:20
}

func TestSprintf(t *testing.T) {
	const name, age = "tom", 20
	val := fmt.Sprint("name:", name, ", age:", age)
	fmt.Println(val) // name:tom, age:20
}

func TestSprintln(t *testing.T) {
	const name, age = "tom", 20
	val := fmt.Sprintln("name:", name, ", age:", age)
	fmt.Println(val)
	// output:
	// name:tom, age:20
	//
}

func TestSscan(t *testing.T) {
	s := "tom 20"
	var name string
	var age int
	n, err := fmt.Sscan(s, &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(name, age)
	// output:
	// 2
	// tom 20
}

func TestSscanf(t *testing.T) {
	s := "name: tom , age: 20" // 扫描参数两边需有空格
	var name string
	var age int
	n, err := fmt.Sscanf(s, "name: %s , age: %d", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(name, age)
	// output:
	// 2
	// tom 20
}

func TestSscanln(t *testing.T) {
	s := "tom 20\n"
	var name string
	var age int
	n, err := fmt.Sscanln(s, &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(name, age)
	// output:
	// 2
	// tom 20
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("name:%s, age:%d", u.Name, u.Age)
}

func TestString(t *testing.T) {
	u := User{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(u) // name:tom, age:20
}
