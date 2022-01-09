package main

import (
	"fmt"
	"time"
)

// v1：退出一个goroutine
func main() {
	cancel := make(chan bool)
	go worker(cancel)

	time.Sleep(time.Second)
	cancel <- true
}

func worker(cancel chan bool) {
	for {
		select {
		case <-cancel:
			break
		default:
			fmt.Println("ok")
		}
		time.Sleep(time.Millisecond * 100)
	}
}
