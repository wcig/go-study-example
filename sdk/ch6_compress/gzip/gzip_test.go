package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// compress/gzip: gzip格式压缩文件的读写,常用于http请求或其他数据传输中的压缩解压.

// gzip包错误
func TestErr(t *testing.T) {
	_ = gzip.ErrChecksum // check sum err
	_ = gzip.ErrHeader   // header invalid
}

// 压缩级别常量
func TestCompressLevel(t *testing.T) {
	_ = gzip.NoCompression      // flate.NoCompression
	_ = gzip.BestSpeed          // flate.BestSpeed
	_ = gzip.BestCompression    // flate.BestCompression
	_ = gzip.DefaultCompression // flate.DefaultCompression
	_ = gzip.HuffmanOnly        // flate.HuffmanOnly
}

// gzip压缩
func TestWrite(t *testing.T) {
	// 创建gzip.Writer
	buf := bytes.NewBuffer(nil)
	gw := gzip.NewWriter(buf)

	// 写入数据
	n, err := gw.Write([]byte("hello world."))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)

	// 刷新数据
	if err = gw.Flush(); err != nil {
		panic(err)
	}

	// 关闭gzip.Writer (非常重要: 不能使用defer关闭,因为需要在使用gzip.Writer压缩的数据前关闭才能保证正常的EOF)
	if err = gw.Close(); err != nil {
		panic(err)
	}

	// 写入数据
	if err = os.WriteFile("tmp.txt.gz", buf.Bytes(), 0755); err != nil {
		panic(err)
	}
}

// gzip解压
func TestRead(t *testing.T) {
	file, err := os.Open("tmp.txt.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	buf := bytes.NewBuffer(nil)
	n, err := buf.ReadFrom(gr)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile("tmp.txt", buf.Bytes(), 0755); err != nil {
		panic(err)
	}
	fmt.Println("成功写入文件,写入字节数:", n)
}
