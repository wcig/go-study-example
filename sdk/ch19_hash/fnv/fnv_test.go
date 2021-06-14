package fnv

import (
	"fmt"
	"hash"
	"hash/fnv"
	"testing"
)

// hash/fnv：实现FNV-1、FNV-1a的非加密哈希函数

// 创建32、64、128位的FNV-1、FNV-1a的哈希
func TestNew(t *testing.T) {
	var (
		h32 hash.Hash32
		h64 hash.Hash64
		h   hash.Hash
	)

	h32 = fnv.New32()
	h32 = fnv.New32a()

	h64 = fnv.New64()
	h64 = fnv.New64a()

	h = fnv.New128()
	h = fnv.New128a()

	fmt.Println(h32, h64, h)
}
