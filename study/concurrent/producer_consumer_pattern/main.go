package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者-消费者模式
func main() {
	ch := make(chan int, 64)

	go producer(3, ch)
	go producer(5, ch)
	go consumer(ch)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
		time.Sleep(200 * time.Millisecond)
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
