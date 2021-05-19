package base32

import (
	"bytes"
	"encoding/base32"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/base32
// base32实现了RFC 4648规定的base32编码。

// base32包默认的2种编码
func TestEncodingTmpl(t *testing.T) {
	// RFC 4648定义的标准base32编码
	_ = base32.StdEncoding
	// RFC 4648中定义的“扩展十六进制字母”。通常在DNS中使用。
	_ = base32.HexEncoding
}

// func (enc *Encoding) EncodeToString(src []byte) string
// 字节切片base32加密输出字符串
func TestEncodeToString(t *testing.T) {
	val := base32.StdEncoding.EncodeToString([]byte("hello"))
	assert.Equal(t, "NBSWY3DP", val)
}

// func (enc *Encoding) EncodedLen(n int) int
// 返回长度为n的输入缓冲区的base32编码的长度(以字节为单位)。
func TestEncodedLen(t *testing.T) {
	src := []byte("hello")
	println(base32.StdEncoding.EncodedLen(len(src))) // 8
}

// func (enc *Encoding) Encode(dst, src []byte)
// Encode使用编码enc对src进行编码，将EncodedLen(len(src))字节写入dst。
// 编码会将输出填充为8字节的倍数，因此Encode不适合在大型数据流的各个块上使用。 请改用NewEncoder()。
func TestEncode(t *testing.T) {
	src := []byte("hello")
	size := base32.StdEncoding.EncodedLen(len(src))
	dst := make([]byte, size)
	base32.StdEncoding.Encode(dst, src)
	assert.Equal(t, "NBSWY3DP", string(dst))
	assert.Equal(t, len([]byte("NBSWY3DP")), len(dst))
}

// func (enc *Encoding) DecodeString(s string) ([]byte, error)
// 字符串base32解密
func TestDecodeString(t *testing.T) {
	b, err := base32.StdEncoding.DecodeString("NBSWY3DP")
	assert.Nil(t, err)
	assert.Equal(t, "hello", string(b))
}

// func (enc *Encoding) DecodedLen(n int) int
// 返回对应于base32编码数据的n个字节的解码数据的最大长度（以字节为单位）
func TestDecodedLen(t *testing.T) {
	src := []byte("NBSWY3DP")
	fmt.Println(base32.StdEncoding.DecodedLen(len(src))) // 5
}

// func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
// 解码使用编码enc解码src。 它最多将DecodedLen(len(src))个字节写入dst，并返回写入的字节数。
// 如果src包含无效的base32数据，它将返回成功写入的字节数和CorruptInputError。 换行符（\ r和\ n）将被忽略。
func TestDecode(t *testing.T) {
	src := []byte("NBSWY3DP")
	size := base32.StdEncoding.DecodedLen(len(src))
	dst := make([]byte, size)
	n, err := base32.StdEncoding.Decode(dst, src)
	assert.Nil(t, err)
	fmt.Println(n, string(dst[0:n])) // 5 hello
}

// func (enc Encoding) WithPadding(padding rune) *Encoding
// 将创建一个与enc相同的新编码，但带有指定的填充字符，或者使用NoPadding禁用填充。 填充字符不得为'\ r'或'\ n'，且不得包含在编码字母中，且其符文必须等于或小于'\ xff'。
func TestWithPadding(t *testing.T) {
	encoding := base32.StdEncoding.WithPadding('@')
	println(encoding.EncodeToString([]byte("Hello World."))) // JBSWY3DPEBLW64TMMQXA@@@@
}

// func NewEncoding(encoder string) *Encoding
// 返回由给定字母定义的新Encoding，该字母必须是32字节的字符串
func TestNewEncoding(t *testing.T) {
	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"
	encoding := base32.NewEncoding(encodeStd)
	println(encoding.EncodeToString([]byte("hello")))
}

// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
// 返回一个新的base32流编码器。 写入返回的写入器的数据将使用enc进行编码，然后写入w。
// Base32编码以5字节块为单位； 完成写入后，调用者必须关闭返回的编码器以刷新任何部分写入的块。
func TestNewEncoder(t *testing.T) {
	bf := &bytes.Buffer{}
	writer := base32.NewEncoder(base32.StdEncoding, bf)
	n, err := writer.Write([]byte("hello"))
	_ = writer.Close()
	fmt.Println(n, err, bf.String()) // 5 <nil> NBSWY3DP
}

// func NewDecoder(enc *Encoding, r io.Reader) io.Reader
// 构造一个新的base32流解码器。
func TestNewDecoder(t *testing.T) {
	src := "NBSWY3DP"
	reader := base32.NewDecoder(base32.StdEncoding, strings.NewReader(src))
	out := make([]byte, base32.StdEncoding.DecodedLen(len([]byte(src))))
	n, err := reader.Read(out)
	fmt.Println(n, err, string(out[0:n])) // 5 <nil> hello
}
