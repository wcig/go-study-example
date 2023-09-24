package ch21_generic

import (
	"fmt"
	"testing"
)

// 基本接口(Basic interfaces): go1.18支持泛型前, interface为方法集合
type FooBar interface {
	Foo(n int)
	Bar(n int)
}

// go1.18支持泛型后, interface可以为类型集合
type Numeric interface {
	int | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func TestNumeric(t *testing.T) {
	s1 := Sum(111, 222)
	fmt.Println(s1)
}

func Sum[T Numeric](a T, b T) T {
	return a + b
}

// 一般接口(General interface): 混合方法和类型集合 (一般不常用)
type ReadWriter interface {
	string | []rune
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// ~: 指定底层类型
func Test1(t *testing.T) {
	type Int interface {
		int | int8 | int16 | int32 | int64
	}

	type Uint interface {
		uint | uint8 | uint16 | uint32
	}

	type Float interface {
		float32 | float64
	}

	type MySlice[T Int | Uint | Float] []T

	// 正确
	var s1 MySlice[int]
	_ = s1

	// 错误
	// type MyInt int
	// var s2 MySlice[MyInt]
	// _ = s2
}

func Test2(t *testing.T) {
	type Int interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	type Uint interface {
		~uint | ~uint8 | ~uint16 | ~uint32
	}

	type Float interface {
		~float32 | ~float64
	}

	type MySlice[T Int | Uint | Float] []T

	var s1 MySlice[int]
	_ = s1

	type MyInt int
	var s2 MySlice[MyInt]
	_ = s2

	type MyMyInt MyInt
	var s3 MySlice[MyMyInt]
	_ = s3

	// 错误
	// var s4 MySlice[Int]
	// _ = s4
}

// any: 任何类型 (type any = interface{})
func TestAny(t *testing.T) {
	type Node[T any] struct {
		elements []T
	}

	n1 := &Node[int]{}
	_ = n1

	n2 := &Node[float64]{}
	_ = n2

	n3 := &Node[string]{}
	_ = n3

	type User struct{}
	n4 := &Node[*User]{}
	_ = n4
}
