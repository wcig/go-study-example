package ch15_expvar

import (
	"expvar"
	"fmt"
	"net/http"
	"testing"
)

// expvar: 为公共变量提供标准化的接口，例如服务器中的操作计数器。

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "ok %s", r.URL.Path[1:])
}

// curl 'http://localhost:28080/live'
// 访问链接查看变量"visits": curl 'http://localhost:28080/debug/vars'
func Test(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":28080", nil)
}
