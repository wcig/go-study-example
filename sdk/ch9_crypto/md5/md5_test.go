package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

// crypto/md5: 包md5实现了RFC 1321中定义的MD5哈希算法。

// 常量
func Test(t *testing.T) {
	_ = md5.BlockSize // MD5块大小（字节）
	_ = md5.Size      // MD5校验和的大小（字节）
}

// 函数
// func New() hash.Hash // 返回一个新的hash.Hash来计算MD5校验和
// func Sum(data []byte) [Size]byte // 返回MD5对data校验和的结果

// 创建hash.Hash并计算MD5
func TestNew(t *testing.T) {
	m1 := md5.New()
	n, err := io.WriteString(m1, "hello world.")
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)

	val := m1.Sum(nil)
	fmt.Printf("md5: %x\n", val)
	fmt.Println("md5:", hex.EncodeToString(val))
	// output:
	// 成功写入字节数: 12
	// md5: 3c4292ae95be58e0c58e4e5511f09647
	// md5: 3c4292ae95be58e0c58e4e5511f09647
}

// 计算MD
func TestSum(t *testing.T) {
	data := "hello world."
	sum := md5.Sum([]byte(data))
	fmt.Println(sum)
	fmt.Println(fmt.Sprintf("%x", sum))
	// output:
	// [60 66 146 174 149 190 88 224 197 142 78 85 17 240 150 71]
	// 3c4292ae95be58e0c58e4e5511f09647
}
