package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

// 编写自定义handler处理请求，任何请求都响应"hello world!"。
func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: &myHandler,
	}
	server.ListenAndServe()
}
