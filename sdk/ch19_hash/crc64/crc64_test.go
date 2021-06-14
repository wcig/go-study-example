package crc64

import (
	"fmt"
	"hash/crc64"
	"testing"
)

// hash/crc64：实现64位循环冗余校验或CRC-64校验和。

// 常量
func TestConst(t *testing.T) {
	_ = crc64.Size // 8 CRC64校验和大小（字节）

	var isoPoly uint64 = crc64.ISO   // ISO多项式
	var ecmaPoly uint64 = crc64.ECMA // ECMA多项式
	fmt.Println(isoPoly, ecmaPoly)
}

// func MakeTable(poly uint64) *Table：根据指定表达式创建表
func TestMakeTable(t *testing.T) {
	table := crc64.MakeTable(crc64.ISO)
	fmt.Println(table != nil)
}

// func Checksum(data []byte, tab *Table) uint64：计算CRC64校验和
func TestChecksum(t *testing.T) {
	src := "hello world."
	checksum := crc64.Checksum([]byte(src), crc64.MakeTable(crc64.ISO))
	fmt.Println(checksum) // 8370449253397092780
}

// func New(tab *Table) hash.Hash64：创建hash.Hash64
func TestNew(t *testing.T) {
	src := "hello world."
	hash64 := crc64.New(crc64.MakeTable(crc64.ISO))
	n, err := hash64.Write([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)
	checksum := hash64.Sum64()
	fmt.Println(checksum)
	// output:
	// 成功写入字节数: 12
	// 8370449253397092780
}

//
func TestUpdate(t *testing.T) {
	src1 := "hello "
	src2 := "world."
	table := crc64.MakeTable(crc64.ISO)

	checksum1 := crc64.Checksum([]byte(src1), table)
	fmt.Println(checksum1)

	checksum2 := crc64.Update(checksum1, table, []byte(src2))
	fmt.Println(checksum2)
	// output:
	// 6461608761098768384
	// 8370449253397092780
}
