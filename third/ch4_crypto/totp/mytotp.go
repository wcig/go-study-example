package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

/**
 * TOTP简单实现, 对于边界异常情况没有处理
 * 1.使用算法: hmac1
 * 2.有效时间: 30s
 * 3.密钥要求: 16/32..位长度
 */

func GenerateCode(secret string) (code string, err error) {
	counter := time.Now().Unix() / int64(30)
	return generateCodeWithCounter(secret, uint64(counter))
}

func generateCodeWithCounter(secret string, counter uint64) (code string, err error) {
	// secret -> key
	secret = strings.TrimSpace(secret)
	if n := len(secret) % 8; n != 0 {
		secret += strings.Repeat("=", 8-n)
	}
	secret = strings.ToUpper(secret)
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// hmac1加密
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(counter))
	mac := hmac.New(sha1.New, secretBytes)
	mac.Write(buf)
	sum := mac.Sum(nil) // 长度20的字节切片

	// 加密结果 -> 6位数字密码
	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))
	mod := math.Pow10(6)
	value = value % int64(mod)
	code = fmt.Sprintf("%06d", value)
	return code, nil
}

func ValidateCode(secret string, code string) bool {
	counter := uint64(time.Now().Unix() / int64(30))
	counters := []uint64{counter - 1, counter, counter + 1}
	for _, item := range counters {
		val, err := generateCodeWithCounter(secret, item)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if val == code {
			return true
		}
	}
	return false
}
