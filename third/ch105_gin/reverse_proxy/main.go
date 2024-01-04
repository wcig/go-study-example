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
	proxyHub  map[string]*url.URL
	transport *http.Transport
)

func main() {
	initProxyResource()
	e := gin.Default()
	e.Any("/:server/*path", reverseProxyHandler)
	_ = e.Run(":28080")
}

func initProxyResource() {
	hostMap := map[string]string{
		"user":  "http://localhost:28081",
		"video": "http://localhost:28082",
	}
	proxyHub = make(map[string]*url.URL)
	for key, host := range hostMap {
		remote, err := url.Parse(host)
		if err != nil {
			panic(err)
		}
		proxyHub[key] = remote
	}

	transport = http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConns = 30
	transport.MaxIdleConnsPerHost = 30
}

func reverseProxyHandler(c *gin.Context) {
	key := c.Param("server")
	path := c.Param("path")
	if key == "" {
		c.JSON(404, map[string]interface{}{"error": "not found"})
		return
	}

	if u, ok := proxyHub[key]; ok {
		py := newReverseProxy(u)
		py.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = u.Host
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
			req.URL.Path = path
		}
		py.ServeHTTP(c.Writer, c.Request)
		return
	}

	c.JSON(200, map[string]interface{}{"code": 0, "key": key, "path": path})
	return
}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	p := httputil.NewSingleHostReverseProxy(target)
	p.Transport = transport
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
