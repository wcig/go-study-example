package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestForBreak(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		if i == 2 {
			break
		}
	}
	fmt.Println()
	// Output:
	// 0 1 2
}

func TestForBreakWithSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break
		}
	}
	// Output:
	// 0 1 2 3 4
}

func TestForBreakWithSwitch2(t *testing.T) {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loop
		}
	}
	// Output:
	// 0 1 2
}

func TestForBreakWithSwitch3(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			return
		}
	}
	// Output:
	// 0 1 2
}

func TestForBreakWithSelect(t *testing.T) {
	ch := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			ch <- struct{}{}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-ch:
				log.Println(">> ch")
			case <-ctx.Done():
				log.Println(">> ctx done")
				// bad: 只会退出select循环,不会退出外层的for循环
				break
			}
		}
	}()

	time.Sleep(time.Hour)
}

func TestForBreakWithSelect2(t *testing.T) {
	ch := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			ch <- struct{}{}
			time.Sleep(time.Second)
		}
	}()

	go func() {
	loop:
		for {
			select {
			case <-ch:
				log.Println(">> ch")
			case <-ctx.Done():
				log.Println(">> ctx done")
				// good: 退出外层的for循环
				break loop
			}
		}
	}()

	time.Sleep(time.Hour)
}

func TestForBreakWithSelect3(t *testing.T) {
	ch := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			ch <- struct{}{}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-ch:
				log.Println(">> ch")
			case <-ctx.Done():
				log.Println(">> ctx done")
				// good: 退出外层的for循环
				return
			}
		}
	}()

	time.Sleep(time.Hour)
}
