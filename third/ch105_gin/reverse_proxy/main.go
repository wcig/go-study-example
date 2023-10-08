package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
			proxy: newReverseProxy(remote),
		}
		proxyMap[key] = proxy
	}
}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	p := httputil.NewSingleHostReverseProxy(target)
	p.ModifyResponse = func(r *http.Response) error {
		if strings.Contains(r.Request.RequestURI, "/user") {
			return nil
		}

		// read raw response
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf(">> reverse proxy modify response, read err: %v", err)
			return err
		}
		defer r.Body.Close()
		log.Printf(">> reverse proxy modify response, body: %s", string(body))

		// modify response body
		m := make(map[string]interface{})
		if err = json.Unmarshal(body, &m); err != nil {
			log.Printf(">> reverse proxy modify response, unmarshal err: %v", err)
			return err
		}
		m["msg"] = "ok"
		newBody, err := json.Marshal(m)
		if err != nil {
			log.Printf(">> reverse proxy modify response, marshal err: %v", err)
			return err
		}
		r.Body = io.NopCloser(bytes.NewReader(newBody))

		// modify response header
		r.Header["Content-Length"] = []string{fmt.Sprint(len(newBody))}
		r.Header["My-Custom-Header"] = []string{"ok"}

		// modify response status code
		r.StatusCode = http.StatusCreated

		return nil
	}
	p.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
		log.Printf(">> reverse proxy error handler, err: %v", err)
		rw.WriteHeader(http.StatusBadGateway)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		badGatewayResponse := []byte(`{"code": 502, "result": false}`)
		if _, err = rw.Write(badGatewayResponse); err != nil {
			log.Printf(">> reverse proxy error handler, write err: %v", err)
		}
	}
	return p
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
