package textproto

// net/textproto: 实现了HTTP、NNTP和SMTP的基于问题的请求、响应协议的通用支持

// 函数:
// func CanonicalMIMEHeaderKey(s string) string // 返回MIME头部键s的规范格式
// func TrimBytes(b []byte) []byte              // 返回b去除前后ASCII空格后内容
// func TrimString(s string) string             // 返回b去除前后ASCII空格后内容

// 类型:
// 1.Conn: 表示文本网络连接协议
// type Conn struct {
//    Reader
//    Writer
//    Pipeline
//    // contains filtered or unexported fields
// }
// func Dial(network, addr string) (*Conn, error)                              // 使用net.Dial连接网络到指定地址，返回新的Conn连接
// func NewConn(conn io.ReadWriteCloser) *Conn                                 // 从io.ReadWriteCloser返回一新的Conn
// func (c *Conn) Close() error                                                // 关闭连接
// func (c *Conn) Cmd(format string, args ...interface{}) (id uint, err error) // 一个遍历的向管道发送命令并等待响应的方法

// 2.Error: 表示来自服务器的数字错误响应
// type Error struct {
//    Code int
//    Msg  string
// }
// func (e *Error) Error() string

// 3.MIMEHeader: MIME峰峰的头映射，字符串键对应值为字符串切片
// type MIMEHeader map[string][]string
// func (h MIMEHeader) Add(key, value string)      // 追加键值到对应的键
// func (h MIMEHeader) Del(key string)             // 删除键
// func (h MIMEHeader) Get(key string) string      // 获取键的第一个值
// func (h MIMEHeader) Set(key, value string)      // 替换现有的键值
// func (h MIMEHeader) Values(key string) []string // 获取键关联的所有值

// 4.Pipeline: 管理管道的有序请求/响应队列
// type Pipeline struct {
//    // contains filtered or unexported fields
// }
// func (p *Pipeline) EndRequest(id uint)    // 通知p已发送给定id的请求 (如果是服务器表示收到)
// func (p *Pipeline) EndResponse(id uint)   // 通知p已经收到给定id的响应 (如果是服务器表示发送)
// func (p *Pipeline) Next() uint            // 返回请求/响应对的下一个id
// func (p *Pipeline) StartRequest(id uint)  // 阻塞直到给定id的请求的发送时间到达 (如果是服务器表示接收)
// func (p *Pipeline) StartResponse(id uint) // 阻塞自导给定id的请求的接收时间到达 (如果是服务器表示发送)

// 5.ProtocolError: 描述违反协议错误,例如无效响应或挂断连接
// type ProtocolError string
// func (p ProtocolError) Error() string

// 6.Reader: 实现了从一文本协议网络连接读取请求或响应的遍历方法
// type Reader struct {
//    R *bufio.Reader
//    // contains filtered or unexported fields
// }
// func NewReader(r *bufio.Reader) *Reader // 从bufio.Reader来创建Reader
// func (r *Reader) DotReader() io.Reader // 从r返回一dot的reader
// func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error) // 读取响应码行
// func (r *Reader) ReadContinuedLine() (string, error) // 从r读取接下来的行,刅空行结束
// func (r *Reader) ReadContinuedLineBytes() ([]byte, error) // 与ReadContinuedLine类似,区别在于返回字节切片
// func (r *Reader) ReadDotBytes() ([]byte, error) // 读取点编码并返回解码的数据
// func (r *Reader) ReadDotLines() ([]string, error) // 读取点编码并返回包含解码线的切片,其中以\r\n或\n结尾
// func (r *Reader) ReadLine() (string, error) // 从r读取一行,去除末尾的\r\n或\n
// func (r *Reader) ReadLineBytes() ([]byte, error) // 与ReadLine类似,区别在于返回字节切片
// func (r *Reader) ReadMIMEHeader() (MIMEHeader, error) // 从r读取MIME风格头并返回MIMEHeader和错误
// func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error) // 读取表单的多行响应

// 7.Writer: 实现了为文本协议网络连接写入请求或响应的遍历方法
// type Writer struct {
//    W *bufio.Writer
//    // contains filtered or unexported fields
// }
// func NewWriter(w *bufio.Writer) *Writer                               // 从bufio.Writer创建一Writer
// func (w *Writer) DotWriter() io.WriteCloser                           // 返回一io.WriteCloser来用于向w写入点编码内容
// func (w *Writer) PrintfLine(format string, args ...interface{}) error // 写入以\r\n结尾的格式化输出
