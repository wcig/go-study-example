package doubleendedqueue

import "go-app/algorithm/linearlist"

// 双端队列
type DoubleEndedQueue struct {
	first *Element
	last  *Element
	size  int
}

type Element struct {
	value interface{}
	prev  *Element
	next  *Element
}

func New() *DoubleEndedQueue {
	return &DoubleEndedQueue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *DoubleEndedQueue) PushBack(value interface{}) {
	newElement := &Element{
		value: value,
		prev:  q.last,
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

func (q *DoubleEndedQueue) PushFront(value interface{}) {
	newElement := &Element{
		value: value,
		prev:  nil,
		next:  q.first,
	}

	if q.size == 0 {
		q.first = newElement
		q.last = newElement
	} else {
		q.first.prev = newElement
		q.first = newElement
	}
	q.size++
}

func (q *DoubleEndedQueue) PopBack() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}

	value = q.last.value
	if q.last.prev == nil {
		q.first = nil
		q.last = nil
	} else {
		q.last.prev.next = nil
		q.last = q.last.prev
	}
	q.size--
	return value, true
}

func (q *DoubleEndedQueue) PopFront() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}

	value = q.first.value
	if q.first.next == nil {
		q.first = nil
		q.last = nil
	} else {
		q.first.next.prev = nil
		q.first = q.first.next
	}
	q.size--
	return value, true
}

func (q *DoubleEndedQueue) PeekBack() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.last.value, true
}

func (q *DoubleEndedQueue) PeekFront() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.first.value, true
}

func (q *DoubleEndedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *DoubleEndedQueue) Size() int {
	return q.size
}

func (q *DoubleEndedQueue) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *DoubleEndedQueue) Values() []interface{} {
	values := make([]interface{}, q.size, q.size)
	for i, e := 0, q.first; i < q.size; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

func (q *DoubleEndedQueue) Iterator() linearlist.Iterator {
	return NewIterator(q)
}

func (q *DoubleEndedQueue) rangeCheck(index int) bool {
	return index >= 0 && index < q.size
}

func (q *DoubleEndedQueue) get(index int) (interface{}, bool) {
	if !q.rangeCheck(index) {
		return nil, false
	}

	e := q.first
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e.value, true
}
