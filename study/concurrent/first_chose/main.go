package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 赢者为王
func main() {
	ch := make(chan string, 10)

	go func() {
		ch <- searchBing("golang")
	}()
	go func() {
		ch <- searchBaidu("golang")
	}()
	go func() {
		ch <- searchGoogle("golang")
	}()

	fmt.Println("result:", <-ch)
	// Output:
	// google: 1580
	// baidu: 1372
	// bing: 1429
	// result: baidu-golang
}

func searchBing(keyword string) string {
	rand.Seed(time.Now().UnixNano())
	duration := 1000 + rand.Intn(1000)
	fmt.Println("bing:", duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	return "bing-" + keyword
}

func searchBaidu(keyword string) string {
	rand.Seed(time.Now().UnixNano())
	duration := 1000 + rand.Intn(1000)
	fmt.Println("baidu:", duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	return "baidu-" + keyword
}

func searchGoogle(keyword string) string {
	rand.Seed(time.Now().UnixNano())
	duration := 1000 + rand.Intn(1000)
	fmt.Println("google:", duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	return "google-" + keyword
}
