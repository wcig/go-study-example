package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const num = 5
	limiter := NewSimpleLimiter(2)
	task := func(i int) {
		time.Sleep(100 * time.Millisecond)
	}
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func(n int) {
			start := time.Now()
			fmt.Println("add task:", n, start)
			defer wg.Done()
			limiter.enter()
			defer limiter.leave()
			task(n)
			fmt.Println("over task:", n, start, time.Since(start))
		}(i)
	}
	wg.Wait()
	// Output:
	// add task: 4 2022-01-09 19:54:00.804498 +0800 CST m=+0.000172323
	// add task: 2 2022-01-09 19:54:00.804464 +0800 CST m=+0.000138474
	// add task: 1 2022-01-09 19:54:00.804474 +0800 CST m=+0.000147857
	// add task: 5 2022-01-09 19:54:00.804455 +0800 CST m=+0.000129177
	// add task: 3 2022-01-09 19:54:00.804536 +0800 CST m=+0.000210537
	// over task: 2 2022-01-09 19:54:00.804464 +0800 CST m=+0.000138474 104.93371ms
	// over task: 4 2022-01-09 19:54:00.804498 +0800 CST m=+0.000172323 104.948548ms
	// over task: 5 2022-01-09 19:54:00.804455 +0800 CST m=+0.000129177 205.034631ms
	// over task: 1 2022-01-09 19:54:00.804474 +0800 CST m=+0.000147857 205.007855ms
	// over task: 3 2022-01-09 19:54:00.804536 +0800 CST m=+0.000210537 305.092622ms
}

// concurrent limiter
type SimpleLimiter struct {
	sem chan struct{}
}

func NewSimpleLimiter(num int64) *SimpleLimiter {
	return &SimpleLimiter{sem: make(chan struct{}, num)}
}

func (sl *SimpleLimiter) enter() {
	sl.sem <- struct{}{}
}

func (sl *SimpleLimiter) leave() {
	<-sl.sem
}
