package singlelinklist

type Iterator struct {
	list      *SingleLinkList
	next      *Node
	nextIndex int
}

func NewIterator(list *SingleLinkList) *Iterator {
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

	val := iterator.next.data
	iterator.next = iterator.next.next
	iterator.nextIndex++
	return val
}
