package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// http2
func main() {
	handler := HelloHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wolrd.\n")
}
