package main

import (
	"fmt"
	"time"
)

// v2：退出多个goroutine
func main() {
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
}

func worker(cancel chan bool) {
	// 注意break只作用于select，此处应使用return
	for {
		select {
		case <-cancel:
			return
		default:
			fmt.Println("ok")
		}
		time.Sleep(time.Millisecond * 100)
	}
}
