package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
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

	var dayOfWeek []string
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		dayOfWeek = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	} else {
		dayOfWeek = []string{}
	}
	tmpl.Execute(w, dayOfWeek)
}
