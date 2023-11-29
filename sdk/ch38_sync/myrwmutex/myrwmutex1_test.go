package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

// 读优先读写锁: 获取写锁不到自旋
type MyRWMutex1 struct {
	mu          sync.Mutex
	readerCount int
}

func (rw *MyRWMutex1) RLock() {
	rw.mu.Lock()
	rw.readerCount++
	rw.mu.Unlock()
}

func (rw *MyRWMutex1) RUnlock() {
	rw.mu.Lock()
	rw.readerCount--
	rw.mu.Unlock()
}

func (rw *MyRWMutex1) Lock() {
	// 获取不到写锁自旋
	for {
		rw.mu.Lock()
		if rw.readerCount > 0 {
			rw.mu.Unlock()
		} else {
			break
		}
	}
}

func (rw *MyRWMutex1) Unlock() {
	rw.mu.Unlock()
}

// ------------------------------------------------------ //

type MyRWMutexCounter1 struct {
	mu      MyRWMutex1
	counter int
}

func (c *MyRWMutexCounter1) Incr(n int) {
	c.mu.Lock()
	c.counter += n
	c.mu.Unlock()
}

func (c *MyRWMutexCounter1) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}

func TestMyRWMutex1(t *testing.T) {
	const n = 1000
	wg := &sync.WaitGroup{}
	wg.Add(n)

	c := &MyRWMutexCounter1{}
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
