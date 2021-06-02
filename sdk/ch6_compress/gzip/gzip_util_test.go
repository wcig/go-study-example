package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试工具函数
func TestGzipGunzipData(t *testing.T) {
	src := "hello world."
	compressVal, err := GzipData([]byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip压缩数据后字节数:", len(compressVal))

	unCompressVal, err := GunzipData(compressVal)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip解压数据后字节数:", len(unCompressVal))
	assert.Equal(t, src, string(unCompressVal))
	// output:
	// gzip压缩数据后字节数: 41
	// gzip解压数据后字节数: 12
}

/* ------------------------------------------------------------------- */

// gzip压缩数据
func GzipData(data []byte) (val []byte, err error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)

	if _, err = gw.Write(data); err != nil {
		return nil, err
	}

	if err = gw.Flush(); err != nil {
		return nil, err
	}

	if err = gw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// gzip解压数据
func GunzipData(data []byte) (val []byte, err error) {
	buf := bytes.NewBuffer(data)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	var buf2 bytes.Buffer
	if _, err = buf2.ReadFrom(gr); err != nil {
		return nil, err
	}

	if err = gr.Close(); err != nil {
		return nil, err
	}

	return buf2.Bytes(), nil
}
