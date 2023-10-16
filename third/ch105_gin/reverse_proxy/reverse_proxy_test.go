package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	tc *http.Client
)

func initTC() {
	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.MaxConnsPerHost = 100
	tp.MaxIdleConns = 30
	tp.MaxIdleConnsPerHost = 30
	tc = &http.Client{
		Transport: tp,
	}
}

func Test(t *testing.T) {
	initTC()

	call("http://localhost:28080/api/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/user/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/video/a/b/c/d/e/f/g?k1=v1&k2=v2")
	call("http://localhost:28080/video/err?k1=v1&k2=v2")
}

func TestConcurrent(t *testing.T) {
	initTC()

	start := time.Now()
	const Num = 100
	wg := new(sync.WaitGroup)
	wg.Add(Num * 2)
	for i := 0; i < Num; i++ {
		go callWithWG("http://localhost:28080/user/a/b/c/d/e/f/g?k1=v1&k2=v2", wg)
		go callWithWG("http://localhost:28080/video/a/b/c/d/e/f/g?k1=v1&k2=v2", wg)
	}
	wg.Wait()
	fmt.Println(">> time cost:", time.Since(start))
}

func callWithWG(urlStr string, wg *sync.WaitGroup) {
	defer wg.Done()
	call(urlStr)
}

func call(url string) {
	// fmt.Println("--------------------------------")
	resp, err := tc.Get(url)
	if err != nil {
		fmt.Printf(">> call err: %v\n", err)
		return
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
	_ = body
	_ = header
	// fmt.Printf("resp statusCode: %d\nheader: %s\nbody:%s\n", resp.StatusCode, string(header), string(body))

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
