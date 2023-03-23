package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(rewriteResponseMiddleware())
	e.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
	_ = e.Run(":28080")
}

func rewriteResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rb := newResponseBuffer(c.Writer)
		c.Writer = rb
		c.Next()
		copyData := rb.body.Bytes()
		fmt.Println(">> raw response data:", string(copyData))

		// reset body
		rb.body.Reset()
		// set response status
		rb.SetStatus(http.StatusInternalServerError)
		// set response header
		rb.SetHeader("Replace-Header", "true")
		// set response body
		_, err := rb.Write([]byte(`{"code": 100, "data": "replace data"}`))
		if err != nil {
			fmt.Println(">> response write err:", err)
			return
		}
		// flush body to ResponseWriter
		rb.Flush()
	}
}

type responseBuffer struct {
	gin.ResponseWriter
	body    *bytes.Buffer
	status  int
	flushed bool
}

func newResponseBuffer(rw gin.ResponseWriter) *responseBuffer {
	return &responseBuffer{
		ResponseWriter: rw,
		body:           &bytes.Buffer{},
		status:         rw.Status(),
	}
}

func (rb *responseBuffer) SetStatus(code int) {
	rb.status = code
}

func (rb *responseBuffer) SetHeader(key string, val string) {
	rb.ResponseWriter.Header().Set(key, val)
}

func (rb *responseBuffer) Write(b []byte) (int, error) {
	return rb.body.Write(b)
}

func (rb *responseBuffer) Flush() {
	if rb.flushed {
		return
	}
	rb.ResponseWriter.WriteHeader(rb.status)
	if rb.body.Len() > 0 {
		_, err := rb.ResponseWriter.Write(rb.body.Bytes())
		if err != nil {
			fmt.Println(">> response buffer flush err:", err)
			return
		}
		rb.body.Reset()
	}
	rb.flushed = true
}
