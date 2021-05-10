package main

import "net/http"

// 最简单示例：只监听端口28080不做任何事情
func main() {
	http.ListenAndServe("127.0.0.1:28080", nil)
}
