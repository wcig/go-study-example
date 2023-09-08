package ch21_generic

import (
	"fmt"
	"testing"
)

// 泛型struct
func TestStruct(t *testing.T) {
	type Foo[T int | int32 | int64] struct {
		Name  string
		Value T
	}

	var f1 Foo[int]
	f1.Name = "aaa"
	f1.Value = 111
	fmt.Println(f1)

	f2 := &Foo[int64]{
		Name:  "bbb",
		Value: 222,
	}
	fmt.Println(f2)
}

// 定义多个类型形参的泛型struct
func TestMultiTypeParamStruct(t *testing.T) {
	type Foo[T string, K int | int32 | int64] struct {
		First  T
		Second K
	}
	f := &Foo[string, int]{
		First:  "tom",
		Second: 123,
	}
	fmt.Println(f)
}

// 泛型不支持匿名struct
func TestAnonymousStruct(t *testing.T) {
	// error
	// testCase := struct[T int | int32 | int64] {
	// 	Input    T
	// 	Expected T
	// }{
	// 	Input:    111,
	// 	Expected: 222,
	// }
}
