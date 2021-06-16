package multipart

import (
	"mime/multipart"
	"os"
	"testing"
)

// mime/multipart：实现了的MIME multipart解析，参见RFC 2046。该实现适用于HTTP（RFC 2388）和常见浏览器生成的multipart主体。

// 错误
func TestErr(t *testing.T) {
	_ = multipart.ErrMessageTooLarge
}

// 类型：multipart.File 多部分消息的文件部分访问接口（只能读取不能写入）
func TestTypeFile(t *testing.T) {
	_ = multipart.File(os.Stdin)
	// type File interface {
	//    io.Reader
	//    io.ReaderAt
	//    io.Seeker
	//    io.Closer
	// }
}

// 类型：multipart.FileHeader 描述multipart请求的文件部分
// func (fh *FileHeader) Open() (File, error)：打开并返回fh关联的文件
func TestTypeFileHeader(t *testing.T) {
	_ = multipart.FileHeader{}
	// type FileHeader struct {
	//    Filename string
	//    Header   textproto.MIMEHeader
	//    Size     int64 // Go 1.9
	//    // contains filtered or unexported fields
	// }
}

// 类型：multipart.Form 解析的multipart表单，键值保存在Value中，文件保存在File中。
// func (f *Form) RemoveAll() error：移除所有f关联的临时文件
func TestTypeForm(t *testing.T) {
	_ = multipart.Form{}
	// type Form struct {
	//    Value map[string][]string
	//    File  map[string][]*FileHeader
	// }
}

// 类型：multipart.Part 一个multipart主体的一部分
// func (p *Part) Close() error：io关闭
// func (p *Part) FileName() string：返回Content-Disposition头的文件名参数
// func (p *Part) FormName() string：如果 p 具有“form-data”类型的 Content-Disposition，则 FormName 返回 name 参数。否则它返回空字符串。
// func (p *Part) Read(d []byte) (n int, err error)：读取一part的内容，在其头之后和下一part之前。
func TestTypePart(t *testing.T) {
	_ = multipart.Part{}
	// type Part struct {
	//    // The headers of the body, if any, with the keys canonicalized
	//    // in the same fashion that the Go http.Request headers are.
	//    // For example, "foo-bar" changes case to "Foo-Bar"
	//    Header textproto.MIMEHeader
	//    // contains filtered or unexported fields
	// }
}

// 类型：multipart.Reader：MIME的multipart主体的多部分迭代器，不支持seek。
// func NewReader(r io.Reader, boundary string) *Reader：给定MIME边界从r创建multipart的Reader（边界boundary参数值通常从消息的"Content-Type"头获取，使用mime.ParseMediaType方法可以解析这样的头）
// func (r *Reader) NextPart() (*Part, error)：返回multipart的下一part或错误，当没有更多part时返回错误io.EOF。
// func (r *Reader) NextRawPart() (*Part, error)：同NextPart基本一样，区别在于对于"Content-Transfer-Encoding: quoted-printable"没有特殊处理。
// func (r *Reader) ReadForm(maxMemory int64) (*Form, error)：解析整个multipart消息。
func TestTypeReader(t *testing.T) {
	_ = multipart.Reader{}
	// type Reader struct {
	//    // contains filtered or unexported fields
	// }
}

// 类型：multipart.Writer 生成multipart消息
// func NewWriter(w io.Writer) *Writer：基于w创建一新的具有随机边界的multipartWriter。
// func (w *Writer) Boundary() string：返回Writer的边界。
// func (w *Writer) Close() error：完成multipart消息，并将尾随的边界结束符写入输出。
// func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)：基于给定字段名调用有头的CreatePart。
// func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)：基于字段名和文件名创建一新的form-data，是CreatePart的便捷包装。
// func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)：基于给定的头创建一新的multipart部分。
// func (w *Writer) FormDataContentType() string：返回具有此Writer便捷的HTTP multipart/form-data的Content-TYpe。
// func (w *Writer) SetBoundary(boundary string) error：使用显示的值覆盖Writer默认随机生成的便捷分隔符。
// func (w *Writer) WriteField(fieldname, value string) error：调用CreateFormFiled，然后写入给定值。
func TestTypeWriter(t *testing.T) {
	_ = multipart.Writer{}
	// type Writer struct {
	//    // contains filtered or unexported fields
	// }
}
