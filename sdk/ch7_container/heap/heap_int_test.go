package heap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// 自定义int类型堆
type IntHeap []int

// 实现heap.Interface的sort.Interface方法
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现heap.Interface的push方法
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
	*h = append(*h, x.(int))
}

// 实现heap.Interface的pop方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeapPushPop(t *testing.T) {
	// heap初始化
	h := &IntHeap{1, 2, 3}
	heap.Init(h)

	// heap插入元素
	heap.Push(h, 10)

	// heap推出元素
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// output:
	// minimum: 1
	// 1 2 3 10
}

func TestHeapRemoveFix(t *testing.T) {
	h := &IntHeap{2, 1, 3}
	heap.Init(h)
	fmt.Println(h)

	v := heap.Remove(h, 1)
	fmt.Println(v)
	fmt.Println(h)

	heap.Fix(h, 1)
	fmt.Println(h)
	// output:
	// &[1 2 3]
	// 2
	// &[1 3]
	// &[1 3]
}

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 3}
	fmt.Println(h)

	sort.Sort(h)
	fmt.Println(h)

	h.Push(10)
	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop().(int))
	}
	// output:
	// &[2 1 3]
	// &[1 2 3]
	// &[1 2 3 10]
	// 10 3 2 1
}
