package zlib

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// compress/zlib

// 错误
func TestErr(t *testing.T) {
	_ = zlib.ErrChecksum
	_ = zlib.ErrDictionary
	_ = zlib.ErrHeader
}

// 常量
func TestConst(t *testing.T) {
	// 压缩级别
	_ = zlib.NoCompression      // flate.NoCompression
	_ = zlib.BestSpeed          // flate.BestSpeed
	_ = zlib.BestCompression    // flate.BestCompression
	_ = zlib.DefaultCompression // flate.DefaultCompression
	_ = zlib.HuffmanOnly        // flate.HuffmanOnly
}

// 创建zlib的reader方法
// func NewReader(r io.Reader) (io.ReadCloser, error)
// func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)

// 创建zlib的writer方法
// func NewWriter(w io.Writer) *Writer
// func NewWriterLevel(w io.Writer, level int) (*Writer, error)
// func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)

// 读写
func TestWriteAndRead(t *testing.T) {
	// 创建zlib的io.Writer
	var buf1 bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buf1, zlib.BestCompression)
	if err != nil {
		panic(err)
	}

	// 基于zlib压缩写入数据
	src := "hello world."
	n, err := writer.Write([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)

	// 刷新底层数据
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	// 注意: 写入完成即关闭io.Writer,保证正常EOF
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 创建zlib的io.ReadCloser
	readCloser, err := zlib.NewReader(&buf1)
	if err != nil {
		panic(err)
	}
	defer readCloser.Close()

	var buf2 bytes.Buffer
	nn, err := buf2.ReadFrom(readCloser)
	if err != nil {
		panic(err)
	}
	fmt.Println("成功读取字节数:", nn)
	assert.Equal(t, src, buf2.String())
	// output:
	// 成功写入字节数: 12
	// 成功读取字节数: 12
}
