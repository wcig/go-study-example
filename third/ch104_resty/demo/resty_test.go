package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client  = resty.New()
	getUrl  = "http://localhost:38080/get"
	postUrl = "http://localhost:38080/post"
	errUrl  = "http://localhost:38080/err"
)

func TestFirst(t *testing.T) {
	fmt.Println(">> TestFirst")
	resp, err := client.R().EnableTrace().Get(getUrl)

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
}

func TestGet(t *testing.T) {
	fmt.Println(">> TestGet")
	headers := map[string]string{
		"uid":   "123",
		"uname": "tom",
	}
	params := map[string]string{
		"id":   "123",
		"name": "tom",
	}

	resp, err := client.R().SetHeaders(headers).SetQueryParams(params).Get(getUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp.String())

	type Result struct {
		Code    int                 `json:"code"`
		Params  map[string][]string `json:"params"`
		Headers map[string][]string `json:"headers"`
	}
	result := new(Result)
	resp2, err := client.R().SetHeaders(headers).SetQueryParams(params).SetResult(result).Get(getUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp2:", resp2.String())
	fmt.Println("result:", toJsonStr(result))
}

func TestPost(t *testing.T) {
	fmt.Println(">> TestPost")
	headers := map[string]string{
		"uid":   "123",
		"uname": "tom",
	}
	params := map[string]interface{}{
		"id":   123,
		"name": "tom",
	}

	resp, err := client.R().SetHeaders(headers).SetBody(params).Post(postUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp.String())

	type Result struct {
		Code    int                    `json:"code"`
		Params  map[string]interface{} `json:"params"`
		Headers map[string][]string    `json:"headers"`
	}
	result := new(Result)
	resp2, err := client.R().SetHeaders(headers).SetBody(params).SetResult(result).Post(postUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp2:", resp2.String())
	fmt.Println("result:", toJsonStr(result))
}

func TestErr(t *testing.T) {
	// 有响应,但http状态码400
	type Result struct {
		Code    int                 `json:"code"`
		Params  map[string][]string `json:"params"`
		Headers map[string][]string `json:"headers"`
		Err     string              `json:"err"`
	}
	result := new(Result)
	resp, err := client.R().SetError(result).Get(errUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp.StatusCode(), resp.String())
	fmt.Println("result:", toJsonStr(result))

	// 没有响应,连接失败
	resp2, err := client.R().Get("http://localhost:48080/err")
	if err != nil {
		fmt.Println("request err:", err.Error())
	}
	fmt.Println("resp2:", resp2.StatusCode(), resp2.String())
}

func TestCustomClient(t *testing.T) {
	// Create a Resty Client
	client2 := resty.New()

	// Retries are configured per client
	client2.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})

	_ = client2
}
