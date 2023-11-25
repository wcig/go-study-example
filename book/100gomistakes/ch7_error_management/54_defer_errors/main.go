package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

// bad
func example1() {
	// 创建一个自定义的http.Client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 禁用自动重定向
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("https://baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

// bad
func example2() {
	// 创建一个自定义的http.Client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 禁用自动重定向
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("https://baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

// good
func example3() {
	// 创建一个自定义的http.Client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 禁用自动重定向
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("https://baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("response body close err: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

// good
func example4() {
	// 创建一个自定义的http.Client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 禁用自动重定向
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("https://baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer closeResource(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func closeResource(r io.Closer) {
	if err := r.Close(); err != nil {
		log.Printf("close resource err: %v", err)
	}
}
