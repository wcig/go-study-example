package ch21_generic

import (
	"fmt"
	"testing"
)

// 泛型方法
type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Push(v T) {
	q.elements = append(q.elements, v)
}

func (q *Queue[T]) Pop() T {
	var value T
	if len(q.elements) == 0 {
		return value
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func TestQueue(t *testing.T) {
	var q Queue[int]
	fmt.Println(q.Size())
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Size())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

// 定义多个类型形参的泛型方法
type MyFoo[T string, K int | int32 | int64] struct {
	Name  T
	Value K
}

func (f *MyFoo[T, K]) Print() {
	fmt.Printf("Foo: %s-%d\n", f.Name, f.Value)
}

func TestMethod(t *testing.T) {
	f := &MyFoo[string, int]{"tom", 123}
	f.Print()
}
