package main

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 写优先读写锁: 使用chan替换底层信号量
type MyRWMutex2 struct {
	mu          sync.Mutex
	writerSem   chan struct{}
	readerSem   chan struct{}
	readerCount atomic.Int32
	readerWait  atomic.Int32
}

const maxReaders = 1 << 30

func NewMyRWMutex2() *MyRWMutex2 {
	return &MyRWMutex2{
		writerSem: make(chan struct{}),
		readerSem: make(chan struct{}),
	}
}

func (rw *MyRWMutex2) RLock() {
	if rw.readerCount.Add(1) < 0 {
		<-rw.readerSem
	}
}

func (rw *MyRWMutex2) RUnlock() {
	if rw.readerCount.Add(-1) < 0 {
		if rw.readerWait.Add(-1) == 0 {
			rw.writerSem <- struct{}{}
		}
	}
}

func (rw *MyRWMutex2) Lock() {
	rw.mu.Lock()
	r := rw.readerCount.Add(-maxReaders) + maxReaders
	if r != 0 && rw.readerWait.Add(r) != 0 {
		<-rw.writerSem
	}
}

func (rw *MyRWMutex2) Unlock() {
	r := rw.readerCount.Add(maxReaders)
	for i := 0; i < int(r); i++ {
		rw.readerSem <- struct{}{}
	}
	rw.mu.Unlock()
}

// ------------------------------------------------------ //

type MyRWMutexCounter2 struct {
	mu      *MyRWMutex2
	counter int
}

func NewMyRWMutexCounter2() *MyRWMutexCounter2 {
	return &MyRWMutexCounter2{
		mu:      NewMyRWMutex2(),
		counter: 0,
	}
}

func (c *MyRWMutexCounter2) Incr(n int) {
	c.mu.Lock()
	c.counter += n
	c.mu.Unlock()
}

func (c *MyRWMutexCounter2) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}

func TestMyRWMutex2(t *testing.T) {
	const n = 1000
	wg := &sync.WaitGroup{}
	wg.Add(n)

	c := NewMyRWMutexCounter2()
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				c.Incr(1)
			}
			wg.Done()
		}()
	}

	go func() {
		for {
			_ = c.Value()
			time.Sleep(time.Microsecond)
		}
	}()

	wg.Wait()
	log.Println(">> end:", c.Value() == n*n)
}
