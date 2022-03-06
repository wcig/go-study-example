package simplequeue

// type Iterator struct {
// 	queue  *ConcurrentQueue
// 	cursor int
// }
//
// func NewIterator(q *ConcurrentQueue) linearlist.Iterator {
// 	return &Iterator{
// 		queue:  q,
// 		cursor: 0,
// 	}
// }
//
// func (i *Iterator) HasNext() bool {
// 	return i.cursor < i.queue.size
// }
//
// func (i *Iterator) Next() interface{} {
// 	if !i.HasNext() {
// 		return nil
// 	}
//
// 	val, _ := i.queue.get(i.cursor)
// 	i.cursor++
// 	return val
// }
