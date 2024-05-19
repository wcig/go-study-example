package main

import (
	"fmt"
	"net/http"
)

// 串联多个处理器
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log..")
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protect..")
		h.ServeHTTP(w, r)
	})
}
