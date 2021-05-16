package ascii85

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// encoding/ascii85
// ascii85包实现了ascii85数据编码（5个ascii字符表示4个字节），该编码用于btoa工具和Adobe的PostScript语言和PDF文档格式。

// func Encode(dst, src []byte) int
// Encode将src最多编码为dst的MaxEncodedLen（len（src））字节，返回实际写入的字节数。
// 该编码处理4字节的块，对最后一个片段使用特殊的编码，因此Encode不适合用于大型数据流的各个块。 请改用NewEncoder（）。
func TestEncode(t *testing.T) {
	src := []byte("hello明天")
	dst := make([]byte, ascii85.MaxEncodedLen(len(src)))
	n := ascii85.Encode(dst, src)
	fmt.Println(n, len(dst), string(dst)) // 14 15 BOu!rDs!FIjeSjU
}

// func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)
// 解码将src解码为dst，同时返回写入dst的字节数和从src消耗的数。 如果src包含无效的ascii85数据，则Decode将返回成功写入的字节数和CorruptInputError。 解码会忽略src中的空格和控制字符。 通常，ASCII85编码的数据被包裹在<〜和〜>符号中。 解码期望这些已被调用方剥离。
// 如果flush为true，则Decode假定src代表输入流的末尾并对其进行完全处理，而不是等待另一个32位块的完成。
// NewDecoder在Decode周围包装了io.Reader接口
func TestDecode(t *testing.T) {
	src := []byte("BOu!rDs!FIjeSjU")
	dst := make([]byte, 4*len(src))
	ndst, nsrc, err := ascii85.Decode(dst, src, false)
	fmt.Println(ndst, nsrc, err, string(dst[0:ndst]))
}

// func MaxEncodedLen(n int) int
// MaxEncodedLen返回n个源字节的编码的最大长度
func TestMaxEncodedLen(t *testing.T) {
	str := "hello"
	fmt.Println(ascii85.MaxEncodedLen(len([]byte(str)))) // 10
}

// func NewEncoder(w io.Writer) io.WriteCloser
// 创建一个将数据编码为ascii85流写入w的编码器。Ascii85编码算法操作32位块，写入结束后，必须调用Close方法将缓存中保留的不完整块刷新到w里。
func TestNewEncoder(t *testing.T) {
	bb := &bytes.Buffer{}
	encoder := ascii85.NewEncoder(bb)
	n, err := encoder.Write([]byte("hello明天"))
	fmt.Println(n, err)
	err = encoder.Close()
	fmt.Println(err)
	fmt.Println(bb.String())
}

// output:
// 11 <nil>
// <nil>
// BOu!rDs!FIjeSj

// func NewDecoder(r io.Reader) io.Reader
// 创建一个从r解码ascii85流的解码器
func TestNewDecoder(t *testing.T) {
	str := "BOu!rDs!FIjeSj"
	decoder := ascii85.NewDecoder(strings.NewReader(str))
	b, err := ioutil.ReadAll(decoder)
	fmt.Println(string(b), err) // hello明天 <nil>
}
