package main

import (
	"fmt"
	"net/http"
)

// 获取cookie方式一: 通过Request方法获取 (推荐)
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
	firstCookie, err := r.Cookie("first_cookie")
	if err != nil {
		panic(err)
	}
	fmt.Println("first_cookie:", firstCookie)

	cookies := r.Cookies()
	for i, v := range cookies {
		fmt.Println(i, v)
	}
	fmt.Fprintln(w, cookies)
}
