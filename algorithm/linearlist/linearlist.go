package linearlist

type LinearList interface {
	Size() int
	IsEmpty() bool
	Clear()
	Values() []interface{}

	Add(v interface{})
	Insert(index int, v interface{}) bool
	Remove(index int) (interface{}, bool)
	Set(index int, v interface{}) bool
	Get(index int) (interface{}, bool)
	Contain(v interface{}) bool
	IndexOf(v interface{}) int
	Iterator() Iterator
}
