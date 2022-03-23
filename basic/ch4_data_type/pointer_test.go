package ch4_data_type

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	// 只声明没有初始化
	var p1 *int
	assert.Nil(t, p1)

	// new声明并初始化
	p2 := new(int)
	assert.NotNil(t, p2)
	assert.Equal(t, 0, *p2)

	// 指向变量地址声明并初始化
	num := 1
	p3 := &num
	assert.NotNil(t, p3)
	assert.Equal(t, 1, *p3)
}

func TestPointer2(t *testing.T) {
	// 只声明没有初始化方式-同一个函数内: 1)不能通过解地址修改其指向的值;2)可以通过修改指针地址方式修改值
	var p1 *int
	// *p1 = 1 // panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	num1 := 1
	p1 = &num1
	fmt.Printf("%p, %v\n", p1, *p1)

	// 声明并初始化-同一个函数内: 1)可以通过解地址修改其指向的值;2)可以通过修改指针地址方式修改值
	p2 := new(int)
	fmt.Printf("%p, %v\n", p2, *p2)
	*p2 = 2
	fmt.Printf("%p, %v\n", p2, *p2)
	num2 := 200
	p2 = &num2
	fmt.Printf("%p, %v\n", p2, *p2)

	// Output:
	// 0xc0000acb00, 1
	// 0xc0000acb08, 0
	// 0xc0000acb08, 2
	// 0xc0000acb10, 200
}

func TestPointer3(t *testing.T) {
	fn1 := func(p *int) {
		n := 1
		p = &n
	}
	fn2 := func(p *int) {
		*p = 2
	}
	fn3 := func(p **int) {
		n := 3
		*p = &n
	}

	// 只声明没有初始化方式-跨函数
	var p1 *int
	fn1(p1)
	fmt.Println(p1) // <nil>
	// fn2(p1)
	// fmt.Println(p1) // panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	fn3(&p1)
	fmt.Println(p1, *p1) // 0xc0000acb00 3

	// 声明并初始化-跨函数
	p2 := new(int)
	fmt.Println(p2, *p2) // 0xc0000acb08 0
	fn1(p2)
	fmt.Println(p2, *p2) // 0xc0000acb08 0
	fn2(p2)
	fmt.Println(p2, *p2) // 0xc0000acb08 2
	fn3(&p2)
	fmt.Println(p2, *p2) // 0xc0000acb10 3
}

type user struct {
	Id   int
	Name string
}

// 注意修改指针的方式
func TestPointerModifyValue(t *testing.T) {
	u := &user{1, "tom"}
	fmt.Printf("%p, %v\n", u, u)
	modify1(u)
	fmt.Printf("%p, %v\n", u, u)
	modify2(u)
	fmt.Printf("%p, %v\n", u, u)
	modify3(u)
	fmt.Printf("%p, %v\n", u, u)

	u.Id = 100
	fmt.Printf("%p, %v\n", u, u)
	*u = user{101, "101"}
	fmt.Printf("%p, %v\n", u, u)
	u = &user{102, "102"}
	fmt.Printf("%p, %v\n", u, u)
	// Output:
	// 0xc0000be048, &{1 tom}
	// modify1 before: 0xc0000be048, &{1 tom}
	// modify1 after: 0xc0000be048, &{2 tom}
	// 0xc0000be048, &{2 tom}
	// modify2 before: 0xc0000be048, &{2 tom}
	// modify2 after: 0xc0000be048, &{3 modify ok}
	// 0xc0000be048, &{3 modify ok}
	// modify3 before: 0xc0000be048, &{3 modify ok}
	// modify3 after: 0xc0000be120, &{4 modify false}
	// 0xc0000be048, &{3 modify ok}
	// 0xc0000be048, &{100 modify ok}
	// 0xc0000be048, &{101 101}
	// 0xc0000be198, &{102 102}
}

// 直接修改结构体指针的域-生效
func modify1(u *user) {
	fmt.Printf("modify1 before: %p, %v\n", u, u)
	u.Id = 2
	fmt.Printf("modify1 after: %p, %v\n", u, u)
}

// 修改结构体指针指向的值-生效
func modify2(u *user) {
	fmt.Printf("modify2 before: %p, %v\n", u, u)
	*u = user{3, "modify ok"}
	fmt.Printf("modify2 after: %p, %v\n", u, u)
}

// 修改结构体指针的地址-生效
func modify3(u *user) {
	fmt.Printf("modify3 before: %p, %v\n", u, u)
	u = &user{4, "modify false"}
	fmt.Printf("modify3 after: %p, %v\n", u, u)
}

func TestNewPointer(t *testing.T) {
	u := new(*user)
	fmt.Println(u, *u)     // 0xc000010060 <nil>
	fmt.Println(u == nil)  // false
	fmt.Println(*u == nil) // true
}
