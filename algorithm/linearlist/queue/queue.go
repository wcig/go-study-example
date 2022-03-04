package queue

type Queue interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	IsEmpty() bool
	Size() int
	Clear()
	Values() []interface{}
}
