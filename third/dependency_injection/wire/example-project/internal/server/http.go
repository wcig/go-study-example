package server

import (
	"go-app/third/dependency_injection/wire/example-project/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHTTPServer)

func NewHTTPServer(c *config.Server) *http.Server {
	gin.SetMode(c.Mode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	server := &http.Server{
		Addr:    c.Address,
		Handler: engine,
	}
	return server
}
