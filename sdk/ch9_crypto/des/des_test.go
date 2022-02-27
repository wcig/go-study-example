package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wumansgy/goEncrypt"
)

// des: 一种对称加密算法 (给定明文数据+秘钥+初始向量iv,输出加密密文)

// 常量
// const BlockSize = 8 // DES块大小(以byte为单位)

// 函数
// func NewCipher(key []byte) (cipher.Block, error)          // 返回DES的cipher.Block
// func NewTripleDESCipher(key []byte) (cipher.Block, error) // 返回3DES的cipher.Block

const (
	iv = "12345678"
)

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

func TestDesCbcEncrypt(t *testing.T) {
	data := []byte("123456")
	key := []byte("12345678")

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	paddingData := PKCS5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	result := make([]byte, len(paddingData))
	blockMode.CryptBlocks(result, paddingData)
	fmt.Println(result, len(result)) // [29 69 254 237 91 71 129 189] 8

	expect, err := goEncrypt.DesCbcEncrypt(data, key, key)
	assert.Nil(t, err)
	assert.Equal(t, expect, result)
}

func TestDesCbcDecrypt(t *testing.T) {
	encryptData := []byte{29, 69, 254, 237, 91, 71, 129, 189}
	key := []byte("12345678")

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	paddingResult := make([]byte, len(encryptData))
	blockMode.CryptBlocks(paddingResult, encryptData)
	result, err := PKCS5UnPadding(paddingResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(result, len(result))

	except, err := goEncrypt.DesCbcDecrypt(encryptData, key, key)
	assert.Nil(t, err)
	assert.Equal(t, except, result)
}

func Test3DesCbcEncrypt(t *testing.T) {
	data := []byte("123456")
	key := []byte("123456789012345678901234")

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	paddingData := PKCS5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	result := make([]byte, len(paddingData))
	blockMode.CryptBlocks(result, paddingData)
	fmt.Println(result, len(result)) // [212 95 28 124 195 33 94 132] 8

	expect, err := goEncrypt.TripleDesEncrypt(data, key, []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, expect, result)
}

func Test3DesCbcDecrypt(t *testing.T) {
	encryptData := []byte{212, 95, 28, 124, 195, 33, 94, 132}
	key := []byte("123456789012345678901234")

	block, err := des.NewTripleDESCipher(key)
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

	except, err := goEncrypt.TripleDesDecrypt(encryptData, key, []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, except, result)
}
