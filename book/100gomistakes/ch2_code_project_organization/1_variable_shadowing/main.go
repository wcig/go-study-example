package main

import (
	"log"
	"net/http"
)

// num1: Unintended variable shadowing: 没意识到的变量隐藏
func main() {
	// bad
	_ = example1()

	// good
	_ = example2()

	// Output:
	// 2023/11/20 21:52:05 >> example1 no trace: true
	// 2023/11/20 21:52:05 >> example1 over: false
	// 2023/11/20 21:52:05 >> example2 no trace: true
	// 2023/11/20 21:52:05 >> example2 over: true
}

var tracing bool

func example1() error {
	var client *http.Client
	if tracing {
		client, err := createClientWithTrace()
		if err != nil {
			return err
		}
		log.Println(">> example1 trace:", client != nil)
	} else {
		client, err := createClient()
		if err != nil {
			return err
		}
		log.Println(">> example1 no trace:", client != nil)
	}
	log.Println(">> example1 over:", client != nil)
	return nil
}

func example2() error {
	var (
		client *http.Client
		err    error
	)
	if tracing {
		client, err = createClientWithTrace()
		if err != nil {
			return err
		}
		log.Println(">> example2 trace:", client != nil)
	} else {
		client, err = createClient()
		if err != nil {
			return err
		}
		log.Println(">> example2 no trace:", client != nil)
	}
	log.Println(">> example2 over:", client != nil)
	return nil
}

func createClientWithTrace() (*http.Client, error) {
	return &http.Client{}, nil
}

func createClient() (*http.Client, error) {
	return &http.Client{}, nil
}
