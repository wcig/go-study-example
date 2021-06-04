package lzw

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// compress/lzw

// 常量
func TestConst(t *testing.T) {
	_ = lzw.LSB // lsb排序,主要用于gif
	_ = lzw.MSB // msb排序,主要用于tiff,pdf
}

// 读写
func TestWriteAndRead(t *testing.T) {
	// 创建lzw的io.WriteCloser
	var buf bytes.Buffer
	writeCloser := lzw.NewWriter(&buf, lzw.LSB, 8)

	// 基于lzw压缩写入数据
	src := "hello world."
	n, err := writeCloser.Write([]byte(src))
	if err != nil {
		panic(err)
	}

	// 写入完成即关闭io.WriteCloser,保证正常EOF
	err = writeCloser.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)

	// 创建lzw的io.ReadCloser
	readCloser := lzw.NewReader(&buf, lzw.LSB, 8)
	defer readCloser.Close()

	// 基于lzw解压读取所有数据
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
