package bcrypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// bcrypt: 一种加密方式 (多次原密码加随机盐加密的加密方式)
// 示例: $2a$10$YmNVMeRoiF0CvZdf/GUV9eDrGPfAazFDhA6q2DcKaV/3OSgiRL1Y.
// 2a: 标识bcrypt加密
// 10: 标识哈希成本因子(循环加密次数)
// YmNVMeRoiF0CvZdf/GUV9e: 16个字节(128bits)的salt经过base64编码得到的22长度的字符
// DrGPfAazFDhA6q2DcKaV/3OSgiRL1Y.: 24个字节(192bits)的hash,经过bash64的编码得到的31长度的字符.

// 常量
// const (
//	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
//	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
//	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
// )

// 变量
// var ErrMismatchedHashAndPassword = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
// var ErrHashTooShort = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")

// 方法
// func CompareHashAndPassword(hashedPassword, password []byte) error   // 比较加密后密码与原密码是否匹配
// func Cost(hashedPassword []byte) (int, error)                        // 返回给定加密后密码的哈希成本因子
// func GenerateFromPassword(password []byte, cost int) ([]byte, error) // 给定原密码与哈希成本因子,返回加密后密码

func TestConst(t *testing.T) {
	fmt.Println(bcrypt.MinCost)
	fmt.Println(bcrypt.MaxCost)
	fmt.Println(bcrypt.DefaultCost)
	// Output:
	// 4
	// 31
	// 10
}

func TestVar(t *testing.T) {
	fmt.Println(bcrypt.ErrMismatchedHashAndPassword)
	fmt.Println(bcrypt.ErrHashTooShort)
	// Output:
	// crypto/bcrypt: hashedPassword is not the hash of the given password
	// crypto/bcrypt: hashedSecret too short to be a bcrypted password
}

func TestExample(t *testing.T) {
	rawPwd := "123456"
	encodePwdBytes, err := bcrypt.GenerateFromPassword([]byte(rawPwd), bcrypt.DefaultCost)
	assert.Nil(t, err)
	fmt.Println("encoded password:", string(encodePwdBytes))

	cost, err := bcrypt.Cost(encodePwdBytes)
	assert.Nil(t, err)
	fmt.Println("cost:", cost)

	cost, err = bcrypt.Cost([]byte(rawPwd))
	assert.NotNil(t, err)
	fmt.Println(cost, err)

	err = bcrypt.CompareHashAndPassword(encodePwdBytes, []byte(rawPwd))
	assert.Nil(t, err)
	err = bcrypt.CompareHashAndPassword(encodePwdBytes, []byte("12345"))
	assert.NotNil(t, err)
	fmt.Println(err)
	// Output:
	// encoded password: $2a$10$YmNVMeRoiF0CvZdf/GUV9eDrGPfAazFDhA6q2DcKaV/3OSgiRL1Y.
	// cost: 10
	// 0 crypto/bcrypt: hashedSecret too short to be a bcrypted password
	// crypto/bcrypt: hashedPassword is not the hash of the given password
}
