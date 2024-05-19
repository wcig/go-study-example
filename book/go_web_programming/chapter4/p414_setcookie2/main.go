package main

import (
	"net/http"
)

// 设置cookie方式二: 通过http方法设置 (推荐)
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/set_cookie", setCookie)
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
