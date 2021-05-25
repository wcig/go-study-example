package ch14_concurrent_pattern

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 赢者为王

func init() {
	rand.Seed(time.Now().Unix())
}

func TestGetFirst(t *testing.T) {
	ch := make(chan string, 32)

	go func() {
		ch <- searchByBaidu("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBing("golang")
	}()

	fmt.Println(<-ch)
}

func searchByBaidu(keyword string) string {
	t := rand.Int63n(500)
	time.Sleep(time.Duration(t) * time.Millisecond)
	return "ok-baidu"
}

func searchByGoogle(keyword string) string {
	t := rand.Int63n(500)
	time.Sleep(time.Duration(t) * time.Millisecond)
	return "ok-google"
}

func searchByBing(keyword string) string {
	t := rand.Int63n(500)
	time.Sleep(time.Duration(t) * time.Millisecond)
	return "ok-bing"
}
