package main

import (
	"html/template"
	"net/http"
)

// 设置动作
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tmpl.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, "hello")
}
