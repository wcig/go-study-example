package url

// net/url: 实现URL解析和查询转义

// 函数
// func PathEscape(s string) string             // 对字符串进行转义，可被安全的添加到URL路径path中
// func PathUnescape(s string) (string, error)  // PathEscape逆转换
// func QueryEscape(s string) string            // 对字符串进行转义，可被安全的添加到URL的query中
// func QueryUnescape(s string) (string, error) // QueryEscape逆转换，与PathUnescape区别在于会将'+'反转义为''

// 类型
// 1.Error: URL操作错误
// type Error struct {
//    Op  string
//    URL string
//    Err error
// }
// func (e *Error) Error() string
// func (e *Error) Temporary() bool
// func (e *Error) Timeout() bool
// func (e *Error) Unwrap() error

// 2.EscapeError
// type EscapeError string
// func (e EscapeError) Error() string

// 3.InvalidHostError
// type InvalidHostError string
// func (e InvalidHostError) Error() string

// 4.URL: URL解析 [scheme:][//[userinfo@]host][/]path[?query][#fragment]
// type URL struct {
//    Scheme      string
//    Opaque      string    // encoded opaque data
//    User        *Userinfo // username and password information
//    Host        string    // host or host:port
//    Path        string    // path (relative paths may omit leading slash)
//    RawPath     string    // encoded path hint (see EscapedPath method); added in Go 1.5
//    ForceQuery  bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
//    RawQuery    string    // encoded query values, without '?'
//    Fragment    string    // fragment for references, without '#'
//    RawFragment string    // encoded fragment hint (see EscapedFragment method); added in Go 1.15
// }
// func Parse(rawurl string) (*URL, error)                // 解析rawurl为URL结构，rawurl可以是绝对的（以scheme开头），也可以是相对的（没有host的路径）
// func ParseRequestURI(rawurl string) (*URL, error)      // 解析rawurl为URL结构，假定rawurl是HTTP请求获取到的，并且rawurl没有#fragment后缀
// func (u *URL) EscapedFragment() string                 // 返回u.Fragment的转义（优先使用u.EscapedFragment()而不是u.RawFragment）
// func (u *URL) EscapedPath() string                     // 返回u.Path的转义（优先使用u.EscapedPath()而不是u.RawPath）
// func (u *URL) Hostname() string                        // 返回u.Host，去除端口号
// func (u *URL) IsAbs() bool                             //  报告URL是否是绝对的
// func (u *URL) MarshalBinary() (text []byte, err error) //
// func (u *URL) Parse(ref string) (*URL, error)          // 在接收者的上下文中解析URL
// func (u *URL) Port() string                            // 返回u.Host的端口，不存在则返回空字符串
// func (u *URL) Query() Values                           // 解析RawQuery并返回对应的值（默认丢弃错误键值对，校验错误使用ParseQuery）
// func (u *URL) Redacted() string                        // 将任何密码替换为"xxxxx"
// func (u *URL) RequestURI() string                      // 返回编码的"path?query"或"opaque?query"
// func (u *URL) ResolveReference(ref *URL) *URL          // 从绝对URI u和引用的相对URI，解析出新的URL
// func (u *URL) String() string                          // 重组为有效的URL字符串
// func (u *URL) UnmarshalBinary(text []byte) error       //

// 5.UserInfo: URL中用户名和密码的封装
// type Userinfo struct {
//    // contains filtered or unexported fields
// }
// func User(username string) *Userinfo                   // 解析字符串为Userinfo
// func UserPassword(username, password string) *Userinfo //根据username和password创建Userinfo
// func (u *Userinfo) Password() (string, bool)           // 返回密码的值和其是否已设置
// func (u *Userinfo) String() string                     // 返回标准格式字符串："username[:password]"
// func (u *Userinfo) Username() string                   // 返回用户名

// 6.Values: 键值对列表，与http.Header的map，这里Values区分大小写
// type Values map[string][]string
// func ParseQuery(query string) (Values, error) // 解析URL编码的query字符串为Values
// func (v Values) Add(key, value string)        // 添加键值对
// func (v Values) Del(key string)               // 删除key对应的值
// func (v Values) Encode() string               // 编码为按key排序的形式（"bar=baz&foo=quux"）
// func (v Values) Get(key string) string        // 获取指定key的值，不存在则返回空字符串
// func (v Values) Set(key, value string)        // 设置键值，会覆盖已有值
