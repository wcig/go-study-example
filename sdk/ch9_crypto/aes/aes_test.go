package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wumansgy/goEncrypt"
)

// aes: 一种对称加密算法

// 常量
// const BlockSize = 16 // AES块大小 (byte为单位)

// 函数
// func NewCipher(key []byte) (cipher.Block, error) // 基于给定key创建一新的cipher.Block (key16,32,64字节对应AES-128,AES-192,AES-256)

const (
	iv = "1234567890123456"
)

func TestAesCbcEncrypt(t *testing.T) {
	data := []byte("123456")
	key := []byte("1234567890123456")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	paddingData := PKCS5Padding(data, block.BlockSize())
	result := make([]byte, len(paddingData))
	blockMode.CryptBlocks(result, paddingData)
	fmt.Println(result, len(result)) // [214 55 115 90 233 226 27 165 12 182 134 183 79 171 141 44] 16

	except, err := goEncrypt.AesCbcEncrypt(data, key, []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, except, result)
}

func TestAesCbcDecrypt(t *testing.T) {
	encryptData := []byte{214, 55, 115, 90, 233, 226, 27, 165, 12, 182, 134, 183, 79, 171, 141, 44}
	key := []byte("1234567890123456")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	paddingResult := make([]byte, len(encryptData))
	blockMode.CryptBlocks(paddingResult, encryptData)
	result, err := PKCS5UnPadding(paddingResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(result, len(result), string(result)) // [49 50 51 52 53 54] 6 123456

	except, err := goEncrypt.AesCbcDecrypt(encryptData, key, []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, except, result)
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - (len(src) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(src, padText...)
	return newText
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	number := int(src[length-1])
	if number > length {
		return nil, errors.New("padding size error")
	}
	return src[:length-number], nil
}
