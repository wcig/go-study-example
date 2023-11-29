package main

import (
	"fmt"
	"sync"
	"testing"
)

type MyMutex struct {
	ch chan struct{}
}

func NewMyMutex() *MyMutex {
	return &MyMutex{ch: make(chan struct{}, 1)}
}

func (m *MyMutex) Lock() {
	m.ch <- struct{}{}
}

func (m *MyMutex) Unlock() {
	<-m.ch
}

func TestMyMutex(t *testing.T) {
	const n = 1000
	wg := &sync.WaitGroup{}
	wg.Add(n)

	num := 0
	mu := NewMyMutex()

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				mu.Lock()
				num++
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(num == n*n)
}
