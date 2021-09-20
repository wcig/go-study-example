package http

import (
	"net/http"
	"testing"
)

// net/http: 提供了HTTP client和server的实现

// 常量
// 1.HTTP方法
// const (
// 	MethodGet     = "GET"
// 	MethodHead    = "HEAD"
// 	MethodPost    = "POST"
// 	MethodPut     = "PUT"
// 	MethodPatch   = "PATCH" // RFC 5789
// 	MethodDelete  = "DELETE"
// 	MethodConnect = "CONNECT"
// 	MethodOptions = "OPTIONS"
// 	MethodTrace   = "TRACE"
// )

// 2.HTTP状态码
// const (
// 	StatusContinue           = 100 // RFC 7231, 6.2.1
// 	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
// 	StatusProcessing         = 102 // RFC 2518, 10.1
// 	StatusEarlyHints         = 103 // RFC 8297
//
// 	StatusOK                   = 200 // RFC 7231, 6.3.1
// 	StatusCreated              = 201 // RFC 7231, 6.3.2
// 	StatusAccepted             = 202 // RFC 7231, 6.3.3
// 	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
// 	StatusNoContent            = 204 // RFC 7231, 6.3.5
// 	StatusResetContent         = 205 // RFC 7231, 6.3.6
// 	StatusPartialContent       = 206 // RFC 7233, 4.1
// 	StatusMultiStatus          = 207 // RFC 4918, 11.1
// 	StatusAlreadyReported      = 208 // RFC 5842, 7.1
// 	StatusIMUsed               = 226 // RFC 3229, 10.4.1
//
// 	StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
// 	StatusMovedPermanently = 301 // RFC 7231, 6.4.2
// 	StatusFound            = 302 // RFC 7231, 6.4.3
// 	StatusSeeOther         = 303 // RFC 7231, 6.4.4
// 	StatusNotModified      = 304 // RFC 7232, 4.1
// 	StatusUseProxy         = 305 // RFC 7231, 6.4.5
//
// 	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
// 	StatusPermanentRedirect = 308 // RFC 7538, 3
//
// 	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
// 	StatusUnauthorized                 = 401 // RFC 7235, 3.1
// 	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
// 	StatusForbidden                    = 403 // RFC 7231, 6.5.3
// 	StatusNotFound                     = 404 // RFC 7231, 6.5.4
// 	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
// 	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
// 	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
// 	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
// 	StatusConflict                     = 409 // RFC 7231, 6.5.8
// 	StatusGone                         = 410 // RFC 7231, 6.5.9
// 	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
// 	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
// 	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
// 	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
// 	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
// 	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
// 	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
// 	StatusTeapot                       = 418 // RFC 7168, 2.3.3
// 	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
// 	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
// 	StatusLocked                       = 423 // RFC 4918, 11.3
// 	StatusFailedDependency             = 424 // RFC 4918, 11.4
// 	StatusTooEarly                     = 425 // RFC 8470, 5.2.
// 	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
// 	StatusPreconditionRequired         = 428 // RFC 6585, 3
// 	StatusTooManyRequests              = 429 // RFC 6585, 4
// 	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
// 	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3
//
// 	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
// 	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
// 	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
// 	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
// 	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
// 	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
// 	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
// 	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
// 	StatusLoopDetected                  = 508 // RFC 5842, 7.2
// 	StatusNotExtended                   = 510 // RFC 2774, 7
// 	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
// )

// 3.其他
// const DefaultMaxHeaderBytes = 1 << 20              // 1 MB // HTTP请求头最大允许大小 (可通过设置Server.MaxHeaderBytes覆盖)
// const DefaultMaxIdleConnsPerHost = 2               // Transport的MaxIdleConnsPerHost
// const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT" // HTTP头生成时间的格式

// 变量
func TestVar(t *testing.T) {
	_ = http.ErrNotSupported
	_ = http.ErrMissingBoundary
	_ = http.ErrNotMultipart

	_ = http.ErrBodyNotAllowed
	_ = http.ErrHijacked
	_ = http.ErrContentLength

	_ = http.ServerContextKey
	_ = http.LocalAddrContextKey

	_ = http.DefaultClient
	_ = http.DefaultServeMux

	_ = http.ErrAbortHandler
	_ = http.ErrBodyNotAllowed
	_ = http.ErrBodyReadAfterClose
	_ = http.ErrHandlerTimeout
	_ = http.ErrLineTooLong
	_ = http.ErrMissingFile
	_ = http.ErrNoCookie
	_ = http.ErrNoLocation
	_ = http.ErrSkipAltProtocol
	_ = http.ErrUseLastResponse
	_ = http.NoBody
}

