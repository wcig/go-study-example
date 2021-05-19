package base32

import (
	"encoding/base32"
	"fmt"
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
func Test(t *testing.T) {
	src := []byte("NBSWY3DP")
	fmt.Println(base32.StdEncoding.DecodedLen(len(src))) // 5
}
