package syslog

// log/syslog
// 为系统日志服务提供了一个简单的接口。它可以使用 UNIX 域套接字、UDP 或 TCP 向系统日志守护进程发送消息。
// syslog 包被冻结并且不接受新功能。一些外部包提供更多功能可以访问：https://godoc.org/?q=syslog

// func NewLogger(p Priority, logFlag int) (*log.Logger, error)
// type Priority
// type Writer
//    func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)
//    func New(priority Priority, tag string) (*Writer, error)
//    func (w *Writer) Alert(m string) error
//    func (w *Writer) Close() error
//    func (w *Writer) Crit(m string) error
//    func (w *Writer) Debug(m string) error
//    func (w *Writer) Emerg(m string) error
//    func (w *Writer) Err(m string) error
//    func (w *Writer) Info(m string) error
//    func (w *Writer) Notice(m string) error
//    func (w *Writer) Warning(m string) error
//    func (w *Writer) Write(b []byte) (int, error)
