package ch203_jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

const (
	verifyKey = "123456" // 密钥
	ttl       = 3600     // 失效时间1小时(单位秒)
)

func TestBuildNoClaimsToken(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenStr, err := token.SignedString([]byte(verifyKey))
	assert.Nil(t, err)
	fmt.Println(tokenStr)
}

func TestParseNoClaimsToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.B_pYxir_XkZ7obWeYgsaV7mqUGKm4OMf6yCu0Ve64tU"
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	assert.Nil(t, err)
	b, _ := json.MarshalIndent(token, "", "\t")
	fmt.Println(string(b))
}

// output:
// {
//	"Raw": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.B_pYxir_XkZ7obWeYgsaV7mqUGKm4OMf6yCu0Ve64tU",
//	"Method": {
//		"Name": "HS256",
//		"Hash": 5
//	},
//	"Header": {
//		"alg": "HS256",
//		"typ": "JWT"
//	},
//	"Claims": {},
//	"Signature": "B_pYxir_XkZ7obWeYgsaV7mqUGKm4OMf6yCu0Ve64tU",
//	"Valid": true
// }

func TestBuildJwtToken(t *testing.T) {
	now := time.Now().Unix()
	claims := jwt.StandardClaims{
		Audience:  "myapp",   // aud (audience)：受众
		ExpiresAt: now + ttl, // exp (expiration time)：过期时间
		Id:        "001",     // jti (JWT ID)：编号
		IssuedAt:  now,       // iat (Issued At)：签发时间
		Issuer:    "wcig",    // iss (issuer)：签发人
		NotBefore: now,       // nbf (Not Before)：生效时间
		Subject:   "user",    // sub (subject)：主题
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(verifyKey))
	assert.Nil(t, err)
	fmt.Println(tokenStr)
}

func TestParseJwtToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJteWFwcCIsImV4cCI6MTYxOTQ0MDk2NCwianRpIjoiMDAxIiwiaWF0IjoxNjE5NDM3MzY0LCJpc3MiOiJ3Y2lnIiwibmJmIjoxNjE5NDM3MzY0LCJzdWIiOiJ1c2VyIn0.IgxK9hhE-UsHUr12ImNPI_EgUnxMFU1QpzAWV3zE78U"
	var claims jwt.StandardClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	assert.Nil(t, err)
	PrettyPrintJson(token)
	PrettyPrintJson(claims)

	myClaims, ok := token.Claims.(*jwt.StandardClaims)
	if ok && token.Valid {
		PrettyPrintJson(myClaims)
	}
}

// output:
// {
//	"Raw": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJteWFwcCIsImV4cCI6MTYxOTQ0MDk2NCwianRpIjoiMDAxIiwiaWF0IjoxNjE5NDM3MzY0LCJpc3MiOiJ3Y2lnIiwibmJmIjoxNjE5NDM3MzY0LCJzdWIiOiJ1c2VyIn0.IgxK9hhE-UsHUr12ImNPI_EgUnxMFU1QpzAWV3zE78U",
//	"Method": {
//		"Name": "HS256",
//		"Hash": 5
//	},
//	"Header": {
//		"alg": "HS256",
//		"typ": "JWT"
//	},
//	"Claims": {
//		"aud": "myapp",
//		"exp": 1619440964,
//		"jti": "001",
//		"iat": 1619437364,
//		"iss": "wcig",
//		"nbf": 1619437364,
//		"sub": "user"
//	},
//	"Signature": "IgxK9hhE-UsHUr12ImNPI_EgUnxMFU1QpzAWV3zE78U",
//	"Valid": true
// }
// {
//	"aud": "myapp",
//	"exp": 1619440964,
//	"jti": "001",
//	"iat": 1619437364,
//	"iss": "wcig",
//	"nbf": 1619437364,
//	"sub": "user"
// }
// {
//	"aud": "myapp",
//	"exp": 1619440964,
//	"jti": "001",
//	"iat": 1619437364,
//	"iss": "wcig",
//	"nbf": 1619437364,
//	"sub": "user"
// }

func PrettyPrintJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

// JWT header/payload base64url 编码算法，会默认删除末尾的 “=” 符号
func TestJWTBase64URL(t *testing.T) {
	rawHeader := `{"alg":"HS256","typ":"JWT"}`
	rawPayload := `{"sub":"1234567890","name":"John Doe","iat":1516239022}`

	header := EncodeBase64URLWithJWT([]byte(rawHeader))
	payload := EncodeBase64URLWithJWT([]byte(rawPayload))
	fmt.Println("base64url encode header:", header)
	fmt.Println("base64url encode payload:", payload)

	decodeHeader, _ := DecodeBase64URLWithJWT(header)
	decodePayload, _ := DecodeBase64URLWithJWT(payload)
	assert.Equal(t, rawHeader, string(decodeHeader))
	assert.Equal(t, rawPayload, string(decodePayload))
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

// 编码字符串为 Base64 URL 编码格式，并去除末尾可能多余的 "=" 符号
func EncodeBase64URLWithJWT(data []byte) string {
	// Base64 URL 编码
	encoded := base64.URLEncoding.EncodeToString(data)

	// 去除末尾可能多余的 "=" 符号
	encoded = strings.TrimRight(encoded, "=")

	return encoded
}