// 方法
// func CanonicalHeaderKey(s string) string // 返回头部键s的规范格式
// func DetectContentType(data []byte) string // 实现解析算法解析数据的Content-Type,始终返回有效的MIME类型,如果无法确定则返回 "application/octet-stream"
// func Error(w ResponseWriter, error string, code int) // 以指定错误消息和HTTP状态码响应请求,其不会结束请求,调用者需确保不再写入数据到w,错误消息必须为纯文本
// func Handle(pattern string, handler Handler) // 在DefaultServeMux一指定模式注册一handler
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) // 在DefaultServeMux以指定魔术注册一handler函数
// func ListenAndServe(addr string, handler Handler) error // 监听指定的TCP地址addr并处理传入的请求,接收的连接被配置为TCP保持连接
// func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error // 与ListenAndServe类似,区别在支持HTTPS连接
// func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser // 与LimitReader类似但其限制输入请求体的大小,与io.LimitReader不同这里返回的是io.ReadCloser,对于超出限制的读取返回非EOF错误并在调用其Close方法时关闭底层reader (主要用于防止刻度恶意发送大请求)
// func NotFound(w ResponseWriter, r *Request) // 返回HTTP 404 not found错误
// func ParseHTTPVersion(vers string) (major, minor int, ok bool) // 解析HTTP版本 (例如"HTTP/1.0"返回1,0,true)
// func ParseTime(text string) (t time.Time, err error) // 解析时间头
// func ProxyFromEnvironment(req *Request) (*url.URL, error) // 返回给定请求的代理URL
// func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error) // 返回一始终相同的URL的代理函数用于传输
// func Redirect(w ResponseWriter, r *Request, url string, code int) // 以指定code重定向请求到url来响应请求
// func Serve(l net.Listener, handler Handler) error // 在监听器l接收传入的请求并以handler处理请求,每个请求会创建一新的go routine来处理
// func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker) // 以指定的ReadSeeker的内容回复请求
// func ServeFile(w ResponseWriter, r *Request, name string) // 以指定文件或目录来响应请求
// func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error // 与Serve类似,区别在于支持HTTPS请求
// func SetCookie(w ResponseWriter, cookie *Cookie) // 为响应设置指定cookie
// func StatusText(code int) string // 返回HTTP状态码,未知则返回空字符串

// 类型
// 1.Client: HTTP客户端
// type Client
// func (c *Client) CloseIdleConnections() // 关闭空闲连接
// func (c *Client) Do(req *Request) (*Response, error) // 发送HTTP请求不难过返回HTTP响应和错误 (非2xx状态码不返回错误)
// func (c *Client) Get(url string) (resp *Response, err error) // 发送HTTP get请求
// func (c *Client) Head(url string) (resp *Response, err error) // 发送HEAD到指定url
// func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) // 发送HTTP post请求
// func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) // 发送post 表单请求,请求头Content-Type为application/x-www-form-urlencoded

// 2.CloseNotifier: 接口由ResponseWriters实现,允许检测底层连接何时小时 (废弃)
// type CloseNotifier

// 3.ConnState: 表示客户端连接到服务器的状态 (被用于可选的Server.ConnState hook)
// type ConnState int
// func (c ConnState) String() string
// 枚举常量
// const (
// 	StateNew ConnState = iota
// 	StateActive
// 	StateIdle
// 	StateHijacked
// 	StateClosed
// )

// 4.Cookie: 表示HTTP的cookie
// type Cookie struct {
//    Name  string
//    Value string
//    Path       string    // optional
//    Domain     string    // optional
//    Expires    time.Time // optional
//    RawExpires string    // for reading cookies only
//    MaxAge   int
//    Secure   bool
//    HttpOnly bool
//    SameSite SameSite // Go 1.11
//    Raw      string
//    Unparsed []string // Raw text of unparsed attribute-value pairs
// }
// func (c *Cookie) String() string

// 5.CookieJar: 管理HTTP请求cookie的存储和使用
// type CookieJar interface {
//    SetCookies(u *url.URL, cookies []*Cookie)
//    Cookies(u *url.URL) []*Cookie
// }

