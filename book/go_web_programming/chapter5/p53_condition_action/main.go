package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// 条件动作
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
	rand.Seed(time.Now().Unix())
	tmpl.Execute(w, rand.Intn(10) > 5)
}
