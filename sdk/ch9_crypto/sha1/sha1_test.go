package sha1

import (
	"crypto/sha1"
	"testing"
)

// crypto/sha1：实现SHA1哈希算法

// 常量
func TestConst(t *testing.T) {
	_ = sha1.BlockSize // 64 SHA1块大小（字节）
	_ = sha1.Size      // 20 SHA1校验和大小（字节）
}

// 函数
// func New() hash.Hash             // 创建计算SHA1校验和的hash.Hash
// func Sum(data []byte) [Size]byte // 返回data的SHA1的校验和
