package simplehashtable

const (
	defaultCapacity = 1 << 4
	loadFactor      = 0.75
)

type SimpleHashTable struct {
	table    []*Element
	hashFunc HashFunc
	size     int
	cap      int
}

type Element struct {
	hash  int
	value interface{}
	next  *Element
}

type HashFunc func(v interface{}) int

func New(hashFunc HashFunc) *SimpleHashTable {
	return &SimpleHashTable{
		table:    make([]*Element, defaultCapacity, defaultCapacity),
		hashFunc: hashFunc,
		size:     0,
		cap:      defaultCapacity,
	}
}

// todo resize
func (t *SimpleHashTable) Add(v interface{}) {
	hash := t.hashFunc(v)
	newElement := &Element{
		hash:  hash,
		value: v,
		next:  nil,
	}

	index := t.getKeyTableIndex(hash)
	e := t.table[index]
	if e == nil {
		t.table[index] = newElement
		t.size++
		return
	}
	for {
		if e != nil && e.hash == newElement.hash && e.value == newElement.value {
			return
		}
		if e.next == nil {
			break
		}
		e = e.next
	}
	e.next = newElement
	t.size++
}

func (t *SimpleHashTable) Contain(v interface{}) bool {
	return t.getElement(t.hash(v), v) != nil
}

func (t *SimpleHashTable) Remove(v interface{}) bool {
	if t.size <= 0 {
		return false
	}

	hash := t.hash(v)
	index := t.getKeyTableIndex(hash)
	first := t.table[index]
	if first == nil {
		return false
	}
	var beforeElement *Element
	e := first
	var exist bool
	for e != nil {
		if e.hash == hash && e.value == v {
			exist = true
			break
		}
		beforeElement = e
		e = e.next
	}
	if exist {
		if beforeElement == nil {
			t.table[index] = e.next
		} else {
			beforeElement.next = e.next
		}
	}
	t.size--
	return true
}

func (t *SimpleHashTable) Size() int {
	return t.size
}

func (t *SimpleHashTable) IsEmpty() bool {
	return t.size == 0
}

func (t *SimpleHashTable) Clear() {
	if t.size > 0 {
		for i := 0; i < len(t.table); i++ {
			t.table[i] = nil
		}
		t.size = 0
	}
}

func (t *SimpleHashTable) getElement(hash int, v interface{}) *Element {
	if t.size <= 0 {
		return nil
	}

	index := t.getKeyTableIndex(hash)
	for e := t.table[index]; e != nil; e = e.next {
		if e.hash == hash && e.value == v {
			return e
		}
	}
	return nil
}

func (t *SimpleHashTable) hash(v interface{}) int {
	return t.hashFunc(v)
}

func (t *SimpleHashTable) getKeyTableIndex(hash int) int {
	return (t.cap - 1) & hash // 等价于: hash%t.cap
}
