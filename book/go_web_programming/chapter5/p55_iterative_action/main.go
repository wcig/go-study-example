package main

import (
	"html/template"
	"net/http"
)

// 迭代动作
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
	dayOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	tmpl.Execute(w, dayOfWeek)
}
