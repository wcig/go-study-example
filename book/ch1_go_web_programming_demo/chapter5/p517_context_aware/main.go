package main

import (
	"html/template"
	"net/http"
)

// 上下文感知
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}
