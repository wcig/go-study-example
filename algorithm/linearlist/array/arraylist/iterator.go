package arraylist

import "go-app/algorithm/linearlist"

type ArrayIterator struct {
	list   *ArrayList
	cursor int
}

func NewIterator(list *ArrayList) linearlist.Iterator {
	return &ArrayIterator{
		list:   list,
		cursor: 0,
	}
}

func (iterator *ArrayIterator) HasNext() bool {
	if iterator.list.rangeCheck(iterator.cursor) {
		return true
	}
	return false
}

func (iterator *ArrayIterator) Next() interface{} {
	if iterator.list.rangeCheck(iterator.cursor) {
		val := iterator.list.data[iterator.cursor]
		iterator.cursor++
		return val
	}
	return nil
}
