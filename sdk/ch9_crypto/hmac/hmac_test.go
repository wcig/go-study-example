package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 函数
// func Equal(mac1, mac2 []byte) bool // 比较2个mac值是否相等,而不泄露时序信息
// func New(h func() hash.Hash, key []byte) hash.Hash // 给定一个hash.Hash类型和一个字节切片key,返回hmac的hash.Hash

func TestHmacMD5(t *testing.T) {
	key := []byte("abc123")
	mac := hmac.New(md5.New, key)

	data := []byte("123456")
	mac.Write(data)
	sum := mac.Sum(nil)
	val := hex.EncodeToString(sum)
	fmt.Println(val, len(val)) // 925c88e6c9088e684be78b5c9e47e17f 32
}

func TestHmacSha256(t *testing.T) {
	key := []byte("abc123")
	mac := hmac.New(sha256.New, key)

	data := []byte("123456")
	mac.Write(data)
	sum := mac.Sum(nil)
	val := hex.EncodeToString(sum)
	fmt.Println(val, len(val)) // 8232e879c618696029c06fb07c7ca3678c0078ab334846afff476d275b25c62d 64
}

func TestHmacEqual(t *testing.T) {
	key := []byte("abc123")
	message := []byte("123456")
	messageMAC, _ := hex.DecodeString("8232e879c618696029c06fb07c7ca3678c0078ab334846afff476d275b25c62d")
	assert.True(t, ValidSha256MAC(message, messageMAC, key))
	messageMAC[0] = byte('0')
	assert.False(t, ValidSha256MAC(message, messageMAC, key))

}

func ValidSha256MAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
