package httputil

import (
	"net/http/httputil"
	"testing"
)

// net/http/httputil: 提供了HTTP实用方法

// 变量
func TestVar(t *testing.T) {
	_ = httputil.ErrLineTooLong // 读取行数据太长
}

// 函数
// func DumpRequest(req *http.Request, body bool) ([]byte, error) // 在HTTP1.x表示返回给定请求 (应只用于服务器表示客户端请求)
// func DumpRequestOut(req *http.Request, body bool) ([]byte, error) // 与DumpRequest类似,但用于传出客户端请求
// func DumpResponse(resp *http.Response, body bool) ([]byte, error) // 与DyumpRequest类似,但用于转存响应
// func NewChunkedReader(r io.Reader) io.Reader // 返回一新的chunkedReader
// func NewChunkedWriter(w io.Writer) io.WriteCloser // 返回一新的chunkedWriter,在写入w之前翻译写入的HTTP chuned格式

// 类型
// 1.BufferPool: 一接口,英语获取和返回临时字节切片给io.CopyBuffer使用
// type BufferPool interface {
//     Get() []byte
//     Put([]byte)
// }

// 2.ClientConn: 早起HTTP实现模块 (已废弃)
// type ClientConn struct {
//    // contains filtered or unexported fields
// }
// func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn
// func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn
// func (cc *ClientConn) Close() error
// func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)
// func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)
// func (cc *ClientConn) Pending() int
// func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)
// func (cc *ClientConn) Write(req *http.Request) error

// 3.ReverseProxy: 一HTTP handler,将传入请求发送给另一个服务器,代理响应返回给客户端
// type ReverseProxy
// func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy // 基于路由URL创建一ReverseProxy
// func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)

// 4.ServerConn: 早起HTTP实现模块 (已废弃)
// type ServerConn struct {
//    // contains filtered or unexported fields
// }
// func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn
// func (sc *ServerConn) Close() error
// func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)
// func (sc *ServerConn) Pending() int
// func (sc *ServerConn) Read() (*http.Request, error)
// func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error
