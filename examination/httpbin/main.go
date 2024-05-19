package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", handler)
	_ = r.Run(":28001")
}

func handler(c *gin.Context) {
	// args
	args := make(map[string]string)
	if err := c.BindQuery(&args); err != nil {
		fmt.Println(err)
	}

	// headers
	headers := make(map[string]string)
	rawHeaders := c.Request.Header
	for k, v := range rawHeaders {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// origin
	origin := c.ClientIP()

	// url
	url := getRawRequestUrl(c)

	// response
	info := &httpBinInfo{
		Args:    args,
		Headers: headers,
		Origin:  origin,
		Url:     url,
	}
	c.JSON(http.StatusOK, info)
}

func getRawRequestUrl(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host + c.Request.RequestURI
}

type httpBinInfo struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}
