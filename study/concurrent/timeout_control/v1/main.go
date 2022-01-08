package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// v1: 内存泄露版本
// 原因：requestWork因超时执行完成，此时done channel没有go routine接收，done <- hardWork(job)将一直卡主占用一个goroutine
func main() {
	num := 1000

	var wg sync.WaitGroup
	wg.Add(num)
	start := time.Now()

	go printNumGoroutine()

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			_ = requestWork(context.Background(), "ok")
		}()
	}

	wg.Wait()
	fmt.Println("time cost:", time.Since(start))
	time.Sleep(10 * time.Second)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
	// Output:
	// 1 6
	// 2 2002
	// 3 1893
	// time cost: 2.004288934s
	// 4 1002
	// 5 1002
	// 6 1002
	// 7 1002
	// 8 1002
	// 9 1002
	// 10 1002
	// 11 1002
	// 12 1002
}

func printNumGoroutine() {
	for i := 1; ; i++ {
		fmt.Println(i, runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan error)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func hardWork(job interface{}) error {
	time.Sleep(5 * time.Second)
	return nil
}
