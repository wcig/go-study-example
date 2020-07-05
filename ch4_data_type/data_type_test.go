package ch4_data_type

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

// 布尔类型
func TestBool(t *testing.T) {
	var f1 bool = true
	f2 := false

	fmt.Println(f1, f2)
}

// 整型
func TestInt(t *testing.T) {
	var i1 int = 0
	var i2 uint = 0

	var i3 int8 = 0
	var i4 uint8 = 0

	var i5 int16 = 0
	var i6 uint16 = 0

	var i7 int32 = 0
	var i8 uint32 = 0

	var i9 int64 = 0
	var i10 uint64 = 0

	fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10)

	var i11 = 0
	fmt.Println(reflect.TypeOf(i11)) // int
}

// 浮点型
func TestFloat(t *testing.T) {
	var f1 float32 = 3.14
	var f2 float64 = 3.14

	fmt.Println(f1, f2)

	var f3 = 3.14
	fmt.Println(reflect.TypeOf(f3)) // float64
}

// 复数类型
func TestComplex(t *testing.T) {
	var c1 complex64 = complex(1, 2)
	var c2 complex128 = complex(1, -2)

	fmt.Println(c1, c2)
	fmt.Println(c1 * c1)

	var c4 = complex(1, 2)
	fmt.Println(reflect.TypeOf(c4)) // complex128
}

// 字符串类型
func TestString(t *testing.T) {
	var str string = "今天天气很好"
	fmt.Println(len(str))                   // 字节长度
	fmt.Println(strings.Count(str, "") - 1) // 字符长度
}

// byte,rune类型
func TestByteRune(t *testing.T) {
	var str string = "今天天气很好"
	bytes := []byte(str)
	fmt.Println(len(bytes)) // 字节长度

	runes := []rune(str)
	fmt.Println(len(runes)) // 字符长度
}

// uintptr类型
func TestUintptr(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	u := user{
		name: "Tom",
		age:  6,
	}
	p := unsafe.Pointer(&u)

	pname := (*string)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(u.name)))
	page := (*int)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(u.age)))
	fmt.Println(*pname, *page)
}
