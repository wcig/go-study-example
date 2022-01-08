package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// v3: panic捕获
func main() {
	num := 1000

	var wg sync.WaitGroup
	wg.Add(num)
	start := time.Now()

	go printNumGoroutine()

	for i := 0; i < num; i++ {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("panic:", p)
				}
			}()

			defer wg.Done()
			_ = requestWork(context.Background(), "ok")
		}()
	}

	wg.Wait()
	fmt.Println("time cost:", time.Since(start))
	time.Sleep(10 * time.Second)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
	// Output:
	// 1 4
	// 2 2002
	// time cost: 2.004640673s
	// 3 1002
	// 4 1002
	// 5 1002
	// 6 2
	// 7 2
	// 8 2
	// 9 2
	// 10 2
	// 11 2
	// 12 2
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

	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

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
