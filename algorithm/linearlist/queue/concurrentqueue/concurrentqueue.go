package simplequeue

import (
	"sync"
)

// 普通队列
type ConcurrentQueue struct {
	mu    sync.Mutex
	first *Element
	last  *Element
	size  int
}

type Element struct {
	value interface{}
	next  *Element
}

func New() *ConcurrentQueue {
	return &ConcurrentQueue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *ConcurrentQueue) Push(value interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

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

func (q *ConcurrentQueue) Pop() (value interface{}, ok bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

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

func (q *ConcurrentQueue) Peek() (value interface{}, ok bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		return nil, false
	}
	return q.first.value, true
}

func (q *ConcurrentQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ConcurrentQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.size
}

func (q *ConcurrentQueue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *ConcurrentQueue) Values() []interface{} {
	values := make([]interface{}, q.size, q.size)
	for i, e := 0, q.first; i < q.size; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

// func (q *ConcurrentQueue) Iterator() linearlist.Iterator {
// 	return NewIterator(q)
// }
//
// func (q *ConcurrentQueue) get(index int) (interface{}, bool) {
// 	if !q.rangeCheck(index) {
// 		return nil, false
// 	}
//
// 	e := q.first
// 	for i := 0; i < index; i++ {
// 		e = e.next
// 	}
// 	return e.value, true
// }
//
// func (q *ConcurrentQueue) rangeCheck(index int) bool {
// 	return index >= 0 && index < q.size
// }
