package ants

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants"
)

func TestFirst(t *testing.T) {
	err := ants.Submit(func() {
		fmt.Println("ants task")
	})
	time.Sleep(time.Millisecond)
	fmt.Println(err)
}

func TestCustom(t *testing.T) {
	myPool, err := ants.NewPool(2)
	if err != nil {
		log.Fatalf("new ants pool err: %v", err)
	}
	defer myPool.Release()

	go func() {
		for {
			log.Printf("goroutines num: %d", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	const num = 1000
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
		err2 := myPool.Submit(func() {
			atomic.AddInt32(&cn, 1)
			defer atomic.AddInt32(&cn, -1)

			atomic.AddInt32(&n, 1)
			time.Sleep(3 * time.Millisecond)
			wg.Done()
		})
		if err2 != nil {
			log.Printf("ants pool task exec err: %v", err)
		}
	}

	wg.Wait()
	log.Println("result num:", n)

	// Output:
	// 2023/11/27 19:43:27 >> concurrent num: 1
	// 2023/11/27 19:43:27 goroutines num: 8
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 goroutines num: 8
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:28 goroutines num: 8
	// 2023/11/27 19:43:28 >> concurrent num: 2
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 goroutines num: 8
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 >> concurrent num: 2
	// 2023/11/27 19:43:29 result num: 1000
}
