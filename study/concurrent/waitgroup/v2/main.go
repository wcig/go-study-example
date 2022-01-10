package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// v2：带错误返回的WaitGroup（使用golang.org/x/sync/errgroup包）
func main() {
	eg := &errgroup.Group{}
	urls := []string{
		"http://baidu.com",
		"http://bing.com",
		"http://google.com",
	}
	for _, url := range urls {
		tempUrl := url
		eg.Go(func() error {
			client := http.Client{Timeout: 3 * time.Second}
			resp, err := client.Get(tempUrl)
			if err == nil {
				_ = resp.Body.Close()
			}
			fmt.Println(tempUrl, err)
			return err
		})
	}
	err := eg.Wait()
	fmt.Println("over >>", err)

	// Output:
	// http://baidu.com <nil>
	// http://bing.com <nil>
	// http://google.com Get "http://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	// over >> Get "http://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
}
