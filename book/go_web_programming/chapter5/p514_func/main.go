package main

import (
	"html/template"
	"net/http"
	"time"
)

// 模板-函数
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	tmpl := template.New("tmpl.html").Funcs(funcMap)
	tmpl, err := tmpl.ParseFiles("tmpl.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, time.Now())
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
