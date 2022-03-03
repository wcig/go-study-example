package doublycircularlinkedlist

import "go-app/algorithm/linearlist"

type Iterator struct {
	list      *DoubleCircularLinkedList
	next      *Node
	nextIndex int
}

func NewIterator(list *DoubleCircularLinkedList) linearlist.Iterator {
	return &Iterator{
		list:      list,
		next:      list.first,
		nextIndex: 0,
	}
}

func (iterator *Iterator) HasNext() bool {
	return iterator.nextIndex < iterator.list.size
}

func (iterator *Iterator) Next() interface{} {
	if !iterator.HasNext() {
		return nil
	}

	val := iterator.next.value
	iterator.next = iterator.next.next
	iterator.nextIndex++
	return val
}
