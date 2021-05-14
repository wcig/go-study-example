package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/write", write)
	http.HandleFunc("/writeheader", writeHeader)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

func write(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "No such service, try next door")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://baidu.com")
	w.WriteHeader(http.StatusFound)
}

type Post struct {
	User    string   `json:"user"`
	Threads []string `json:"threads"`
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "tom",
		Threads: []string{"first", "second", "third"},
	}
	bytes, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
