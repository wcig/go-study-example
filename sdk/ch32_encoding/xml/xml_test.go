package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// encoding/xml
// xml文件构造和解析

type user struct {
	XMLName    xml.Name `xml:"user"`
	Id         int      `xml:"id"`
	Name       string   `xml:"name"`
	CreateTime int64    `xml:"create_time"`
}

// func Marshal(v interface{}) ([]byte, error)
// 生成v的xml编码
func TestMarshal(t *testing.T) {
	user := &user{Id: 1, Name: "<>tom", CreateTime: time.Now().Unix()}
	bytes, err := xml.Marshal(user)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
	// output:
	// <user><id>1</id><name>tom</name><create_time>1621746536</create_time></user>
}

// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
// 格式化方式生成v的xml编码
func TestMarshalIndent(t *testing.T) {
	u := &user{Id: 1, Name: "tom", CreateTime: time.Now().Unix()}
	b, err := xml.MarshalIndent(u, "", "\t")
	assert.Nil(t, err)
	fmt.Println(string(b))
	// output:
	// <user>
	//	<id>1</id>
	//	<name>tom</name>
	//	<create_time>1621746611</create_time>
	// </user>
}

// func Unmarshal(data []byte, v interface{}) error
// xml解码
func TestUnmarshal(t *testing.T) {
	src := "<user><id>1</id><name>tom</name><create_time>1621746536</create_time></user>"
	u := &user{}
	err := xml.Unmarshal([]byte(src), u)
	assert.Nil(t, err)
	fmt.Println(u) // &{{ user} 1 tom 1621746536}
}

// func CopyToken(t Token) Token
// 返回token的拷贝 (深拷贝)
func TestCopyToken(t *testing.T) {
	src := []byte("hello")
	token := xml.CharData(src)

	copyToken := xml.CopyToken(token)
	fmt.Printf("%v, %p, %p\n", token, token, copyToken) // [104 101 108 108 111], 0xc0000a2b00, 0xc0000a2b05
	assert.True(t, reflect.DeepEqual(token, copyToken))

	src[0] = 'a'
	assert.True(t, !reflect.DeepEqual(token, copyToken))
}

// func EscapeText(w io.Writer, s []byte) error
// EscapeText将适当地转义的XML等效于纯文本数据写入w。
func TestEscapeText(t *testing.T) {
	var buf bytes.Buffer
	err := xml.EscapeText(&buf, []byte("<div>"))
	assert.Nil(t, err)
	fmt.Println(buf.String()) // &lt;div&gt;
}

// func Escape(w io.Writer, s []byte)
// 等价于EscapeText(),区别在于忽略错误 (不建议使用)
func TestEscape(t *testing.T) {
	var buf bytes.Buffer
	xml.Escape(&buf, []byte("<div>"))
	fmt.Println(buf.String()) // &lt;div&gt;
}

// func NewEncoder(w io.Writer) *Encoder
// 返回一个包装w后的编码器
func TestNewEncoder(t *testing.T) {
	u := &user{Id: 1, Name: "tom", CreateTime: time.Now().Unix()}

	// func (enc *Encoder) Encode(v interface{}) error
	// 编码器将v的XML编码写入流中,在返回数据时已调用Flush()
	var buf1 bytes.Buffer
	ec1 := xml.NewEncoder(&buf1)
	err1 := ec1.Encode(u)
	assert.Nil(t, err1)
	fmt.Println(buf1.String())
	// output:
	// <user><id>1</id><name>tom</name><create_time>1621746904</create_time></user>

	// func (enc *Encoder) Indent(prefix, indent string)
	// 设置编码器格式化风格
	var buf2 bytes.Buffer
	ec2 := xml.NewEncoder(&buf2)
	ec2.Indent("", "\t")
	_ = ec2.Encode(u)
	fmt.Println(buf2.String())
	// output:
	// <user>
	//	<id>1</id>
	//	<name>tom</name>
	//	<create_time>1621748731</create_time>
	// </user>

	// func (enc *Encoder) EncodeToken(t Token) error
	// EncodeToken将给定的XML令牌写入流中。如果StartElement和EndElement标记不正确匹配，它将返回错误。(没有必要调用Flush())
	var buf3 bytes.Buffer
	ec3 := xml.NewEncoder(&buf3)
	err3 := ec3.EncodeToken(xml.CharData("hello"))
	assert.Nil(t, err3)
	_ = ec3.Encode(u)
	fmt.Println(buf3.String())
	// output:
	// hello<user><id>1</id><name>tom</name><create_time>1621761496</create_time></user>

	// func (enc *Encoder) EncodeElement(v interface{}, start StartElement) error
	// EncodeElement使用start作为编码中最外面的标记，将v的XML编码写入流中。
	// err := ec1.EncodeElement(u, start)

	// func (enc *Encoder) Flush() error
	// 刷新将所有缓冲的XML刷新到基础编写器。
	// err := ec1.Flush()
}

// func NewTokenDecoder(t TokenReader) *Decoder
// NewTokenDecoder使用基础令牌流创建一个新的XML解析器。
func TestNewTokenDecoder(t *testing.T) {
	decoder := xml.NewTokenDecoder(xml.NewDecoder(os.Stdin))
	assert.NotNil(t, decoder)
}

// xml设置属性值、子元素、可忽略元素、注释
func TestXmlOptions(t *testing.T) {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{City: "Hanga Roa", State: "Easter Island"}

	output, err := xml.MarshalIndent(v, "", "\t")
	assert.Nil(t, err)
	fmt.Println(string(output))
	// output:
	// <person id="13">
	//	<name>
	//		<first>John</first>
	//		<last>Doe</last>
	//	</name>
	//	<age>42</age>
	//	<Married>false</Married>
	//	<City>Hanga Roa</City>
	//	<State>Easter Island</State>
	//	<!-- Need more details. -->
	// </person>
}
