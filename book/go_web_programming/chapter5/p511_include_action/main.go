package main

import (
	"html/template"
	"net/http"
)

// 嵌套动作： 默认template.ParseFiles()的模板只对第一个html模板文件生效
func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("t1.html", "t2.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, "Hello World!")
}
