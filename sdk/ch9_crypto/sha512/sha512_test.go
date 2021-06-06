package sha512

import (
	"crypto/sha512"
	"testing"
)

// crypto/sha512
// 实现了 FIPS 180-4 中定义的 SHA384、SHA512、SHA512/224 和 SHA512/256 哈希算法。
// 该包返回的所有 hash.Hash 实现也实现 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler 来编组和解组散列的内部状态。

// 常量
func TestConst(t *testing.T) {
	_ = sha512.BlockSize // 128 512/224, SHA512/256, SHA384 和 SHA512块大小（字节）
	_ = sha512.Size      // 64 SHA512校验和大小（字节）
	_ = sha512.Size224   // 28 SHA512/224校验和大小（字节）
	_ = sha512.Size256   // 32 SHA512/256校验和大小（字节）
	_ = sha512.Size384   // 48 SHA5384校验和大小（字节）
}

// 函数
// func New() hash.Hash                                // 返回计算SHA512校验和的hash.Hash
// func New384() hash.Hash                             // 返回计算SHA384校验和的hash.Hash
// func New512_224() hash.Hash                         //  返回计算SHA512/224校验和的hash.Hash
// func New512_256() hash.Hash                         //  返回计算SHA512/256校验和的hash.Hash
// func Sum512(data []byte) [Size]byte                 // 返回data的SHA512校验和
// func Sum384(data []byte) (sum384 [Size384]byte)     // 返回data的SHA384校验和
// func Sum512_224(data []byte) (sum224 [Size224]byte) // 返回data的SHA512/224校验和
// func Sum512_256(data []byte) (sum256 [Size256]byte) // 返回data的SHA512/256校验和
