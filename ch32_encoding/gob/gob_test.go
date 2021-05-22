package gob

import (
	"encoding/gob"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/gob
// gob(go binary)文件编解码器,可以作为rpc调用的参数使用

type user struct {
	Id   int
	Name string
}

// func NewEncoder(w io.Writer) *Encoder
// 返回一个新的编码器，该编码器将在io.Writer上传输。
func TestNewEncoder(t *testing.T) {
	encoder := gob.NewEncoder(os.Stdout)
	assert.NotNil(t, encoder)
}

// func NewDecoder(r io.Reader) *Decoder
// 返回一个新的解码器，该解码器从io.Reader中读取。 如果r还没有实现io.ByteReader，它将包装在bufio.Reader中。
func TestNewDecoder(t *testing.T) {
	decoder := gob.NewDecoder(os.Stdin)
	assert.NotNil(t, decoder)
}

// func (enc *Encoder) Encode(e interface{}) error
// Encode传输由空接口值表示的数据项，从而确保所有必需的类型信息都已首先传输。 将nil指针传递给Encoder会感到恐慌，因为它们无法通过gob传输。
func TestEncode(t *testing.T) {
	file, err := os.Create("tmp.gob")
	assert.Nil(t, err)
	defer file.Close()

	encoder := gob.NewEncoder(file)
	u := &user{Id: 1, Name: "tom"}
	err = encoder.Encode(u)
	assert.Nil(t, err)
}

// func (dec *Decoder) Decode(e interface{}) error
// 解码从输入流中读取下一个值，并将其存储在由空接口值表示的数据中。 如果e为nil，则该值将被丢弃。 否则，e底下的值必须是指向接收到的下一个数据项的正确类型的指针。
// 如果输入在EOF处，则Decode返回io.EOF且不修改e
func TestDecode(t *testing.T) {
	file, err := os.Open("tmp.gob")
	assert.Nil(t, err)
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var u user
	err = decoder.Decode(&u)
	assert.Nil(t, err)
	fmt.Println(u)
}