// 6.Dir: 实现了FileSystem,用于指定目录树的本机文件系统
// type Dir string
// func (d Dir) Open(name string) (File, error) // 使用os.Open实现FileSystem

// 7.File: 具有FileSystem的Open方法,并且由FileServer实现提供服务
// type File interface {
//    io.Closer
//    io.Reader
//    io.Seeker
//    Readdir(count int) ([]fs.FileInfo, error)
//    Stat() (fs.FileInfo, error)
// }

// 8.FileSystem: 实现命名文件集合的访问
// type FileSystem interface {
//     Open(name string) (File, error)
// }
// func FS(fsys fs.FS) FileSystem // 转换fsys为FileSystem实现

// 9.Flusher: 该接口由ResponseWriters实现,允许HTTP handler将缓冲数据刷新到客户端
// type Flusher interface {
//     Flush()
// }

// 10.Handler: 响应HTTP请求
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
// func FileServer(root FileSystem) Handler // 返回一处理HTTP请求的handler,该处理程序是以root为根的文件系统
// func NotFoundHandler() Handler // 返回一响应404 page not found的handler
// func RedirectHandler(url string, code int) Handler // 返回一以指定code指向url的重定向handler
// func StripPrefix(prefix string, h Handler) Handler // 返回一处理HTTP请求的handler,该handler从请求URL的Path删除指定前缀然后调用handler h
// func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler // 返回一限制时间的handler

// 11.HandlerFunc: 允许普通函数作为HTTP handler的适配器
// type HandlerFunc ServeHTTP(ResponseWriter, *Request)
//    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

// 12.Header: 表示HTTP头的键值对
// type Header
// func (h Header) Add(key, value string)
// func (h Header) Clone() Header
// func (h Header) Del(key string)
// func (h Header) Get(key string) string
// func (h Header) Set(key, value string)
// func (h Header) Values(key string) []string
// func (h Header) Write(w io.Writer) error
// func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error

// 13.Hijacker: 接口由ResponseWriters实现,允许HTTP handler接管连接
// type Hijacker interface {
//     Hijack(net.Conn, *bufio.ReadWriter, error)
// }

// 14.ProtocolError: 表示HTTP协议错误 (弃用)
// type ProtocolError struct {
//     ErrorString string
// }
// func (pe *ProtocolError) Error() string

// 15.PushOptions: 描述Pusher.Push的选项
// type PushOptions struct {
//    Method string
//    Header Header
// }

// 16.Pusher: 接口由ResponseWriters实现,支持HTTP/2服务器推送
// type Pusher interface {
//    Push(target string, opts *PushOptions)
// }

// 17.Request: 表示HTTP请求
// type Request
// func NewRequest(method, url string, body io.Reader) (*Request, error) // 使用background context包装NewRequestWithContext
// func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) // 以指定ctx, 方法, url和可选body创建Request
// func ReadRequest(b *bufio.Reader) (*Request, error) // 从b读取并解析输入的请求
// func (r *Request) AddCookie(c *Cookie) // 添加cookie
// func (r *Request) BasicAuth() (username, password string, ok bool) // 返回请求Authorization头的用户名和密码 (如果请求使用HTTP基本身份认证)
// func (r *Request) Clone(ctx context.Context) *Request // 返回r的深拷贝并替换上线文为ctx (ctx必须为非空)
// func (r *Request) Context() context.Context // 返回请求的上线文,修改上下文使用WithContext方法
// func (r *Request) Cookie(name string) (*Cookie, error) // 返回请求中指定名称name的cookie,没有则返回ErrNoCookie
// func (r *Request) Cookies() []*Cookie // 解析并返回发送请求的Cookie
// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) // 返回表单key的第一个文件
// func (r *Request) FormValue(key string) string // 返回key对应的第一个值
// func (r *Request) MultipartReader() (*multipart.Reader, error) // 如果是multipart/form-data或mutipart/mixed的POST请求返回一MIME的多部分reader,否则返回nil和错误
// func (r *Request) ParseForm() error // 填充r.Form和r.PostForm
// func (r *Request) ParseMultipartForm(maxMemory int64) error // 解析请求体为multipart/form-data,整个请求体的解析最多maxMemory字节数据被存储在内存中,其他存储在磁盘的临时文件中
// func (r *Request) PostFormValue(key string) string // 返回key的第一个值
// func (r *Request) ProtoAtLeast(major, minor int) bool // 报告请求使用的HTTP协议是否至少是major.minor版本
// func (r *Request) Referer() string // 返回发送请求的引用URL
// func (r *Request) SetBasicAuth(username, password string) // 设置请求的HTTP基本身份认证
// func (r *Request) UserAgent() string // 返回发送请求客户端的User-Agent
// func (r *Request) WithContext(ctx context.Context) *Request // 返回替换上下文为ctx后r的浅拷贝,ctx不能为空
// func (r *Request) Write(w io.Writer) error // 写入HTTP/1.1请求
// func (r *Request) WriteProxy(w io.Writer) error // 与Write类似,但是以HTTP代理期望形式写入请求

