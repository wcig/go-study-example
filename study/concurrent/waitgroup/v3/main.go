package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// v3：带错误返回 + 错误一个直接返回的WaitGroup
// https://github.com/go-kratos/kratos/tree/v1.0.x/pkg/sync/errgroup
func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	urls := []string{
		"http://localhost:8081",
		"http://bing.com",
		"http://google.com",
	}
	for _, url := range urls {
		tempUrl := url
		eg.Go(func() error {
			done := make(chan error, 1)
			go func() {
				done <- request(tempUrl)
			}()
			select {
			case err := <-done:
				return err
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}
	err := eg.Wait()
	fmt.Println("over >>", err)

	// Output:
	// http://localhost:8081 Get "http://localhost:8081": dial tcp 127.0.0.1:8081: connect: connection refused
	// over >> Get "http://localhost:8081": dial tcp 127.0.0.1:8081: connect: connection refused
}

func request(tempUrl string) error {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(tempUrl)
	if err == nil {
		_ = resp.Body.Close()
	}
	fmt.Println(tempUrl, err)
	return err
}
