package linkedstack

import (
	"go-app/algorithm/linearlist"
)

type Iterator struct {
	stack  *LinkedStack
	cursor int
}

func NewIterator(s *LinkedStack) linearlist.Iterator {
	return &Iterator{
		stack:  s,
		cursor: 0,
	}
}

func (i *Iterator) HasNext() bool {
	return i.cursor < i.stack.size
}

func (i *Iterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	val, _ := i.stack.get(i.cursor)
	i.cursor++
	return val
}
