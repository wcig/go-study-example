package httptest

// net/http/httptest: 提供HTTP测试的方法

// 常量
// const DefaultRemoteAddr = "1.2.3.4"

// 函数
// func NewRequest(method, target string, body io.Reader) *http.Request // 返回一新传入的服务器请求,适用于http.Handler进行测试

// 类型
// 1.ResponseRecorder: 实现了http.ResponseWriter接口,可用于记录
// type ResponseRecorder struct {
//     Code int
//     HeaderMap http.Header
//     Body *bytes.Buffer
//     Flushed bool
// }
// func NewRecorder() *ResponseRecorder // 创建一新ResponseRecorder
// func (rw *ResponseRecorder) Flush() // 实现http.Flushed,用于测试Flush是否被调用
// func (rw *ResponseRecorder) Header() http.Header // 实现http.ResponseWriter,返回响应头
// func (rw *ResponseRecorder) Result() *http.Response // 返回响应
// func (rw *ResponseRecorder) Write(buf []byte) (int, error) // 实现http.ResponseWriter接口,将buf的数据写入到rw.Body如果非空
// func (rw *ResponseRecorder) WriteHeader(code int) // 实现http.ResponseWriter
// func (rw *ResponseRecorder) WriteString(str string) (int, error) // 实现io.StringWriter,如果str非空则将其数据写入rw.Body

// 2.Server: HTTP服务器,用于端到端的HTTP测试
// type Server struct {
//     URL string
//     Listener net.Listener
//     EnableHTTP2 bool
//     TLS *tls.Config
//     Config *http.Server
// }
// func NewServer(handler http.Handler) *Server // 基于handler创建一新Server
// func NewTLSServer(handler http.Handler) *Server // 与NewServer类似,区别在于支持TLS
// func NewUnstartedServer(handler http.Handler) *Server // 创建一新的Server,默认不启动 (后续可修改配置再启动)
// func (s *Server) Certificate() *x509.Certificate // 返回服务器使用的证书 (服务器没有使用TLS则返回nil)
// func (s *Server) Client() *http.Client  // 返回配置与服务器请求的HTTP客户端
// func (s *Server) Close() // 关闭服务器,阻塞直到所有请求完成
// func (s *Server) CloseClientConnections() // 关闭所有到测试Server打开的HTTP连接
// func (s *Server) Start() // 打开从NewUnstartedServer创建的Server
// func (s *Server) StartTLS() // 与Start类似,区别在于支持TLS
