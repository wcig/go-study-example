package simplequeue

import "go-app/algorithm/linearlist"

// 普通队列
type SimpleQueue struct {
	first *Element
	last  *Element
	size  int
}

type Element struct {
	value interface{}
	next  *Element
}

func New() *SimpleQueue {
	return &SimpleQueue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *SimpleQueue) Push(value interface{}) {
	newElement := &Element{
		value: value,
		next:  nil,
	}

	if q.size == 0 {
		q.first = newElement
		q.last = newElement
	} else {
		q.last.next = newElement
		q.last = newElement
	}
	q.size++
}

func (q *SimpleQueue) Pop() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}

	element := q.first
	value = element.value
	q.first = q.first.next
	q.size--
	element = nil
	return value, true
}

func (q *SimpleQueue) Peek() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.first.value, true
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

func (q *SimpleQueue) Iterator() linearlist.Iterator {
	return NewIterator(q)
}

func (q *SimpleQueue) get(index int) (interface{}, bool) {
	if !q.rangeCheck(index) {
		return nil, false
	}

	e := q.first
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e.value, true
}

func (q *SimpleQueue) rangeCheck(index int) bool {
	return index >= 0 && index < q.size
}
