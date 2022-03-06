package blockingqueue

import (
	"errors"
	"go-app/algorithm/linearlist"
)

var (
	InvalidSizeError = errors.New("invalid blocking queue size")
)

// 阻塞队列
type BlockingQueue struct {
	first *Element
	last  *Element
	size  int
	cap   int
}

type Element struct {
	value interface{}
	next  *Element
}

func New(capacity int) (*BlockingQueue, error) {
	if capacity <= 0 {
		return nil, InvalidSizeError
	}
	return &BlockingQueue{
		first: nil,
		last:  nil,
		size:  0,
		cap:   capacity,
	}, nil
}

func (q *BlockingQueue) Push(value interface{}) bool {
	if q.size >= q.cap {
		return false
	}

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
	return true
}

func (q *BlockingQueue) Pop() (value interface{}, ok bool) {
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

func (q *BlockingQueue) Peek() (value interface{}, ok bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.first.value, true
}

func (q *BlockingQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *BlockingQueue) Size() int {
	return q.size
}

func (q *BlockingQueue) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *BlockingQueue) Values() []interface{} {
	values := make([]interface{}, q.size, q.size)
	for i, e := 0, q.first; i < q.size; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

func (q *BlockingQueue) Iterator() linearlist.Iterator {
	return NewIterator(q)
}

func (q *BlockingQueue) get(index int) (interface{}, bool) {
	if !q.rangeCheck(index) {
		return nil, false
	}

	e := q.first
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e.value, true
}

func (q *BlockingQueue) rangeCheck(index int) bool {
	return index >= 0 && index < q.size
}
