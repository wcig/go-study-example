package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// fetch 并发版本。

type Website struct {
	index int
	url   string
}

func main() {
	start := time.Now()
	urls := readAllURL()
	ch := make(chan string)
	for _, url := range urls {
		go fetchWithTimeout(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readAllURL() []string {
	file, err := os.Open("top100.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var sites []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		sites = append(sites, line)
	}
	return sites
}

func fetchWithTimeout(url string, ch chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	done := make(chan string, 1)
	go func() {
		done <- fetch(url)
	}()

	select {
	case result := <-done:
		ch <- result
	case <-ctx.Done():
		ch <- fmt.Sprintf("fetch: %s timeout", url)
	}
}

func fetch(url string) string {
	start := time.Now()
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("fetch: %v", err)
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("fetch: reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs    %7d    %s", secs, nbytes, url)
}
