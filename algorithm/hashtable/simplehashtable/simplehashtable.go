package simplehashtable

const (
	defaultSize = 1 << 4
	loadFactor  = 0.75
)

type SimpleHashTable struct {
	table []*Element
	hash  Hash
	len   int
	cap   int
}

type Element struct {
	value interface{}
	next  *Element
}

type Hash func(v interface{}, tableSize int) int

func (t *SimpleHashTable) Add(v interface{}) {}

func (t *SimpleHashTable) Contain(v interface{}) bool { return false }

func (t *SimpleHashTable) Remove(v interface{}) {}

func (t *SimpleHashTable) Size() int { return 0 }

func (t *SimpleHashTable) IsEmpty() bool { return false }

func (t *SimpleHashTable) Clear() {}
