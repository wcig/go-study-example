package main

import (
	"fmt"
	"net/http"
)

// 获取cookie方式一: 通过请求头获取 (不推荐)
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "first_cookie",
		Value:    "tom",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "second_cookie",
		Value:    "jerry",
		HttpOnly: true,
	}
	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	vals, ok := r.Header["Cookie"]
	if !ok {
		fmt.Fprintln(w, "no cookie")
		return
	}
	for i, v := range vals {
		fmt.Println(i, v)
	}
	fmt.Fprintln(w, vals)
}
