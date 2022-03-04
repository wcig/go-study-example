package simplequeue

type SimpleQueue struct {
	first *Element
	last  *Element
	size  int
}

type Element struct {
	value interface{}
	next  *Element
}

func (q *SimpleQueue) Push(value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (q *SimpleQueue) Pop() (value interface{}, ok bool) {
	// TODO implement me
	panic("implement me")
}

func (q *SimpleQueue) Peek() (value interface{}, ok bool) {
	// TODO implement me
	panic("implement me")
}

func (q *SimpleQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *SimpleQueue) Size() int {
	return q.size
}

func (q *SimpleQueue) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *SimpleQueue) Values() []interface{} {
	values := make([]interface{}, q.size, q.size)
	for i, e := 0, q.first; i < q.size; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}
