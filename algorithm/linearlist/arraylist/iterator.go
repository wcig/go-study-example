package arraylist

type ArrayIterator struct {
	list   *ArrayList
	cursor int
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
