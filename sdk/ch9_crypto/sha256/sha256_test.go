package sha256

import (
	"crypto/sha256"
	"testing"
)

// crypto/sha256: 实现了FIPS 180-4中定义的SHA224和SHA256哈希算法

// 常量
func TestConst(t *testing.T) {
	_ = sha256.BlockSize // 64，sha256、sha224块大小（字节）
	_ = sha256.Size      // 32，sha256校验和大小（字节）
	_ = sha256.Size224   // 28，sha224校验和大小（字节）
}

// 函数
// func New() hash.Hash                            // 创建一个hash.Hash用于计算sha256校验和
// func Sum256(data []byte) [Size]byte             // 计算data数据的sha256校验和
// func New224() hash.Hash                         // 创建一个hash.Hash用于计算sha224校验和
// func Sum224(data []byte) (sum224 [Size224]byte) // 计算data数据的sha224校验和
