package httptrace

// net/http/httptrace: 提供了追踪HTTP客户端请求机制

// 函数
// func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context // 传入一ctx和trace返回一新的上下文

// 类型
// 1.ClientTrace: 运行在输出HTTP请的不同阶段的钩子集合
// type ClientTrace
// func ContextClientTrace(ctx context.Context) *ClientTrace // 基于ctx创建一ClientTrace

// 2.DNSDoneInfo: DNS查询结果信息
// type DNSDoneInfo struct {
//     Addrs []net.IPAddr
//     Err error
//     Coalesced bool
// }

// 3.DBSStartInfo: DNS请求信息
// type DNSStartInfo struct {
//    Host string
// }

// 4.GotConnInfo: 包含获取连接的信息 (ClientTrace.GotConn函数参数)
// type GotConnInfo struct {
//     Conn net.Conn
//     Reused bool
//     WasIdle bool
//     IdleTime time.Duration
// }

// 5.WroteRequestInfo: 提供给WroteRequest钩子的信息
// type WroteRequestInfo struct {
//     Err error
// }
