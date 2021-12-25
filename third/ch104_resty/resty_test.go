package ch104_resty

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// resty使用注意事项
// 1.复用同一个*resty.Client，不要每个请求创建一个*resty.Client，这样会造成内存问题并加长请求耗时。
// 2.使用同一个*resty.Client时，注意设置超时时间等参数会相互影响，同时也注意默认值。
// 3.底层包net/http.Client已实现连接池功能，使用resty不需要管连接池。

func TestFirst(t *testing.T) {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func TestDisableLogger(t *testing.T) {
	go timeoutServer()

	doRequest := func(client *resty.Client) {
		client.R().Get("http://localhost:28080")
	}

	// 默认logger
	fmt.Println("---")
	c1 := resty.New().SetTimeout(10 * time.Millisecond).SetRetryCount(1)
	doRequest(c1)

	// 方式一
	fmt.Println("---")
	l := logrus.New()
	l.Out = ioutil.Discard
	c2 := resty.New().SetTimeout(10 * time.Millisecond).SetRetryCount(1).SetLogger(l)
	doRequest(c2)

	// 方式二
	fmt.Println("---")
	c3 := resty.New().SetTimeout(10 * time.Millisecond).SetRetryCount(1).SetLogger(&restyLogger{})
	doRequest(c3)

	// Output:
	// ---
	// 2021/12/25 23:41:22.591072 ERROR RESTY Get "http://localhost:28080": context deadline exceeded (Client.Timeout exceeded while awaiting headers), Attempt 1
	// 2021/12/25 23:41:22.704099 ERROR RESTY Get "http://localhost:28080": context deadline exceeded (Client.Timeout exceeded while awaiting headers), Attempt 2
	// ---
	// ---
}

func timeoutServer() {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, "ok")
	}

	http.HandleFunc("/", handleFunc)
	_ = http.ListenAndServe(":28080", nil)
}

type restyLogger struct{}

func (rl *restyLogger) Errorf(format string, v ...interface{}) {}

func (rl *restyLogger) Warnf(format string, v ...interface{}) {}

func (rl *restyLogger) Debugf(format string, v ...interface{}) {}
