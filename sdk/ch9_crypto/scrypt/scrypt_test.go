package scrypt

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/scrypt"
)

// scrypt 是一种基于密码的密钥派生函数（Password-Based Key Derivation Function, PBKDF）, 由Colin Percival为其备份服务Tarsnap开发.
// 它特别设计用于抵抗暴力破解攻击, 通过需要大量内存来运行, 从而增加了尝试穷举攻击的成本. scrypt算法在2016年作为RFC 7914标准发布.

// func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
// password: 用户密码
// salt: 随机盐值
// N: 是CPU/内存的成本参数 (必须是大于1的偶数)
// r, p: 要求满足 r * p < 2³⁰
// keyLen: 生成的密文字节数组长度
// 推荐参数为: N=32768(1<<15), r=8, p=1
func TestScrypt(t *testing.T) {
	rawPwd := "123456"
	salt := genRandomSalt(32)
	key, err := scrypt.Key([]byte(rawPwd), salt, 1<<15, 8, 1, 32)
	assert.Nil(t, err)
	fmt.Println("salt:", len(salt), hex.EncodeToString(salt))
	fmt.Println("key:", len(key), hex.EncodeToString(key))

	// Output:
	// salt: 32 64f5f7de8005b1a312f6a6f7d5346e63244e3b6178bdc90e243abcf2e4889e83
	// key: 32 9927591ba0a14b57d6287aec32b76a86bedc97f4030312a41411a1e8eb4b6efb
}

func genRandomSalt(size int) []byte {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}
