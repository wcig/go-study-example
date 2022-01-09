package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// v4：使用context退出
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		n := i
		go worker(ctx, n, &wg)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("break", n)
			return
		default:
			fmt.Println("ok", n)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
