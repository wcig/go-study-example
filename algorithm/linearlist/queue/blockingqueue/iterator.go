package blockingqueue

import "go-app/algorithm/linearlist"

type Iterator struct {
	queue  *BlockingQueue
	cursor int
}

func NewIterator(q *BlockingQueue) linearlist.Iterator {
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
