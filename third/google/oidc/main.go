package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const (
	clientID = "497486293189-sing5lgvgp4nmdgk3gv1q5lfd4k04og6.apps.googleusercontent.com"
)

func main() {
	router := gin.Default()
	router.Use(referrerMiddleware)
	router.LoadHTMLGlob("templates/*")
	router.GET("/web/google/integration", integrationHandler)
	router.POST("/api/google/auth", authHandler)
	if err := router.Run(":28082"); err != nil {
		panic(err)
	}
}

func referrerMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	c.Next()
}

func integrationHandler(c *gin.Context) {
	obj := gin.H{"clientID": clientID}
	c.HTML(http.StatusOK, "index.html", obj)
}

type authReq struct {
	Credential string `form:"credential" binding:"required"`
	GCsrfToken string `form:"g_csrf_token"`
}

func authHandler(c *gin.Context) {
	var req authReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	jwtTokenStr := req.Credential
	payload, err := idtoken.Validate(context.Background(), jwtTokenStr, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	log.Printf("google api parse jwt token success, payload: %s", ToJsonStr(payload, true))

	segments := strings.Split(jwtTokenStr, ".")
	jwtHeaderBytes, _ := DecodeBase64URLWithJWT(segments[0])
	jwtPayloadBytes, _ := DecodeBase64URLWithJWT(segments[1])
	jwtHeader, jwtPayload := string(jwtHeaderBytes), string(jwtPayloadBytes)
	log.Printf("jwt token parse header:\n%s\npayload:\n%s", JsonStrPretty(jwtHeader), JsonStrPretty(jwtPayload))
	obj := gin.H{
		"jwtHeader":  jwtHeader,
		"jwtPayload": jwtPayload,
	}
	c.HTML(http.StatusOK, "welcome.html", obj)
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
		log.Print("ERR | json marshal err: %v", err)
	}
	return string(data)
}

func JsonStrPretty(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// 解码 Base64 URL 编码的字符串，补全可能缺少的 "=" 符号
func DecodeBase64URLWithJWT(encoded string) ([]byte, error) {
	// 补全缺少的 "=" 符号
	missingPadding := len(encoded) % 4
	if missingPadding > 0 {
		encoded += strings.Repeat("=", 4-missingPadding)
	}

	// Base64 URL 解码
	return base64.URLEncoding.DecodeString(encoded)
}

// Output:
// [GIN] 2024/05/23 - 21:26:50 | 200 |    1.693333ms |             ::1 | GET      "/web/google/integration"
// 2024/05/23 21:26:54 google api parse jwt token success, payload: {
//        "iss": "https://accounts.google.com",
//        "aud": "497486293189-sing5lgvgp4nmdgk3gv1q5lfd4k04og6.apps.googleusercontent.com",
//        "exp": 1716474413,
//        "iat": 1716470813,
//        "sub": "106151430417141188220"
// }
// 2024/05/23 21:26:54 jwt token parse header:
// {
//        "alg": "RS256",
//        "kid": "323b214ae6975a0f034ea77354dc0c25d03642dc",
//        "typ": "JWT"
// }
// payload:
// {
//        "iss": "https://accounts.google.com",
//        "azp": "497486293189-sing5lgvgp4nmdgk3gv1q5lfd4k04og6.apps.googleusercontent.com",
//        "aud": "497486293189-sing5lgvgp4nmdgk3gv1q5lfd4k04og6.apps.googleusercontent.com",
//        "sub": "106151430417141188220",
//        "email": "xxx@gmail.com",
//        "email_verified": true,
//        "nbf": 1716470513,
//        "name": "xxx",
//        "picture": "https://lh3.googleusercontent.com/a/ACg8ocLpbBbm7CWZn-J6Reedb-bI2wLnZD_gYZ083UxEdu7gv0__vg=s96-c",
//        "given_name": "xxx",
//        "family_name": "xxx",
//        "iat": 1716470813,
//        "exp": 1716474413,
//        "jti": "7a76a03975bce7a874d48e132d28c66481c3198a"
// }
// [GIN] 2024/05/23 - 21:26:54 | 200 |  1.188670959s |             ::1 | POST     "/api/google/auth"
