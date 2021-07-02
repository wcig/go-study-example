package smtp

// net/smtp: 实现了RFC 5231定义的邮件传输协议
// 发送带附件的邮件可使用第三方库："https://github.com/go-gomail/gomail"
// 或参考："http://www.361way.com/golang-email-attachment/5856.html"

// 函数
// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error // 发送邮箱，addr：邮箱服务器，a：身份认证，from：发送邮件邮箱地址，to：接收邮箱邮箱地址，msg：邮箱内容

// type Auth // smtp身份认证机制
// func CRAMMD5Auth(username, secret string) Auth                 // 返回实现RFC 2195定义的CRAM-MD5身份认证机制
// func PlainAuth(identity, username, password, host string) Auth // 返回实现RFC 4616定义的PLAIN身份认证机制

// type Client // 连接smtp服务器客户端
// func Dial(addr string) (*Client, error)                                    // 返回连接addr的smtp服务器客户端（地址需包含端口）
// func NewClient(conn net.Conn, host string) (*Client, error)                // 基于已有连接和host作为服务器名来认证创建的客户端
// func (c *Client) Auth(a Auth) error                                        // 为客户端设置指定的身份认证
// func (c *Client) Close() error                                             // 关闭连接
// func (c *Client) Data() (io.WriteCloser, error)                            // 向服务器发出 DATA 命令并返回可用于编写邮件标题和正文的编写器。
// func (c *Client) Extension(ext string) (bool, string)                      // 报告服务器是否支持此扩展
// func (c *Client) Hello(localName string) error                             // 将 HELO 或 EHLO 作为给定的主机名发送到服务器。
// func (c *Client) Mail(from string) error                                   // 设置发送邮箱地址
// func (c *Client) Noop() error                                              // 想服务器发送NOOP命令，只用于检车与服务器连接是否正常
// func (c *Client) Quit() error                                              // 发送QUIT命令并关闭与服务器连接
// func (c *Client) Rcpt(to string) error                                     // 使用to邮箱地址向服务器发送RCPT命令
// func (c *Client) Reset() error                                             // 向服务器发送RSET命令，中止当前mail事务
// func (c *Client) StartTLS(config *tls.Config) error                        // 发送 STARTTLS 命令并加密所有进一步的通信。只有通告 STARTTLS 扩展的服务器才支持此功能。
// func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool) // 返回客户端的 TLS 连接状态。如果 StartTLS 不成功，则返回值是它们的零值。
// func (c *Client) Verify(addr string) error                                 // 验证检查服务器上电子邮件地址的有效性。

// 记录SMTP服务器信息
// type ServerInfo struct {
//    Name string   // SMTP server name
//    TLS  bool     // using TLS, with valid certificate for Name
//    Auth []string // advertised authentication mechanisms
// }
