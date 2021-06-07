package base64

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/base64
// base64实现了RFC 4648规定的base64编码。

// base64默认的4种编码器变量
func Test4EncodingVariable(t *testing.T) {
	// RFC 4648定义的标准base64编码字符集
	_ = base64.StdEncoding
	// RFC 4648定义的另一个base64编码字符集,主要用于url和文件名
	_ = base64.URLEncoding
	// 与StdEncoding类似,区别在于省略填充字符
	_ = base64.RawStdEncoding
	// 与URLEncoding类似,区别在于省略填充字符
	_ = base64.RawURLEncoding
}

// func (enc *Encoding) EncodeToString(src []byte) string
// base64编码
func TestEncodeToString(t *testing.T) {
	val := base64.StdEncoding.EncodeToString([]byte("hello"))
	assert.Equal(t, "aGVsbG8=", val)
}

// func (enc *Encoding) EncodedLen(n int) int
// 返回n字节base64编码的数据解码后的最大长度。
func TestEncodedLen(t *testing.T) {
	src := []byte("hello")
	fmt.Println(base64.StdEncoding.EncodedLen(len(src))) // 8
}

// func (enc *Encoding) Encode(dst, src []byte)
// base64编码
// Encode使用编码enc对src进行编码，将EncodedLen(len(src))字节写入dst。
// 编码会将输出填充为4字节的倍数，因此Encode不适合在大型数据流的各个块上使用。 请改用NewEncoder()。
func TestEncode(t *testing.T) {
	src := []byte("hello")
	size := base64.StdEncoding.EncodedLen(len(src))
	dst := make([]byte, size)
	base64.StdEncoding.Encode(dst, src)
	assert.Equal(t, "aGVsbG8=", string(dst))
	assert.Equal(t, len([]byte("aGVsbG8=")), len(dst))
}

// func (enc *Encoding) DecodeString(s string) ([]byte, error)
// base64解码
func TestDecodeString(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("aGVsbG8=")
	assert.Nil(t, err)
	assert.Equal(t, "hello", string(b))
}

// func (enc *Encoding) DecodedLen(n int) int
// DecodedLen返回与n个字节的base64编码数据相对应的解码数据的最大长度（以字节为单位）
func TestDecodedLen(t *testing.T) {
	src := []byte("aGVsbG8=")
	fmt.Println(base64.StdEncoding.DecodedLen(len(src))) // 6
}

// func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
// base64解码
// 它最多将DecodedLen(len(src))字节写入dst，并返回写入的字节数。
// 如果src包含无效的base64数据，它将返回成功写入的字节数和CorruptInputError。 换行符（\ r和\ n）将被忽略。
func TestDecode(t *testing.T) {
	src := []byte("aGVsbG8=")
	size := base64.StdEncoding.DecodedLen(len(src))
	dst := make([]byte, size)
	n, err := base64.StdEncoding.Decode(dst, src)
	assert.Nil(t, err)
	fmt.Println(n, string(dst[0:n])) // 5 hello
}

// func (enc Encoding) Strict() *Encoding
// 除启用严格解码外，严格创建与enc相同的新编码。 在这种模式下，解码器要求尾随填充位为零，如RFC 4648第3.5节所述。
func TestStrict(t *testing.T) {
	encoding := base64.StdEncoding.Strict()
	assert.NotNil(t, encoding)
}

// func (enc Encoding) WithPadding(padding rune) *Encoding
// WithPadding将创建一个与enc相同的新编码，但带有指定的填充字符，或者使用NoPadding禁用填充。 填充字符不得为'\ r'或'\ n'，且不得包含在编码字母中，且其符文必须等于或小于'\ xff'。
func TestWithPadding(t *testing.T) {
	encoding := base64.StdEncoding.WithPadding(rune('@'))
	assert.NotNil(t, encoding)
}

// func NewEncoding(encoder string) *Encoding
// NewEncoding返回由给定字母定义的新的填充编码，该编码必须是不包含填充字符或CR / LF（'\ r'，'\ n'）的64字节字符串。
// 生成的编码使用默认的填充字符（'='），可以通过WithPadding更改或禁用默认的填充字符。
func TestNewEncoding(t *testing.T) {
	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	encoding := base64.NewEncoding(encodeStd)
	println(encoding.EncodeToString([]byte("hello"))) // aGVsbG8=
}

// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
// NewEncoder返回一个新的base64流编码器。写入返回的写入器的数据将使用enc进行编码，然后写入w。
// Base64编码以4字节块为单位；完成写入后，调用者必须关闭返回的编码器以刷新任何部分写入的块。
func TestNewEncoder(t *testing.T) {
	encoding := base64.StdEncoding
	bf := &bytes.Buffer{}
	writer := base64.NewEncoder(encoding, bf)
	n, err := writer.Write([]byte("hello"))
	_ = writer.Close()
	fmt.Println(n, err, bf.String()) // 5 <nil> aGVsbG8=
}

// func NewDecoder(enc *Encoding, r io.Reader) io.Reader
// NewDecoder构造一个新的base64流解码器。
func TestNewDecoder(t *testing.T) {
	encoding := base64.StdEncoding
	src := "aGVsbG8="
	reader := base64.NewDecoder(encoding, strings.NewReader(src))
	out := make([]byte, encoding.DecodedLen(len([]byte(src))))
	n, err := reader.Read(out)
	fmt.Println(n, err, string(out[0:n]))
}
