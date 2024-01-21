package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 修改 fetchall 程序，连续两次运行 fetchall

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchTwoTimes(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchTwoTimes(url string, ch chan string) {
	first := fetch(url)
	second := fetch(url)
	ch <- fmt.Sprintf("url: %s\nresp1: %s\nresp2: %s", url, first, second)
}

func fetch(url string) string {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("fetch: %v\n", err)
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("fetch: reading %s: %v\n", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs    %7d    %s", secs, nbytes, url)
}
