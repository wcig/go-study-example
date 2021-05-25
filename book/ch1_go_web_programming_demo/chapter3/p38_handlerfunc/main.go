package main

import (
	"fmt"
	"net/http"
)

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func worldHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

// 定义handler函数处理请求
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/hello", helloHandlerFunc)
	http.HandleFunc("/world", worldHandlerFunc)
	server.ListenAndServe()
}
