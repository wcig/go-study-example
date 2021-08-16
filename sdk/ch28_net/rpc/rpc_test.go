package rpc

// net/rpc: rpc系统，支持客户端和服务端（只适用于Go，因为内部使用encoding/gob编解码）

// 要求：
// 1.方法类型是导出类型
// 2.方法是可导出的
// 3.方法必须有2个参数，且都为可导出的或Go内置类型
// 4.方法的第2个参数是指针
// 5.方法返回值为为error了下

// 方法模板
// func (t *T) MethodName(argType T1, replyType *T2) error
// T为导出类型，T1，T2为可导出或内置类型，T1位客户端请求参数，T2位服务端返回参数。

// 常量
// const (
//    // Defaults used by HandleHTTP
//    DefaultRPCPath   = "/_goRPC_"
//    DefaultDebugPath = "/debug/rpc"
// )

// 变量
// var DefaultServer = NewServer()
// var ErrShutdown = errors.New("connection is shut down")

// 函数
// func Accept(lis net.Listener)                          // 接受监听器的连接，并为每个传入连接请求到DefaultServer，接收块，调用者一般在go语句中调用它
// func HandleHTTP()                                      // 在DefaultRPCPath上为DefaultServer注册一RPC消息的HTTP handler，在DefaultDebugPath注册一调试handler，一般在go语句中使用，后续需调用http.Serve()
// func Register(rcvr interface{}) error                  // 在DefaultServer中发布接收者的方法
// func RegisterName(name string, rcvr interface{}) error // 与Register类似，但使用提供的类型名称代替接收者的具体类型
// func ServeCodec(codec ServerCodec)                     // 与ServeConn类似，但是使用指定的解码器codec来解码请求和编码响应
// func ServeConn(conn io.ReadWriteCloser)                // 在单个连接上运行DefaultServer
// func ServeRequest(codec ServerCodec) error             // 与ServeConn类似，但是是同步处理单个请求，万能充后不会关闭编解码器

// 类型
// 1.Call: 活动的RPC
// type Call struct {
//    ServiceMethod string      // The name of the service and method to call.
//    Args          interface{} // The argument to the function (*struct).
//    Reply         interface{} // The reply from the function (*struct).
//    Error         error       // After completion, the error status.
//    Done          chan *Call  // Receives *Call when Go is complete.
// }

// 2.Client: 表示一个RPC客户端
// type Client struct {
//    // contains filtered or unexported fields
// }
// func Dial(network, address string) (*Client, error) // 以特定的网络地址连接RPC服务器
// func DialHTTP(network, address string) (*Client, error) // 以特定的网络地址连接以HTTP RPC服务器，监听在默认HTTP RPC路径
// func DialHTTPPath(network, address, path string) (*Client, error) // 以特定的网络地址和路径连接HTTP RPC服务器
// func NewClient(conn io.ReadWriteCloser) *Client // 返回一新的客户端
// func NewClientWithCodec(codec ClientCodec) *Client // 与NewClient类似，但是以指定的编解码器
// func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error // 同步调用指定方法
// func (client *Client) Close() error // 调用底层编解码器Close方法，如果连接已关闭，ErrShutdown返回
// func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call // 异步调用方法，返回表示调用的Call来进行信号传出

// 3.ClientCodec: RPC客户端的编解码器
// type ClientCodec interface {
//    WriteRequest(*Request, interface{}) error
//    ReadResponseHeader(*Response) error
//    ReadResponseBody(interface{}) error
//
//    Close() error
// }

// 4.Request: 每一个RPC调用之前的写入header，内部使用，在这里用于方便调试。
// type Request struct {
//    ServiceMethod string // format: "Service.Method"
//    Seq           uint64 // sequence number chosen by client
//    // contains filtered or unexported fields
// }

// 5.Response: 每一个RPC返回的响应前写入header
// type Response

// 6.Server: RPC服务器
// type Server struct {
//    // contains filtered or unexported fields
// }
// func NewServer() *Server // 返回一新的Server
// func (server *Server) Accept(lis net.Listener) // 在监听器上接受请求，并处理每一个出阿奴的连接
// func (server *Server) HandleHTTP(rpcPath, debugPath string) // 在rpcPath上注册一RPC消息的HTTP handler，在debugPath注册一调试handler，一般在go语句中使用，后续需调用http.Serve()
// func (server *Server) Register(rcvr interface{}) error // 发布接收者的方法
// func (server *Server) RegisterName(name string, rcvr interface{}) error // 与Register类似，但使用提供的类型名称代替接收者的具体类型
// func (server *Server) ServeCodec(codec ServerCodec) // 与ServeConn类似，但是使用指定的解码器codec来解码请求和编码响应
// func (server *Server) ServeConn(conn io.ReadWriteCloser) // 在单个连接上运行Server
// func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) // 实现http.Handler来回答RPC请求
// func (server *Server) ServeRequest(codec ServerCodec) error // 与ServeCodec类似，但是是同步来处理单个请

// 7.ServerCodec RPC服务端的编解码器
// type ServerCodec interface {
//    ReadRequestHeader(*Request) error
//    ReadRequestBody(interface{}) error
//    WriteResponse(*Response, interface{}) error
//
//    // Close can be called multiple times and must be idempotent.
//    Close() error
// }

// 8.ServerError: 表示从RPC连接的远程侧返回的错误
// type ServerError string
// func (e ServerError) Error() string
