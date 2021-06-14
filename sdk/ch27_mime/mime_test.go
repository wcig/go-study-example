package ch27_mime

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"testing"
)

// mime: 实现了MIME标准的一部分

// 常量
func TestConst(t *testing.T) {
	_ = mime.BEncoding // RFC2045定义的Base64编码方案
	_ = mime.QEncoding // RFC2047定义的Q编码方案
}

// 变量
func TestVar(t *testing.T) {
	_ = mime.ErrInvalidMediaParameter // mime媒体类型解析错误
}

// func AddExtensionType(ext, typ string) error：将扩展名ext与MIME类型typ管理，扩展名以'.'开头。
func TestAddExtensionType(t *testing.T) {
	typ := "image/apng"
	ext := ".apng"

	result, err := mime.ExtensionsByType(typ)
	fmt.Println(result, err)

	err = mime.AddExtensionType(ext, typ)
	if err != nil {
		panic(err)
	}

	result, err = mime.ExtensionsByType(typ)
	fmt.Println(result, err)
	// output:
	// [] <nil>
	// [.apng] <nil>
}

// func ExtensionsByType(typ string) ([]string, error)：返回已知与MIME类型type管理的扩展，返回的扩展以'.'开头，没有则返回nil。
func TestExtensionsByType(t *testing.T) {
	result, err := mime.ExtensionsByType("image/jpeg")
	fmt.Println(result)
	fmt.Println(err)
	// output:
	// [.jpe .jpeg .jpg]
	// <nil>
}

// func FormatMediaType(t string, param map[string]string) string：将媒体类型t和参数params序列化为符合RFC2045和RFC2616的媒体类型。
func TestFormatMediaType(t *testing.T) {
	mediaType := "text/html"
	params := map[string]string{
		"charset": "utf-8",
	}
	result := mime.FormatMediaType(mediaType, params)
	fmt.Println(result) // text/html; charset=utf-8
}

// func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
// 根据RFC1521解析媒体类型值和可选参数，媒体类型是 Content-Type 和 Content-Disposition 标头 (RFC 2183) 中的值。
func TestParseMediaType(t *testing.T) {
	mediaType, params, err := mime.ParseMediaType("text/html; charset=utf-8")
	if err != nil {
		panic(err)
	}
	fmt.Println("media type:", mediaType)
	fmt.Println("params:", params)
	// output:
	// media type: text/html
	// params: map[charset:utf-8]
}

// func TypeByExtension(ext string) string：返回与文件扩展名ext关联的MIME类型，扩展名以'.'开头，没有类型匹配返回空字符串。
// 内置表很小，在unix下会使用本地系统的mime.types进行扩展：/etc/mime.types, /etc/apache2/mime.types, /etc/apache/mime.types，windows是从注册表获取，默认文本字符集参数设置为"utf-8"
func TestTypeByExtension(t *testing.T) {
	typ := mime.TypeByExtension(".jpg")
	fmt.Println(typ) // image/jpeg
}

// 类型：mime.WordDecoder 解码包含 RFC 2047 编码字的 MIME 标头。
func TestTypeWordDecoder(t *testing.T) {
	_ = mime.WordDecoder{}
}

// func (d *WordDecoder) Decode(word string) (string, error)：解码RFC 2047编码字。
func TestDecode(t *testing.T) {
	dec := new(mime.WordDecoder)
	header, err := dec.Decode("=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := io.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}
	header, err = dec.Decode("=?x-case?q?hello!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)
	// output:
	// ¡Hola, señor!
	// HELLO!
}

// func (d *WordDecoder) DecodeHeader(header string) (string, error)
// 解码给定字符串的所有编码字。当且仅当 d 的 CharsetReader 返回错误时，它才返回错误。
func TestDecodeHeader(t *testing.T) {
	dec := new(mime.WordDecoder)
	header, err := dec.DecodeHeader("=?utf-8?q?=C3=89ric?= <eric@example.org>, =?utf-8?q?Ana=C3=AFs?= <anais@example.org>")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	header, err = dec.DecodeHeader("=?utf-8?q?=C2=A1Hola,?= =?utf-8?q?_se=C3=B1or!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := io.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}
	header, err = dec.DecodeHeader("=?x-case?q?hello_?= =?x-case?q?world!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)
	// output:
	// Éric <eric@example.org>, Anaïs <anais@example.org>
	// ¡Hola, señor!
	// HELLO WORLD!
}

// 类型：mime.WordEncoder RFC 2047编码字编码器
func TestTypeWordEncoder(t *testing.T) {
	_ = mime.WordEncoder('b')
}

// func (e WordEncoder) Encode(charset, s string) string：返回s的编码字形式。
func TestEncode(t *testing.T) {
	fmt.Println(mime.QEncoding.Encode("utf-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("utf-8", "Hello!"))
	fmt.Println(mime.BEncoding.Encode("UTF-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("ISO-8859-1", "Caf\xE9"))
	// output:
	// =?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=
	// Hello!
	// =?UTF-8?b?wqFIb2xhLCBzZcOxb3Ih?=
	// =?ISO-8859-1?q?Caf=E9?=
}
