package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:28080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	cookie := &http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, cookie)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := &http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, rc)
		val, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, string(val))
	}
}
