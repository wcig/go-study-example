package main

import (
	"fmt"
	"io"
	"net/http"
)

// 获取请求体
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	num, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println("read byte size:", num)
	fmt.Fprintln(w, string(body))
}
