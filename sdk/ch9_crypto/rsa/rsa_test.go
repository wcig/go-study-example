package rsa

import (
	"crypto/rsa"
	"testing"
)

// crypto/rsa：包 rsa 实现了 PKCS #1 和 RFC 8017 中指定的 RSA 加密。

// 常量
func TestConstants(t *testing.T) {
	_ = rsa.PSSSaltLengthAuto       // 0
	_ = rsa.PSSSaltLengthEqualsHash // -1
}

// 变量
func TestVariables(t *testing.T) {
	_ = rsa.ErrDecryption     // crypto/rsa: decryption error
	_ = rsa.ErrMessageTooLong // crypto/rsa: message too long for RSA public key size
	_ = rsa.ErrVerification   // crypto/rsa: verification error
}

// 对应加解密
// 1.OEAP: EncryptOAEP (公钥) -> DecryptOAEP (私钥)
// 2.PKCS1v15: EncryptPKCS1v15 (公钥) -> DecryptPKCS1v15 (私钥)
// 3.SignPKCS1v15: SignPKCS1v15 (私钥) -> VerifyPKCS1v15 (公钥)
// 4.SignPSS: SignPSS (私钥) -> VerifyPSS (公钥)

// 函数
// 1.RSA-OAEP解密
// func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
// 2.RSA-OAEP加密
// func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
// 3.RSA PKCS #1 v1.5解密
// func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
// 4.RSA PKCS #1 v1.5加回话密钥解密
// func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error
// 5.RSA PKCS #1 v1.5加密
// func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
// 6.使用来自 RSA PKCS #1 v1.5 的 RSASSA-PKCS1-V1_5-SIGN 计算散列的签名
// func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
// 7.使用 PSS 计算摘要的签名。
// func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, opts *PSSOptions) ([]byte, error)
// 8.验证 RSA PKCS #1 v1.5 签名。
// func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
// 9.验证 PSS 签名
// func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error

// 类型
// type PrivateKey
// 1.使用随机源生成给定大小的RSA密钥
// func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)
// 2.使用随机源生成给定大小的多质数RSA密钥
// func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)
// 3.使用priv密钥解密密文
// func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
// 4.报告priv和x是否有相等的值
// func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool
// 5.执行计算加快未来私钥操作
// func (priv *PrivateKey) Precompute()
// 6.返回priv对应公钥
// func (priv *PrivateKey) Public() crypto.PublicKey
// 7.生成签名
// func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

// type PubicKey
// 1.报告pub和x是否有相同的值
// func (pub *PublicKey) Equal(x crypto.PublicKey) bool
// 2.大小返回以字节为单位的模数大小。此公钥的原始签名和密文将具有相同的大小。
// func (pub *PublicKey) Size() int
