package concurrentblockingqueue

import "errors"

var (
	InvalidSizeError = errors.New("invalid concurrent blocking queue size")
)

// 并发阻塞队列 TODO
type ConcurrentBlockingQueue struct {
	queue chan interface{}
}

func New(capacity int) *ConcurrentBlockingQueue {
	return &ConcurrentBlockingQueue{
		queue: make(chan interface{}, capacity),
	}
}

func (q *ConcurrentBlockingQueue) Push(value interface{}) bool {
	select {
	case q.queue <- value:
		return true
	default:
		return false
	}
}

func (q *ConcurrentBlockingQueue) Pop() (value interface{}, ok bool) {
	select {
	case value = <-q.queue:
		return value, true
	default:
		return nil, false
	}
}

func (q *ConcurrentBlockingQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ConcurrentBlockingQueue) Size() int {
	return len(q.queue)
}
