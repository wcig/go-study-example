package main

import (
	"fmt"
	"net/http"
)

// post x-www-form-urlencoded表单请求解析
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for k, v := range r.PostForm {
		fmt.Printf("postForm key:%s, val:%v\n", k, v)
	}
	for k, v := range r.Form {
		fmt.Printf("Form key:%s, val:%v\n", k, v)
	}
	fmt.Fprintln(w, r.Form)
}
