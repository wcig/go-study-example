package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	fmt.Printf("init: size: %d, empty: %t\n", h.Size(), h.Empty())

	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)
	fmt.Printf("after push: size: %d, empty: %t\n", h.Size(), h.Empty())

	top := h.Top()
	fmt.Println("top:", top) // 1

	fmt.Println(h.Pop()) // 1
	fmt.Println(h.Pop()) // 2
	fmt.Println(h.Pop()) // 3
	fmt.Println(h.Pop()) // 4
	fmt.Println(h.Pop()) // 5
	fmt.Printf("after pop: size: %d, empty: %t\n", h.Size(), h.Empty())
}

func TestMinHeapWithData(t *testing.T) {
	h := NewMinHeap([]int{5, 4, 3, 2, 1}...)
	fmt.Printf("init: size: %d, empty: %t\n", h.Size(), h.Empty())

	top := h.Top()
	fmt.Println("top:", top) // 1

	fmt.Println(h.Pop()) // 1
	fmt.Println(h.Pop()) // 2
	fmt.Println(h.Pop()) // 3
	fmt.Println(h.Pop()) // 4
	fmt.Println(h.Pop()) // 5
	fmt.Printf("after pop: size: %d, empty: %t\n", h.Size(), h.Empty())
}
