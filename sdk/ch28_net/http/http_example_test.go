package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCanonicalHeaderKey(t *testing.T) {
	fmt.Println(http.CanonicalHeaderKey("accept-encoding"))
	fmt.Println(http.CanonicalHeaderKey("content-type"))
	// output:
	// Accept-Encoding
	// Content-Type
}

func TestDetectContentType(t *testing.T) {
	fmt.Println(http.DetectContentType([]byte("ok"))) // text/plain; charset=utf-8
}

func TestStatusText(t *testing.T) {
	fmt.Println(http.StatusText(http.StatusOK))
	fmt.Println(http.StatusText(http.StatusNotFound))
	// output:
	// OK
	// Not Found
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func MyHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func TestHandler(t *testing.T) {
	handler := &MyHandler{}
	http.Handle("/", handler)
	err := http.ListenAndServe(":28080", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandleFunc(t *testing.T) {
	http.HandleFunc("/", MyHandleFunc)
	err := http.ListenAndServe(":28080", nil)
	if err != nil {
		t.Fatal(err)
	}
}
