package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	call("http://localhost:28080/api/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/user/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/video/a/b/c/d/e/f/g?k1=v1&k2=v2")
}

func call(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
