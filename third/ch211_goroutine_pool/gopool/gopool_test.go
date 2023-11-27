package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
)

func TestFirst(t *testing.T) {
	gopool.Go(func() {
		fmt.Println("gopool task")
	})
	time.Sleep(time.Millisecond)
}

func TestCustomPool(t *testing.T) {
	myPool := gopool.NewPool("myPool", 2, gopool.NewConfig())
	const num = 1000

	go func() {
		for {
			log.Printf("goroutines num: %d", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(num)

	var n int32 = 0
	var cn int32 = 0

	go func() {
		for {
			log.Printf(">> concurrent num: %d", atomic.LoadInt32(&cn))
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < num; i++ {
		myPool.Go(func() {
			atomic.AddInt32(&cn, 1)
			defer atomic.AddInt32(&cn, -1)

			atomic.AddInt32(&n, 1)
			time.Sleep(3 * time.Millisecond)
			wg.Done()
		})
	}

	wg.Wait()
	log.Println("result num:", n)

	//  Output:
	// 2023/11/27 19:43:04 >> concurrent num: 0
	// 2023/11/27 19:43:04 goroutines num: 6
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:04 goroutines num: 6
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:04 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 goroutines num: 6
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:05 goroutines num: 6
	// 2023/11/27 19:43:05 >> concurrent num: 2
	// 2023/11/27 19:43:06 >> concurrent num: 2
	// 2023/11/27 19:43:06 result num: 1000
}
