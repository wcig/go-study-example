package test

import (
	"fmt"
	"testing"
)

// struct方法定义
type foo struct{}

func (f foo) method() {
	//...
}

type userService struct{}

func (service *userService) method() {
	//...
}

// 相同的数据类型struct直接转换
type T1 struct {
	A int
	B int
}
type T2 struct {
	A int
	B int
}

func TestTransStruct(t *testing.T) {
	t1 := T1{
		A: 1,
		B: 2,
	}
	t2 := T2{
		A: t1.A,
		B: t1.B,
	}
	fmt.Println("t2:", t2)

	//可直接转换
	t3 := T2(t1)
	fmt.Println("t3:", t3)
}

// struct定义
type user struct {
	Id   int64
	Name string
}

func TestStructInit(t *testing.T) {
	u1 := user{10001, "Nick"}

	u2 := user{
		Id:   10002,
		Name: "Tom",
	}

	fmt.Println(u1, u2)
}
