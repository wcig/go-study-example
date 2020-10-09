package ch14_concurrent_pattern

import (
	"fmt"
	"testing"
)

// 控制并发数

// 控制并发数示例
var limit = make(chan bool, 3)

func TestControlConcurrentNum1(t *testing.T) {
	work := func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println("over")
	}
	works := []func(){work, work, work}
	for _, w := range works {
		limit <- true
		w()
		<-limit
	}
	select {}
}

// 控制并发数抽象
func TestControlConcurrentNum2(t *testing.T) {
	work := func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println("over")
	}
	works := []func(){work, work, work}

	limit := 3
	gate := make(chan struct{}, limit)
	enter := func() { gate <- struct{}{} }
	leave := func() { <-gate }

	for _, w := range works {
		enter()
		w()
		leave()
	}
	select {}
}
