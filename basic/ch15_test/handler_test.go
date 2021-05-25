package ch15

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	Routes()
}

func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("create request failed, err", err)
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	t.Log("code:", rw.Code)
	t.Log("body:", rw.Body.String())
}
