package httptest

// net/http/httptest: 提供HTTP测试的方法

// 常量
// const DefaultRemoteAddr = "1.2.3.4"

// 函数
// func NewRequest(method, target string, body io.Reader) *http.Request

// 类型
// type ResponseRecorder
//    func NewRecorder() *ResponseRecorder
//    func (rw *ResponseRecorder) Flush()
//    func (rw *ResponseRecorder) Header() http.Header
//    func (rw *ResponseRecorder) Result() *http.Response
//    func (rw *ResponseRecorder) Write(buf []byte) (int, error)
//    func (rw *ResponseRecorder) WriteHeader(code int)
//    func (rw *ResponseRecorder) WriteString(str string) (int, error)
// type Server
//    func NewServer(handler http.Handler) *Server
//    func NewTLSServer(handler http.Handler) *Server
//    func NewUnstartedServer(handler http.Handler) *Server
//    func (s *Server) Certificate() *x509.Certificate
//    func (s *Server) Client() *http.Client
//    func (s *Server) Close()
//    func (s *Server) CloseClientConnections()
//    func (s *Server) Start()
//    func (s *Server) StartTLS()
