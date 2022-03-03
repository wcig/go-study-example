package singlycircularlinkedlist

import "go-app/algorithm/linearlist"

type Iterator struct {
	list      *SinglyCircularLinkedList
	next      *Node
	nextIndex int
}

func NewIterator(list *SinglyCircularLinkedList) linearlist.Iterator {
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
