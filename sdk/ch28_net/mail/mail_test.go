package mail

import (
	"net/mail"
	"testing"
)

// net/mail: 邮件消息解析

// 错误
func TestErr(t *testing.T) {
	_ = mail.ErrHeaderNotPresent // mail: header not in message
}

// 函数
// func ParseDate(date string) (time.Time, error) // RFC 5322日期字符串解析

// 结构体
// 1.单个邮件地址
// type Address struct {
// 	Name    string // Proper name; may be empty.
// 	Address string // user@domain
// }
// func ParseAddress(address string) (*Address, error)    // 解析单个RFC 5322字符串地址
// func ParseAddressList(list string) ([]*Address, error) // 解析RFC 5322字符串为多个地址
// func (a *Address) String() string                      // 格式化为RFC 5322地址

// 2.RFC 5322地址解析器
// type AddressParser struct {
//    // WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
//    WordDecoder *mime.WordDecoder
// }
// func (p *AddressParser) Parse(address string) (*Address, error) // 解析格式为“Gogh Fir <gf@example.com>”或“foo@example.com”的单个 RFC 5322 地址。
// func (p *AddressParser) ParseList(list string) ([]*Address, error) // 将给定的字符串解析为“Gogh Fir <gf@example.com>”或“foo@example.com”形式的逗号分隔地址列表。

// 3.邮件消息头的key-value键值对
// type Header map[string][]string
// func (h Header) AddressList(key string) ([]*Address, error) // 将头字段解析为地址列表
// func (h Header) Date() (time.Time, error)                   // 解析日期头字段
// func (h Header) Get(key string) string                      // 从头获取指定key的值，不区分大小写，没有则返回空

// 4.已解析的邮件消息
// type Message struct {
//    Header Header
//    Body   io.Reader
// }
// func ReadMessage(r io.Reader) (msg *Message, err error) // 从r读取消息，header已被解析，消息正文可以从msg.Body读取
