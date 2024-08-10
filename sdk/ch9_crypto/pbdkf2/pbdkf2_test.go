package pbdkf2

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"golang.org/x/crypto/pbkdf2"
)

// pbdkf2 (Password-Based Key Derivation Function 2) 是一种基于密码的密钥派生函数, 它通过迭代地应用哈希函数来增强密码的安全性.
// PBKDF2算法的工作原理是将用户密码与一个随机生成的盐值（salt）结合, 然后重复应用一个伪随机函数（通常是HMAC）多次, 以生成密钥.
// 这个过程被称为密钥加强, 可以显著增加破解密码所需的计算成本和时间, 有效防止暴力攻击和彩虹表攻击.

// func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte
// password: 用户密码
// salt: 随机盐值
// iter: 迭代次数 (次数约多加密和破解的耗时越长)
// keyLen: 生成的密文字节数组长度
// h: 加密使用的 hash 函数
func TestPbdkf2(t *testing.T) {
	rawPwd := "123456"
	salt := genRandomSalt(32)
	iter := 10000
	keyLen := 32
	key := pbkdf2.Key([]byte(rawPwd), salt, iter, keyLen, sha256.New)
	fmt.Println("salt:", len(salt), hex.EncodeToString(salt))
	fmt.Println("key:", len(key), hex.EncodeToString(key))

	// Output:
	// salt: 32 84ac8410d2ecb982b0e3734258b232aa1a9a3eabdef4ad1935a84e38451b79ca
	// key: 32 e52b34d8733b8c50cdf65ef9cb9cd82ca96fa35205306fb3b34b76f929e4b0e4
}

func genRandomSalt(size int) []byte {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}
