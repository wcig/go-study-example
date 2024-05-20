package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logFile = flag.String("log_file", "", "log file")

	logger *log.Logger
)

func main() {
	initLogger()
	runServer()
}

func initLogger() {
	flag.Parse()
	var w io.Writer
	if logFile != nil && *logFile != "" {
		file, err := os.OpenFile(*logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		w = io.MultiWriter(file, os.Stdout)
	} else {
		w = os.Stdout
	}
	logger = log.New(w, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func runServer() {
	app := gin.Default()
	app.Use(printMiddleware, gin.LoggerWithWriter(io.Discard))
	app.Any("/*path", handler)
	if err := app.Run(":28080"); err != nil {
		log.Fatal(err)
	}
}

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"code": 0})
}

func printMiddleware(c *gin.Context) {
	// summary
	method := c.Request.Method
	proto := c.Request.Proto
	host := c.Request.Host
	path := c.Request.URL.Path
	clientIp := c.ClientIP()
	logger.Printf(">> summary: %s | %s | %s | %s | %s", method, proto, host, path, clientIp)

	// header
	headers := c.Request.Header
	logger.Printf(">> header: %s", ToJsonStr(headers, true))

	// body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Printf("ERR | read body err: %v", err)
	}
	_ = c.Request.Body.Close()
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	logger.Printf(">> body: %s", body)
}

func ToJsonStr(v interface{}, p bool) string {
	var (
		data []byte
		err  error
	)
	if p {
		data, err = json.MarshalIndent(v, "", "\t")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		logger.Print("ERR | json marshal err: %v", err)
	}
	return string(data)
}
