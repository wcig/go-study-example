package main

import (
	"fmt"
	"sync"
	"time"
)

// v3：所有任务处理完成 + 退出多个goroutine
func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		n := i
		go worker(n, &wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}

func worker(n int, wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		case <-cancel:
			fmt.Println("break", n)
			return
		default:
			fmt.Println("ok", n)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
