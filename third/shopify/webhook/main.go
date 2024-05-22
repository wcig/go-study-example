package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logFile       = flag.String("log_file", "webhook.log", "log file")
	shopifySecret = os.Getenv("SHOPIFY_APP_CLIENT_SECRET")

	logger *log.Logger

	shopifyWebhookHeaderKey = "shopifyWebhookHeader"
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
	app.Use(gin.LoggerWithWriter(io.Discard))
	app.Use(printMiddleware, webhookVerify)
	app.POST("/shopify/webhook", webhookHandler)
	if err := app.Run(":28080"); err != nil {
		log.Fatal(err)
	}
}

func webhookHandler(c *gin.Context) {
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

type ShopifyWebhookHeader struct {
	Topic       string `header:"X-Shopify-Topic" binding:"required"`
	HmacSha256  string `header:"X-Shopify-Hmac-Sha256" binding:"required"`
	ShopsDomain string `header:"X-Shopify-Shop-Domain" binding:"required"`
	APIVersion  string `header:"X-Shopify-API-Version"`
	WebhookID   string `header:"X-Shopify-Webhook-Id"`
	TriggeredAt string `header:"X-Shopify-Triggered-At"`
}

func webhookVerify(c *gin.Context) {
	// bind header
	var header ShopifyWebhookHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		logger.Printf("ERR | bind header err: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// verify
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Printf("ERR | read body err: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_ = c.Request.Body.Close()
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	err = CheckHmacSha256([]byte(shopifySecret), body, header.HmacSha256)
	if err != nil {
		logger.Printf("ERR | verify signature err: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set(shopifyWebhookHeaderKey, header)
	c.Next()
}

func CheckHmacSha256(secret, data []byte, expect string) error {
	hash := hmac.New(sha256.New, secret)
	hash.Write(data)
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if result != expect {
		return fmt.Errorf("check hmac sha256 failed, expect: %s, result: %s", expect, result)
	}
	return nil
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
