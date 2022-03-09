package simplehashtable

const (
	defaultCapacity = 1 << 4
	loadFactor      = 0.75
)

type SimpleHashTable struct {
	table []*Element
	hash  Hash
	size  int
	cap   int
}

type Element struct {
	value interface{}
	next  *Element
}

type Hash func(v interface{}, tableSize int) int

func New(hash Hash) *SimpleHashTable {
	return &SimpleHashTable{
		table: make([]*Element, defaultCapacity, defaultCapacity),
		hash:  hash,
		size:  0,
		cap:   defaultCapacity,
	}
}

func (t *SimpleHashTable) Add(v interface{}) {
	ae := &Element{
		value: v,
		next:  nil,
	}

	hash := t.hash(v, t.cap)
	e := t.table[hash]
	if e == nil {
		t.table[hash] = ae
		return
	}

	for {
		e = e.next
	}
}

func (t *SimpleHashTable) Contain(v interface{}) bool { return false }

func (t *SimpleHashTable) Remove(v interface{}) {}

func (t *SimpleHashTable) Size() int { return 0 }

func (t *SimpleHashTable) IsEmpty() bool { return false }

func (t *SimpleHashTable) Clear() {}
