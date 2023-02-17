package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	proxyMap map[string]*Proxy
)

type Proxy struct {
	host  *url.URL
	proxy *httputil.ReverseProxy
}

func init() {
	hostMap := map[string]string{
		"user":  "http://localhost:28081",
		"video": "http://localhost:28082",
	}
	proxyMap = make(map[string]*Proxy)
	for key, host := range hostMap {
		remote, err := url.Parse(host)
		if err != nil {
			panic(err)
		}
		proxy := &Proxy{
			host:  remote,
			proxy: httputil.NewSingleHostReverseProxy(remote),
		}
		proxyMap[key] = proxy
	}
}

func main() {
	e := gin.Default()
	e.Any("/:server/*path", reverseProxyHandler)
	_ = e.Run(":28080")
}

func reverseProxyHandler(c *gin.Context) {
	key := c.Param("server")
	path := c.Param("path")
	if key == "" {
		c.JSON(404, map[string]interface{}{"error": "not found"})
		return
	}

	if proxy, ok := proxyMap[key]; ok {
		proxy.proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = proxy.host.Host
			req.URL.Scheme = proxy.host.Scheme
			req.URL.Host = proxy.host.Host
			req.URL.Path = path
		}
		proxy.proxy.ServeHTTP(c.Writer, c.Request)
		return
	}

	c.JSON(200, map[string]interface{}{"code": 0, "key": key, "path": path})
	return
}
