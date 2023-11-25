package main

import (
	"fmt"
	"unsafe"
)

func main() {
	foo1 := Foo1{}
	fmt.Println(unsafe.Sizeof(foo1)) // 24

	foo2 := Foo2{}
	fmt.Println(unsafe.Sizeof(foo2)) // 16
}

// 1-8-1: 总共占用24字节
type Foo1 struct {
	b1 byte
	i  int64
	b2 byte
}

func sum1(foos []Foo1) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

// 8-1-1: 总共占用16字节
type Foo2 struct {
	i  int64
	b1 byte
	b2 byte
}

func sum2(foos []Foo2) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}
