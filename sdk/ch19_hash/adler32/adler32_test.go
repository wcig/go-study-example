package adler32

import (
	"fmt"
	"hash/adler32"
	"testing"
)

// hash/adler32 实现Adler-32校验和

// 常量
func TestConst(t *testing.T) {
	_ = adler32.Size // Adler-32校验和大小（字节）
}

// func Checksum(data []byte) uint32: 计算Adler-32校验和
func TestChecksum(t *testing.T) {
	src := "hello world."
	checksum := adler32.Checksum([]byte(src))
	fmt.Println(checksum) // 513148043
}

// func New() hash.Hash32: 创建一计算Adler-32校验和的hash.Hash32
func TestNew(t *testing.T) {
	src := "hello world."
	hash32 := adler32.New()
	fmt.Printf("init hash32 size:%d, blockSize:%d\n", hash32.Size(), hash32.BlockSize())

	n, err := hash32.Write([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)
	sum := hash32.Sum32()
	fmt.Println("checksum:", sum)
	fmt.Printf("after writer data hash32 size:%d, blockSize:%d\n", hash32.Size(), hash32.BlockSize())
}
