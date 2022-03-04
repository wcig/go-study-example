package arraystack

import (
	"go-app/algorithm/linearlist"
)

type Iterator struct {
	stack  *ArrayStack
	cursor int
}

func NewIterator(s *ArrayStack) linearlist.Iterator {
	return &Iterator{
		stack:  s,
		cursor: 0,
	}
}

func (i *Iterator) HasNext() bool {
	return i.cursor < i.stack.Size()
}

func (i *Iterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	val, _ := i.stack.get(i.stack.Size() - i.cursor - 1)
	i.cursor++
	return val
}
