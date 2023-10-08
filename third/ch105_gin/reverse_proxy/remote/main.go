package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	go runUserServer()
	go runVideoServer()

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func runUserServer() {
	e := gin.Default()
	e.Any("/*path", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"code": 0, "server": "user"})
	})
	_ = e.Run(":28081")
}

func runVideoServer() {
	e := gin.Default()
	e.Any("/*path", func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/err") {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("test bad request"))
		} else {
			c.JSON(200, map[string]interface{}{"code": 0, "server": "video"})
		}
	})
	_ = e.Run(":28082")
}
