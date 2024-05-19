package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// 串联两个处理器函数
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
