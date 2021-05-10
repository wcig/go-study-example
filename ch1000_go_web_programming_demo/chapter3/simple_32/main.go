package main

import "net/http"

// 简单示例：可自定义配置（只监听端口28080不做任何事情）
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: nil,
	}
	server.ListenAndServe()
}
