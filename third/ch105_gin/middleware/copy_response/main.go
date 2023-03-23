// refer: https://github.com/gin-gonic/gin/issues/3384
package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(copyResponseMiddleware())
	e.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
	_ = e.Run(":28080")
}

type responseBuffer struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func newResponseBuffer(rw gin.ResponseWriter) *responseBuffer {
	return &responseBuffer{
		ResponseWriter: rw,
		buf:            &bytes.Buffer{},
	}
}

func (rb *responseBuffer) Write(b []byte) (int, error) {
	return rb.buf.Write(b)
}

func copyResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rb := newResponseBuffer(c.Writer)
		c.Writer = rb
		c.Next()
		copyData := rb.buf.Bytes()
		fmt.Println(">> copy response data:", string(copyData))
	}
}
