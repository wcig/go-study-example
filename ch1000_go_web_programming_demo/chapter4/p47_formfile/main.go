package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(bytes))
}
