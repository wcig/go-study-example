package main

import (
	"html/template"
	"net/http"
)

// go模板默认转义html标签来抵御xss攻击
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, r.FormValue("comment")) // <script>alert('Pwnd'!)</script>
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}
