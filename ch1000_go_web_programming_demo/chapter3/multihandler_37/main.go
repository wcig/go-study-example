package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

// 定义多个handler实例处理请求：不同请求路径做不同处理。
func main() {
	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.Handle("/hello", &helloHandler)
	http.Handle("/world", &worldHandler)
	server.ListenAndServe()
}
