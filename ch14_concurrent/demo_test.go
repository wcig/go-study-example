package ch14_concurrent

import (
	"fmt"
	"sync"
	"testing"
)

// 顺序一致性问题：在同一个goroutine可以保证，在不同goroutine之间无法保证
var a string
var done bool

func setup() {
	a = "hello world"
	done = true
}

func TestGoroutineSequence(t *testing.T) {
	go setup() // 在main goroutine看来，setup goroutine中a的赋值操作可能在done赋值之后
	for !done {
	}
	fmt.Println(a)
}

func TestSeq1(t *testing.T) {
	go println("ok")
}

func TestSeq2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("ok")
	}()
	wg.Wait()
}

func TestSeq3(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		println("ok")
		ch <- struct{}{}
	}()
	<-ch
}

func TestSeq4(t *testing.T) {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		println("ok")
		mu.Unlock()
	}()
	mu.Lock()
}
