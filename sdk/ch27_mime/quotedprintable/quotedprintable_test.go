package quotedprintable

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"os"
	"strings"
	"testing"
)

// mime/quotedprintable: 实现了RFC 2045 指定的quoted-printable编码。
// Quoted-printable或QP encoding，没有规范的中文译名，可译为可打印字符引用编码或使用可打印字符的编码。Quoted-printable是使用可打印的ASCII字符（如字母、数字与“=”）表示各种编码格式下的字符，以便能在7-bit数据通路上传输8-bit数据, 或者更一般地说在非8-bit clean媒体上正确处理数据[注 1]。这被定义为MIME content transfer encoding，用于e-mail。
// 参考: https://zh.wikipedia.org/wiki/Quoted-printable

// Type quotedprintable.Reader: quote-printable解码器
//    func NewReader(r io.Reader) *Reader: 基于r创建一quotedprintable Reader
//    func (r *Reader) Read(p []byte) (n int, err error): 从底层Reader读取和解码quote-printable数据
func TestTypeReader(t *testing.T) {
	_ = quotedprintable.Reader{}
	// type Reader struct {
	//    // contains filtered or unexported fields
	// }
}

// Type quotedprintable.Writer: 实现io.WriteCloser的quote-printable Writer
//    func NewWriter(w io.Writer) *Writer: 基于w创建一quotedprintable.Writer
//    func (w *Writer) Close() error: 关闭 Writer，将任何未写入的数据刷新到底层 io.Writer，但不会关闭底层 io.Writer。
//    func (w *Writer) Write(p []byte) (n int, err error): 使用quotedprintable编码对 p 进行编码，并将其写入底层 io.Writer。它将行长度限制为 76 个字符。在 Writer 关闭之前，不一定会刷新编码的字节。
func TestTypeWriter(t *testing.T) {
	_ = quotedprintable.Writer{}
	// type Writer struct {
	//    // Binary mode treats the writer's input as pure binary and processes end of
	//    // line bytes as binary data.
	//    Binary bool
	//    // contains filtered or unexported fields
	// }
}

func TestReader(t *testing.T) {
	for _, s := range []string{
		`=48=65=6C=6C=6F=2C=20=47=6F=70=68=65=72=73=21`,
		`invalid escape: <b style="font-size: 200%">hello</b>`,
		"Hello, Gophers! This symbol will be unescaped: =3D and this will be written in =\r\none line.",
	} {
		b, err := io.ReadAll(quotedprintable.NewReader(strings.NewReader(s)))
		fmt.Printf("%s %v\n", b, err)
	}
}

func TestWriter(t *testing.T) {
	w := quotedprintable.NewWriter(os.Stdout)
	w.Write([]byte("These symbols will be escaped: = \t"))
	w.Close()
}
