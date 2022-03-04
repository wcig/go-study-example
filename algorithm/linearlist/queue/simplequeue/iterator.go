package simplequeue

import "go-app/algorithm/linearlist"

type Iterator struct {
	queue  *SimpleQueue
	cursor int
}

func NewIterator(q *SimpleQueue) linearlist.Iterator {
	return &Iterator{
		queue:  q,
		cursor: 0,
	}
}

func (i *Iterator) HasNext() bool {
	return i.cursor < i.queue.size
}

func (i *Iterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}

	val, _ := i.queue.get(i.cursor)
	i.cursor++
	return val
}
