package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 第三方多路复用器httprouter使用
func main() {
	router := httprouter.New()
	router.GET("/hello/:name", hello)
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: router,
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
