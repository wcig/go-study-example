package ch14_concurrent_pattern

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// 生产者-消费者模式
func TestProducerConsumer(t *testing.T) {
	ch := make(chan int, 10)
	go producer(2, ch)
	go producer(4, ch)
	go consumer(ch)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func producer(factor int, ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i * factor
	}
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
