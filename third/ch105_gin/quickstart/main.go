package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	RegisterHandler(app)
	_ = app.Run(":28080")
}

func RegisterHandler(app *gin.Engine) {
	app.GET("/print", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		log.Println(">> name:", name)
		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}

// Output:
// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// - using env:   export GIN_MODE=release
// - using code:  gin.SetMode(gin.ReleaseMode)
//
// [GIN-debug] GET    /print                    --> main.RegisterHandler.func1 (3 handlers)
// [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
// Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
// [GIN-debug] Listening and serving HTTP on :28080
// 2023/11/14 22:46:25 >> name: tom
// [GIN] 2023/11/14 - 22:46:25 | 200 |     333.792µs |       127.0.0.1 | GET      "/print?name=tom"

// Request:
// $ curl -i http://localhost:28080/print?name=tom
// HTTP/1.1 200 OK
// Content-Type: application/json; charset=utf-8
// Date: Tue, 14 Nov 2023 14:46:25 GMT
// Content-Length: 10
//
// {"code":0}
