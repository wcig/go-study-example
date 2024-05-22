package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	appURL         = "http://localhost:28081/web/shopify/integration/index.html"
	appRedirectURL = "http://localhost:28081/api/shopify/auth"
	appClientID    = "e94a778cd3d7e48b3f55f24ad315cb49"

	shopDomain    = "test20240319.myshopify.com"
	authorizePath = "/admin/oauth/authorize"
	scope         = "read_orders,read_products"
	state         = "202403190123"
)

var (
	appClientSecret = os.Getenv("SHOPIFY_APP_CLIENT_SECRET")
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/web/shopify/integration", integrationHandler)
	router.GET("/api/shopify/auth", authHandler)
	if err := router.Run(":28081"); err != nil {
		panic(err)
	}
}

func integrationHandler(c *gin.Context) {
	shopifyIntegrationUrl := fmt.Sprintf("https://%s%s?client_id=%s&redirect_uri=%s&scope=%s&state=%s&response_type=code", shopDomain, authorizePath, appClientID, appRedirectURL, scope, state)
	obj := gin.H{"shopifyIntegrationUrl": shopifyIntegrationUrl}
	c.HTML(http.StatusOK, "index.html", obj)
}

func authHandler(c *gin.Context) {
	// bind query params
	var req authReq
	if err := c.Bind(&req); err != nil {
		log.Printf("ERR | bind auth req err: %v", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	log.Printf("auth req: %+v", req)

	// get access token
	apiEndpoint := fmt.Sprintf("https://%s/admin/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		req.Shop, appClientID, appClientSecret, req.Code)
	var tokenInfo tokenRes
	resp, err := resty.New().R().SetResult(&tokenInfo).Post(apiEndpoint)
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Printf("ERR | request access token err: %v, %d", err, resp.StatusCode())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	log.Printf("access token: %+v", tokenInfo)

	obj := gin.H{"scope": tokenInfo.Scope, "accessToken": tokenInfo.AccessToken}
	c.HTML(http.StatusOK, "welcome.html", obj)
}

type authReq struct {
	Code      string `form:"code" json:"code"`
	Hmac      string `form:"hmac" json:"hmac"`
	Host      string `form:"host" json:"host"`
	Shop      string `form:"shop" json:"shop"`
	State     string `form:"state" json:"state"`
	Timestamp int64  `form:"timestamp" json:"timestamp"`
}

type tokenRes struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}
