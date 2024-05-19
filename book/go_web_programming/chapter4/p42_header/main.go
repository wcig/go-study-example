package main

import (
	"fmt"
	"net/http"
)

// 获取请求头 (header结构: type Header map[string][]string)
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/header", headers)
	server.ListenAndServe()
}

func headers(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k, v := range header {
		fmt.Printf("key:%s, val:%v\n", k, v)
	}
	fmt.Fprintln(w, header)
}
