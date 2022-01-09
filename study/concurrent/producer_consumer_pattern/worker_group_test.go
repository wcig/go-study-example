package main

import (
	"fmt"
	"sync"
	"testing"
)

type WorkerGroup struct {
	job     func()
	workers int
}

func NewWorkerGroup(job func(), workers int) WorkerGroup {
	return WorkerGroup{
		job:     job,
		workers: workers,
	}
}

func (w WorkerGroup) Start() {
	var wg sync.WaitGroup
	wg.Add(w.workers)
	for i := 0; i < w.workers; i++ {
		go func() {
			defer wg.Done()
			w.job()
		}()
	}
	wg.Wait()
}

func TestWorkerGroup(t *testing.T) {
	w := NewWorkerGroup(func() {
		fmt.Println("ok")
	}, 4)
	w.Start()
}
