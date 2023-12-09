package heap

// 小顶堆
type MinHeap struct {
	data []int
}

func NewMinHeap(nums ...int) *MinHeap {
	h := &MinHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *MinHeap) Push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *MinHeap) up(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.data[p] <= h.data[i] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MinHeap) Pop() int {
	if h.Empty() {
		panic("min heap empty")
	}
	h.swap(0, h.Size()-1)
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return v
}

func (h *MinHeap) down(i int) {
	for {
		left, right, min := h.left(i), h.right(i), i
		if left < len(h.data) && h.data[left] < h.data[min] {
			min = left
		}
		if right < len(h.data) && h.data[right] < h.data[min] {
			min = right
		}
		if min == i {
			break
		}
		h.swap(i, min)
		i = min
	}
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Top() int {
	if h.Empty() {
		panic("min heap empty")
	}
	return h.data[0]
}

func (h *MinHeap) Empty() bool {
	return h.Size() == 0
}
func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) left(i int) int {
	return 2*i + 1
}

func (h *MinHeap) right(i int) int {
	return 2*i + 2
}
