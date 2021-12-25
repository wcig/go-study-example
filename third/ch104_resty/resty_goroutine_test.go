package ch104_resty

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

// memory leak test (recommend：re-use one resty client)

func TestNewClient1(t *testing.T) {
	start := time.Now()

	for i := 0; i < 100; i++ {
		client := resty.New()
		_, err := client.R().Get("https://baidu.com")
		fmt.Println(i, err, runtime.NumGoroutine())
		i++
	}

	fmt.Println(time.Since(start))
	time.Sleep(5 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
	// Output:
	// 0 <nil> 6
	// 1 <nil> 10
	// 2 <nil> 14
	// 3 <nil> 18
	// 4 <nil> 22
	// 5 <nil> 26
	// 6 <nil> 30
	// 7 <nil> 34
	// 8 <nil> 38
	// 9 <nil> 42
	// 4.10375054s
	// over >> 42
}

func TestOneClient1(t *testing.T) {
	client := resty.New()

	start := time.Now()
	for i := 0; i < 100; i++ {
		_, err := client.R().Get("https://baidu.com")
		fmt.Println(i, err, runtime.NumGoroutine())
		i++
	}

	fmt.Println(time.Since(start))
	time.Sleep(2 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
	// Output:
	// 0 <nil> 6
	// 1 <nil> 6
	// 2 <nil> 6
	// 3 <nil> 6
	// 4 <nil> 6
	// 5 <nil> 6
	// 6 <nil> 6
	// 7 <nil> 6
	// 8 <nil> 6
	// 9 <nil> 6
	// 1.920742976s
	// over >> 6
}

func TestNewClient2(t *testing.T) {
	start := time.Now()

	num := 100
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			client := resty.New()
			_, err := client.R().Get("https://baidu.com")
			fmt.Println(err, runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	time.Sleep(5 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
	// Output:
}

func TestOneClient2(t *testing.T) {
	start := time.Now()

	num := 100
	var wg sync.WaitGroup
	wg.Add(num)

	client := resty.New()
	for i := 0; i < num; i++ {
		go func() {
			_, err := client.R().Get("https://baidu.com")
			fmt.Println(err, runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	time.Sleep(5 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
	// Output:
}

func TestRawClient1(t *testing.T) {
	start := time.Now()

	for i := 0; i < 100; i++ {
		client := http.Client{}
		_, err := client.Get("https://baidu.com")
		fmt.Println(i, err, runtime.NumGoroutine())
		i++
	}

	fmt.Println(time.Since(start))
	time.Sleep(5 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
}

func TestRawClient2(t *testing.T) {
	start := time.Now()

	for i := 0; i < 100; i++ {
		_, err := http.Get("https://baidu.com")
		fmt.Println(i, err, runtime.NumGoroutine())
		i++
	}

	fmt.Println(time.Since(start))
	time.Sleep(5 * time.Second)
	fmt.Println("over >>", runtime.NumGoroutine())
}
