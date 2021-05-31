package flate

import (
	"bytes"
	"compress/flate"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// compress/flate
// flate包实现了deflate压缩数据格式，参见RFC 1951。gzip包和zlib包实现了对基于deflate的文件格式的访问。
// DEFLATE是同时使用了LZ77算法与哈夫曼编码（Huffman Coding）的一个无损数据压缩算法。

// 压缩参数常量
func TestConstCompressLevel(t *testing.T) {
	_ = flate.BestCompression    // 最佳压缩比
	_ = flate.BestSpeed          // 最快速度压缩
	_ = flate.NoCompression      // 不压缩
	_ = flate.DefaultCompression // 默认压缩
}

// func NewReader(r io.Reader) io.ReadCloser
// 返回一个新的ReadCloser，可用于读取r的未压缩版本。 如果r还没有实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。
// 完成阅读后调用 ReadCloser 上的 Close 是调用者的责任。 NewReader 返回的 ReadCloser 也实现了 Resetter。
func TestNewReader(t *testing.T) {
	readCloser := flate.NewReader(strings.NewReader("hello"))
	defer readCloser.Close()
	assert.NotNil(t, readCloser)
}

// func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
// 与 NewReader 类似，但使用预设字典初始化阅读器。 返回的 Reader 表现得好像未压缩的数据流以给定的字典开始，该字典已经被读取。
// NewReaderDict 通常用于读取由 NewWriterDict 压缩的数据。 NewReader 返回的 ReadCloser 也实现了 Resetter。
func TestNewReaderDict(t *testing.T) {
	readCloser := flate.NewReaderDict(strings.NewReader("hello"), []byte("key"))
	defer readCloser.Close()
	assert.NotNil(t, readCloser)
}

// func NewWriter(w io.Writer, level int) (*Writer, error)
// 返回在给定级别压缩数据的新 Writer。 在 zlib 之后，级别范围从 1 (BestSpeed) 到 9 (BestCompression)；
// 较高级别通常运行速度较慢但压缩更多。 级别 0 (NoCompression) 不尝试任何压缩； 它只添加了必要的 DEFLATE 框架。 级别-1（DefaultCompression）使用默认压缩级别。 Level -2 (HuffmanOnly) 将仅使用 Huffman 压缩，为所有类型的输入提供非常快的压缩，但牺牲了相当大的压缩效率。
// 如果级别在 [-2, 9] 范围内，则返回的错误将为 nil。 否则返回的错误将非零。
func TestNewWriter(t *testing.T) {
	writer, err := flate.NewWriter(os.Stdout, flate.BestCompression)
	defer writer.Close()
	assert.Nil(t, err)
	assert.NotNil(t, writer)
}

// func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
// 与NewWriter一样，唯一区别在初始化添加预制字典
func TestNewWriterDict(t *testing.T) {
	writer, err := flate.NewWriterDict(os.Stdout, flate.BestSpeed, []byte("key"))
	defer writer.Close()
	assert.Nil(t, err)
	assert.NotNil(t, writer)
}

// flate.Writer 方法:
// func (w *Writer) Close() error
// func (w *Writer) Flush() error
// func (w *Writer) Reset(dst io.Writer)
// func (w *Writer) Write(data []byte) (n int, err error)

// 读写示例
func TestReadWrite(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	w, err := flate.NewWriter(buf, flate.BestSpeed)
	assert.Nil(t, err)
	defer w.Close()

	content := "hello world."
	n, err := w.Write([]byte(content))
	assert.Nil(t, err)
	fmt.Println("成功写入字节数:", n)

	err = w.Flush()
	assert.Nil(t, err)

	rc := flate.NewReader(buf)
	defer rc.Close()

	val := make([]byte, n)
	nn, err := rc.Read(val)
	assert.Nil(t, err)
	assert.Equal(t, content, string(val))
	fmt.Println("成功读取字节数:", nn)
	// output:
	// 成功写入字节数: 12
	// 成功读取字节数: 12
}

// 带字典读写示例
func TestReadWriteWithDict(t *testing.T) {
	dict := []byte("key")

	buf := bytes.NewBuffer(nil)
	w, err := flate.NewWriterDict(buf, flate.BestSpeed, dict)
	assert.Nil(t, err)
	defer w.Close()

	content := "hello world."
	n, err := w.Write([]byte(content))
	assert.Nil(t, err)
	fmt.Println("成功写入字节数:", n)

	err = w.Flush()
	assert.Nil(t, err)

	rc := flate.NewReaderDict(buf, dict)
	defer rc.Close()

	val := make([]byte, n)
	nn, err := rc.Read(val)
	assert.Nil(t, err)
	assert.Equal(t, content, string(val))
	fmt.Println("成功读取字节数:", nn)
	// output:
	// 成功写入字节数: 12
	// 成功读取字节数: 12
}
