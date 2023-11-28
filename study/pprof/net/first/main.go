package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 浏览器访问 http://localhost:6060/debug/pprof/ 查看详细信息
func main() {
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatalf("run pprof server err: %v", err)
	}
}
