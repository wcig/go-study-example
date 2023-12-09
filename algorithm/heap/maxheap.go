package heap

// 大顶堆
type MaxHeap struct {
	data []int
}

func NewMaxHeap(nums ...int) *MaxHeap {
	h := &MaxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *MaxHeap) Push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *MaxHeap) up(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.data[p] >= h.data[i] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MaxHeap) Pop() int {
	if h.Empty() {
		panic("max heap empty")
	}
	h.swap(0, h.Size()-1)
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return v
}

func (h *MaxHeap) down(i int) {
	for {
		left, right, max := h.left(i), h.right(i), i
		if left < len(h.data) && h.data[left] > h.data[max] {
			max = left
		}
		if right < len(h.data) && h.data[right] > h.data[max] {
			max = right
		}
		if max == i {
			break
		}
		h.swap(i, max)
		i = max
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Top() int {
	if h.Empty() {
		panic("max heap empty")
	}
	return h.data[0]
}

func (h *MaxHeap) Empty() bool {
	return h.Size() == 0
}
func (h *MaxHeap) Size() int {
	return len(h.data)
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) left(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) right(i int) int {
	return 2*i + 2
}