// 18.Response: HTTP请求的响应
// type Response
// func Get(url string) (resp *Response, err error) // 向url发送GET请求并返回响应和错误
// func Head(url string) (resp *Response, err error) // 向url发送HEAD请求并返回响应和错误
// func Post(url, contentType string, body io.Reader) (resp *Response, err error) // 向url发送POST请求并返回响应和错误
// func PostForm(url string, data url.Values) (resp *Response, err error) // 向url发送POST表单请求并返回响应和错误 (Content-Type头被设置为application/x-www-form-urlencoded)
// func ReadResponse(r *bufio.Reader, req *Request) (*Response, error) // 从r读取并返回响应
// func (r *Response) Cookies() []*Cookie // 解析响应并返回Set-Cookie头设置的cookie
// func (r *Response) Location() (*url.URL, error) // 返回响应的Location头
// func (r *Response) ProtoAtLeast(major, minor int) bool // 报告响应使用的HTTP协议是否至少major.minor版本
// func (r *Response) Write(w io.Writer) error // 以HTTP/1.x写入w到响应

// 19.ResponseWriter: 接口被HTTP handler用来构造HTTP响应
// type ResponseWriter interface {
//     Header() Header
//     Write([]byte) (int, error)
//     WriteHeader(statusCode int)
// }

// 20.RoundTripper: 接口,表示执行单个HTTP事务的能力,获取给定请求的响应
// type RoundTripper interface {
//    RoundTrip(*Request) (*Response, error)
// }
// func NewFileTransport(fs FileSystem) RoundTripper // 以提供的FileSystem来返回一新的RoundTriper

// 21.SameSite: 允许服务器定义cookie属性,来使得浏览器无法发送跨站点请求
// type SameSite int
// 常量
// const (
//    SameSiteDefaultMode SameSite = iota + 1
//    SameSiteLaxMode
//    SameSiteStrictMode
//    SameSiteNoneMode
// )

// 22.ServeMux: HTTP请求多路复用器.将每个传入请求的URL与注册魔术列表匹配,并调用最接近URL匹配的handler
// type ServeMux
// func NewServeMux() *ServeMux // 新建以ServeMutex
// func (mux *ServeMux) Handle(pattern string, handler Handler) // 为指定模式注册handler (如果模式已存在一handler则panic)
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) // 为指定模式注册一HandleFunc函数
// func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) // 返回指定请求的handler
// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) // 将请求调度到最匹配请求URL的handler

// 23.Server: 定义运行HTTP服务器的参数 (零值是有效配置)
// type Server
// func (srv *Server) Close() error // 立刻关闭所有活动的net.Listeners和任意状态为StateNew,StateActive,StateIdle的连接 (优雅关机请使用Shutdown)
// func (srv *Server) ListenAndServe() error // 开始TCP网络监听
// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error // 与ListenAndServe类似,区别在于支持TLS
// func (srv *Server) RegisterOnShutdown(f func()) // 注册一Shutdown时的函数调用
// func (srv *Server) Serve(l net.Listener) error // 使用l开始监听
// func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error // 与Serve类似,区别在于支持TLS
// func (srv *Server) SetKeepAlivesEnabled(v bool) // 控制HTTP keep-alices是否开启 (默认开启)
// func (srv *Server) Shutdown(ctx context.Context) error // 优雅关机

// 24.Transport: 支持HTTP,HTTPS和HTTP代理的实现RoundTripper接口
// type Transport
// func (t *Transport) CancelRequest(req *Request) // 通过关闭连接来取消传输中的请求 (只有在RoundTrip返回才可被调用)
// func (t *Transport) Clone() *Transport // 返回一深拷贝副本
// func (t *Transport) CloseIdleConnections() // 关闭空闲连接
// func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper) // 注册协议
// func (t *Transport) RoundTrip(req *Request) (*Response, error) // 实现RoundTripper接口
