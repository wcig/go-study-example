package crc32

import (
	"fmt"
	"hash/crc32"
	"testing"
)

// hash/crc32：实现32位循环冗余校验或CRC-32校验和。

// 常量
func TestConst(t *testing.T) {
	_ = crc32.Size // CRC32校验和的大小（字节）

	_ = crc32.IEEE       // 最常用的CRC32多项式（用于ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...）
	_ = crc32.Castagnoli // Castagnoli多项式，具有比IEEE更好的错误检测特性
	_ = crc32.Koopman    // Koopman多项式，具有比IEEE更好的错误检测特性
}

// 变量
func TestVar(t *testing.T) {
	_ = crc32.IEEETable // IEEE多项式表
}

// func MakeTable(poly uint32) *Table：创建表
func TestCreateTable(t *testing.T) {
	table := crc32.MakeTable(crc32.IEEE)
	fmt.Println(table != nil)
}

// func Checksum(data []byte, tab *Table) uint32：计算CRC32校验和
func TestChecksum(t *testing.T) {
	src := "hello world."
	checksum := crc32.Checksum([]byte(src), crc32.IEEETable)
	fmt.Println(checksum) // 2467028988
}

// func ChecksumIEEE(data []byte) uint32：使用IEEE多项式计算的CRC32校验和
func TestChecksumIEEE(t *testing.T) {
	src := "hello world."
	checksum := crc32.ChecksumIEEE([]byte(src))
	fmt.Println(checksum) // 2467028988
}

// func New(tab *Table) hash.Hash32：根据指定表创建hash.Hash32
func TestNew(t *testing.T) {
	src := "hello world."
	hash32 := crc32.New(crc32.IEEETable)
	n, err := hash32.Write([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)
	checksum := hash32.Sum32()
	fmt.Println(checksum)
	// output:
	// 成功写入字节数: 12
	// 2467028988
}

// func NewIEEE() hash.Hash32：使用IEEE多项式创建hash.Hash32
func TestNewIEEE(t *testing.T) {
	src := "hello world."
	hash32 := crc32.NewIEEE()
	n, err := hash32.Write([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)
	checksum := hash32.Sum32()
	fmt.Println(checksum)
	// output:
	// 成功写入字节数: 12
	// 2467028988
}

// func Update(crc uint32, tab *Table, p []byte) uint32：返回添加p字节数据给crc重新计算结果
func TestUpdate(t *testing.T) {
	src1 := "hello "
	src2 := "world."

	checksum1 := crc32.ChecksumIEEE([]byte(src1))
	fmt.Println(checksum1)

	checksum2 := crc32.Update(checksum1, crc32.IEEETable, []byte(src2))
	fmt.Println(checksum2)
	// output:
	// 3984718326
	// 2467028988
}
