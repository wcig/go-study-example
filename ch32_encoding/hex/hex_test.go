package hex

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/hex
// 十六进制编解码
// 建议直接使用EncodeToString(),DecodeString()方法

// hex包定义错误常量
func TestErr(t *testing.T) {
	_ = hex.ErrLength // 解码时输入字节长度错误: 为奇数而不是偶数
}

// func EncodeToString(src []byte) string
// 输出src十六进制编码后的字符串
func TestEncodeToString(t *testing.T) {
	val := hex.EncodeToString([]byte("hello"))
	assert.Equal(t, "68656c6c6f", val)
}

// func EncodedLen(n int) int
// 返回n个字节切片大小的十六进制编码后大小
func TestEncodedLen(t *testing.T) {
	src := []byte("hello")
	size := hex.EncodedLen(len(src))
	fmt.Println(size) // 10
}

// func Encode(dst, src []byte) int
// Encode将src编码为dst的EncodedLen(len(src))字节。为方便起见，它返回写入dst的字节数，但此值始终为EncodedLen(len(src))。编码实现十六进制编码。
func TestEncode(t *testing.T) {
	src := []byte("hello")
	size := hex.EncodedLen(len(src))
	dst := make([]byte, size)

	n := hex.Encode(dst, src)
	assert.Equal(t, size, n)
	fmt.Println(n, string(dst)) // 10 68656c6c6f
}

// func DecodeString(s string) ([]byte, error)
// DecodeString返回由十六进制字符串s表示的字节。DecodeString期望src仅包含十六进制字符，并且src具有偶数长度。 如果输入格式错误，则DecodeString返回错误之前解码的字节。
func TestDecodeString(t *testing.T) {
	b, err := hex.DecodeString("68656c6c6f")
	assert.Nil(t, err)
	assert.Equal(t, "hello", string(b))
}

// func DecodedLen(x int) int
// 返回十六进制解码后长度
func TestDecodedLen(t *testing.T) {
	src := []byte("68656c6c6f")
	size := hex.DecodedLen(len(src))
	fmt.Println(size) // 5
}

// func Decode(dst, src []byte) (int, error)
// Decode将src解码为DecodedLen(len(src)字节，返回写入dst的实际字节数。
// 解码期望src仅包含十六进制字符，并且src具有偶数长度。 如果输入格式错误，则Decode返回错误发生之前解码的字节数。
func TestDecode(t *testing.T) {
	src := []byte("68656c6c6f")
	size := hex.DecodedLen(len(src))
	dst := make([]byte, size)

	n, err := hex.Decode(dst, src)
	assert.Nil(t, err)
	assert.Equal(t, size, n)
	assert.Equal(t, "hello", string(dst))
}

// func NewEncoder(w io.Writer) io.Writer
// NewEncoder返回一个io.Writer，它将小写的十六进制字符写入w。
func TestNewEncoder(t *testing.T) {
	var buf bytes.Buffer
	writer := hex.NewEncoder(&buf)
	n, err := writer.Write([]byte("hello"))
	assert.Nil(t, err)
	fmt.Println(n, buf.String()) // 5 68656c6c6f
}

// func NewDecoder(r io.Reader) io.Reader
// NewDecoder返回io.Reader，该io.Reader从r解码十六进制字符。 NewDecoder期望r仅包含偶数个十六进制字符。
func TestNewDecoder(t *testing.T) {
	src := "68656c6c6f"
	reader := hex.NewDecoder(strings.NewReader(src))
	out := make([]byte, hex.DecodedLen(len([]byte(src))))
	n, err := reader.Read(out)
	assert.Nil(t, err)
	fmt.Println(n, string(out)) // 5 hello
}

// func Dump(data []byte) string
// 转储返回一个字符串，其中包含给定数据的十六进制转储。十六进制转储的格式与命令行上“ hexdump -C”的输出匹配。
func TestDump(t *testing.T) {
	src := []byte("hello")
	dump := hex.Dump(src)
	fmt.Println(dump) // 00000000  68 65 6c 6c 6f                                    |hello|
}

// func Dumper(w io.Writer) io.WriteCloser
// Dumper返回一个WriteCloser，它将所有已写入数据的十六进制转储写入w。 转储的格式与命令行上“ hexdump -C”的输出匹配。
func TestDumper(t *testing.T) {
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}

	stdoutDumper := hex.Dumper(os.Stdout)

	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}
	// output:
	// 00000000  47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|
	// 00000010  75 72 63 65 20 70 72 6f  67 72 61 6d 6d 69 6e 67  |urce programming|
	// 00000020  20 6c 61 6e 67 75 61 67  65 2e 0a 57 65 20 65 6e  | language..We en|
	// 00000030  63 6f 75 72 61 67 65 20  61 6c 6c 20 47 6f 20 75  |courage all Go u|
	// 00000040  73 65 72 73 20 74 6f 20  73 75 62 73 63 72 69 62  |sers to subscrib|
	// 00000050  65 20 74 6f 20 67 6f 6c  61 6e 67 2d 61 6e 6e 6f  |e to golang-anno|
	// 00000060  75 6e 63 65 2e                                    |unce.|
}
