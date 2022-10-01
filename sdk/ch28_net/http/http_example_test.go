package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestClientGet(t *testing.T) {
	resp, err := http.Get("https://baidu.com")
	assert.Nil(t, err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	fmt.Println("response:", string(data))
}

func timeoutServer() {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprint(w, "ok")
	}

	http.HandleFunc("/", handleFunc)
	_ = http.ListenAndServe(":28081", nil)
}

func TestClientSettingWithTimeout1(t *testing.T) {
	go timeoutServer()

	client := http.Client{Timeout: time.Second}
	fmt.Println(time.Now())

	_, err := client.Get("http://localhost:28081")
	fmt.Println(time.Now())
	fmt.Println("timeout err:", err)

	time.Sleep(time.Minute)
}

func TestClientSettingWithTimeout2(t *testing.T) {
	go timeoutServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", "http://localhost:28081", nil)
	assert.Nil(t, err)
	fmt.Println(time.Now())

	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	fmt.Println(time.Now())
	fmt.Println("timeout err:", err)

	time.Sleep(time.Minute)
}

func TestClientSettingWithTimeout3(t *testing.T) {
	go timeoutServer()

	requestFunc := func() error {
		_, err := http.Get("http://localhost:28081")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(time.Now())

	done := make(chan error)
	go func() {
		done <- requestFunc()
	}()

	select {
	case <-done:
		fmt.Println(time.Now())
		fmt.Println("done err:", done)
	case <-ctx.Done():
		fmt.Println(time.Now())
		fmt.Println("timeout err:", ctx.Err())
	}

	time.Sleep(time.Minute)
}

// 注意:
// 1.go的http.client自带连接池功能,默认连接数没有限制,可以通过设置transport生效.
// 2.resp.Body记得关闭资源, 并且只有读取了resp.Body数据才生效.
func TestClientSettingWithTransport(t *testing.T) {
	// 方式一
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 10
	transport.MaxIdleConns = 10
	transport.MaxIdleConnsPerHost = 10

	// 方式二
	// transport := &http.Transport{
	// 	MaxConnsPerHost:     10,
	// 	MaxIdleConns:        10,
	// 	MaxIdleConnsPerHost: 10,
	// }
	client := http.Client{
		Transport: transport,
	}

	resp, err := client.Get("https://baidu.com")
	assert.Nil(t, err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	fmt.Println("response:", string(data))
}
