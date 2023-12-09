package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	fmt.Printf("init: size: %d, empty: %t\n", h.Size(), h.Empty())

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)
	fmt.Printf("after push: size: %d, empty: %t\n", h.Size(), h.Empty())

	top := h.Top()
	fmt.Println("top:", top) // 1

	fmt.Println(h.Pop()) // 5
	fmt.Println(h.Pop()) // 4
	fmt.Println(h.Pop()) // 3
	fmt.Println(h.Pop()) // 2
	fmt.Println(h.Pop()) // 1
	fmt.Printf("after pop: size: %d, empty: %t\n", h.Size(), h.Empty())
}

func TestMaxHeapWithData(t *testing.T) {
	h := NewMaxHeap([]int{1, 2, 3, 4, 5}...)
	fmt.Printf("init: size: %d, empty: %t\n", h.Size(), h.Empty())

	top := h.Top()
	fmt.Println("top:", top) // 1

	fmt.Println(h.Pop()) // 5
	fmt.Println(h.Pop()) // 4
	fmt.Println(h.Pop()) // 3
	fmt.Println(h.Pop()) // 2
	fmt.Println(h.Pop()) // 1
	fmt.Printf("after pop: size: %d, empty: %t\n", h.Size(), h.Empty())
}
