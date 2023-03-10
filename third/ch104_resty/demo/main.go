package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/get", getHandler)
	e.POST("/post", postHandler)
	e.GET("/err", errHandler)
	_ = e.Run(":38080")
}

func getHandler(c *gin.Context) {
	params := c.Request.URL.Query()
	headers := c.Request.Header
	result := map[string]interface{}{
		"code":    0,
		"params":  params,
		"headers": headers,
	}
	c.JSON(200, result)
}

func postHandler(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	headers := c.Request.Header
	result := map[string]interface{}{
		"code":    0,
		"params":  json.RawMessage(data),
		"headers": headers,
	}
	c.JSON(200, result)
}

func errHandler(c *gin.Context) {
	result := map[string]interface{}{
		"code": -1,
		"err":  "not found",
	}
	c.JSON(400, result)
}

func toJsonStr(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
