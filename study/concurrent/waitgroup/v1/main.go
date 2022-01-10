package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// v1：不带错误返回的WaitGroup（使用sync包）
func main() {
	wg := &sync.WaitGroup{}
	urls := []string{
		"http://baidu.com",
		"http://bing.com",
		"http://google.com",
	}
	for _, url := range urls {
		wg.Add(1)
		tempUrl := url

		go func() {
			defer wg.Done()

			client := http.Client{Timeout: 3 * time.Second}
			resp, err := client.Get(tempUrl)
			if err == nil {
				_ = resp.Body.Close()
			}
			fmt.Println(tempUrl, err)
		}()
	}
	wg.Wait()
	fmt.Println("over >>")

	// Output:
	// http://baidu.com <nil>
	// http://bing.com <nil>
	// http://google.com Get "http://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	// over >>
}
