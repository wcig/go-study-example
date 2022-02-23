package arraylist

const (
	defaultCapacity = 10

	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

// 数组结构线性表
type ArrayList struct {
	data []interface{}
	size int
}

func New() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0, defaultCapacity),
		size: 0,
	}
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) IsEmpty() bool {
	return l.Size() == 0
}

func (l *ArrayList) Clear() {
	l.data = make([]interface{}, 0, defaultCapacity)
	l.size = 0
}

func (l *ArrayList) Values() []interface{} {
	newData := make([]interface{}, l.size, l.size)
	copy(newData, l.data)
	return newData
}

func (l *ArrayList) Add(v interface{}) {
	l.grow(l.size + 1)
	l.data = append(l.data, v)
	l.size++
}

func (l *ArrayList) Insert(index int, v interface{}) bool {
	if !l.rangeCheck(index) {
		if index == l.size {
			l.Add(v)
			return true
		}
		return false
	}

	l.grow(l.size + 1)
	l.data = append(l.data, nil)
	for i := l.size; i > index; i-- {
		l.data[i] = l.data[i-1]
	}
	l.data[index] = v
	l.size++
	return true
}

func (l *ArrayList) Remove(index int) (interface{}, bool) {
	if !l.rangeCheck(index) {
		return nil, false
	}

	v := l.data[index]
	for i := index; i < l.size-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.size--
	l.data = l.data[:l.size]
	l.shrink()
	return v, true
}

func (l *ArrayList) Set(index int, v interface{}) bool {
	if !l.rangeCheck(index) {
		return false
	}

	l.data[index] = v
	return true
}

func (l *ArrayList) Get(index int) (interface{}, bool) {
	if !l.rangeCheck(index) {
		return nil, false
	}
	return l.data[index], true
}

func (l *ArrayList) Contain(v interface{}) bool {
	return l.IndexOf(v) > -1
}

func (l *ArrayList) IndexOf(v interface{}) int {
	if l.size == 0 {
		return -1
	}
	for i := 0; i < l.size; i++ {
		if l.data[i] == v {
			return i
		}
	}
	return -1
}

func (l *ArrayList) Iterator() *ArrayIterator {
	return &ArrayIterator{
		list:   l,
		cursor: 0,
	}
}

func (l *ArrayList) rangeCheck(index int) bool {
	return index >= 0 && index < l.size
}

func (l *ArrayList) grow(minCapacity int) {
	currentCapacity := cap(l.data)
	currentSize := len(l.data)
	if currentCapacity < minCapacity {
		newCapacity := int(float32(currentCapacity) * growthFactor)
		newData := make([]interface{}, currentSize, newCapacity)
		copy(newData, l.data)
		l.data = newData
	}
}

func (l *ArrayList) shrink() {
	currentCapacity := cap(l.data)
	currentSize := l.size
	newCapacity := int(float32(currentCapacity) * shrinkFactor)
	if currentSize <= newCapacity {
		newData := make([]interface{}, currentSize, newCapacity)
		copy(newData, l.data)
		l.data = newData
	}
}
