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
	if err := r.ParseMultipartForm(1024); err != nil {
		panic(err)
	}
	fileHeaders, ok := r.MultipartForm.File["uploaded"]
	if !ok || len(fileHeaders) == 0 {
		panic("file empty")
	}
	fileHeader := fileHeaders[0]
	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(bytes))
}
