package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	call("http://localhost:28080/api/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/user/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/video/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/video/err?k1=v1&k2=v2")
}

func call(url string) {
	fmt.Println("--------------------------------")
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
	header, err := json.Marshal(resp.Header)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp statusCode: %d\nheader: %s\nbody:%s\n", resp.StatusCode, string(header), string(body))

	// Output:
	// --------------------------------
	// resp statusCode: 200
	// header: {"Content-Length":["46"],"Content-Type":["application/json; charset=utf-8"],"Date":["Sun, 08 Oct 2023 13:16:18 GMT"]}
	// body:{"code":0,"key":"api","path":"/a/b/c/d/e/f/g"}
	// --------------------------------
	// resp statusCode: 200
	// header: {"Content-Length":["26"],"Content-Type":["application/json; charset=utf-8"],"Date":["Sun, 08 Oct 2023 13:16:18 GMT"]}
	// body:{"code":0,"server":"user"}
	// --------------------------------
	// resp statusCode: 201
	// header: {"Content-Length":["38"],"Content-Type":["application/json; charset=utf-8"],"Date":["Sun, 08 Oct 2023 13:16:18 GMT"],"My-Custom-Header":["ok"]}
	// body:{"code":0,"msg":"ok","server":"video"}
	// --------------------------------
	// resp statusCode: 502
	// header: {"Content-Length":["30"],"Content-Type":["application/json; charset=utf-8"],"Date":["Sun, 08 Oct 2023 13:16:18 GMT"]}
	// body:{"code": 502, "result": false}
}
