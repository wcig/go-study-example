package main

import "net/http"

// https服务：cert.pem为ssl证书，key.pem为证书对应私钥
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
