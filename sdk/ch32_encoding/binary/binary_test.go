package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/binary
// 二进制程序包实现数字和字节序列之间的简单转换以及varint的编码和解码。
// 通过读取和写入固定大小的值来转换数字。 固定大小的值可以是固定大小的算术类型（bool，int8，uint8，int16，float32，complex64等），也可以是仅包含固定大小值的数组或结构。
// varint函数使用可变长度编码对单个整数值进行编码和解码； 较小的值需要较少的字节。 有关规范，请参见https://developers.google.com/protocol-buffers/docs/encoding。
// 该软件包偏向于简单而不是效率。 需要高性能序列化的客户，特别是对于大型数据结构的客户，应该查看更高级的解决方案，例如编码/目标包或协议缓冲区。

// func Write(w io.Writer, order ByteOrder, data interface{}) error
// w:写入数据地方, order:排序, data:需要写入的数据
func TestWrite(t *testing.T) {
	bf := &bytes.Buffer{}
	var pi float64 = math.Pi
	err := binary.Write(bf, binary.LittleEndian, pi)
	assert.Nil(t, err)
	fmt.Println(bf.Bytes()) // [24 45 68 84 251 33 9 64]
}

// func Read(r io.Reader, order ByteOrder, data interface{}) error
// r:读取数据, order:排序, data:读取数据后需要写入的地方
func TestReader(t *testing.T) {
	var pi float64
	src := []byte{24, 45, 68, 84, 251, 33, 9, 64}
	buf := bytes.NewReader(src)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	assert.Nil(t, err)
	fmt.Println(pi) // 3.141592653589793
}
