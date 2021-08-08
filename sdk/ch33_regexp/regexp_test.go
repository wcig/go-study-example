package ch33_regexp

// regexp: 正则表达式搜索。

// 函数
// func Match(pattern string, b []byte) (matched bool, err error)              // 报告字节切片b是否匹配正则表达式pattern
// func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) // 报告RuneReader r是否匹配正则表达式pattern
// func MatchString(pattern string, s string) (matched bool, err error)        // 报告字符串s是否匹配正则表达式pattern
// func QuoteMeta(s string) string                                             // 对正则表达式进行编码

// 类型
// Regexp：编译的正则表达式
// type Regexp struct {
//     // contains filtered or unexported fields
// }
// func Compile(expr string) (*Regexp, error)                                                  // 编译解析正则表达式并返回
// func CompilePOSIX(expr string) (*Regexp, error)                                             // 跟Compile一样，不过限制正则表达式为POSIX ERE语法，并改变匹配语义诶最左边最长
// func MustCompile(str string) *Regexp                                                        // 跟Compile一样，但是表达式无法被解析将出现panic
// func MustCompilePOSIX(str string) *Regexp                                                   // 跟CompilePOSIX一样，但是错误就导致panic
// func (re *Regexp) Copy() *Regexp                                                            // 从re复制返回一新的Regexp对象，在副本调用不会影响另一个
// func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte       // 将src内容的match部分，按re正则表达式匹配为模板template格式，并将输出内容添加到src
// func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte // 同Expand基本一样，参数换为字符串
// func (re *Regexp) Find(b []byte) []byte                                                     // 返回内容b中匹配正则表达式re的最左匹配为本切片，没有返回nil
// func (re *Regexp) FindAll(b []byte, n int) [][]byte                                         // Find的all版本，n为-1返回所有匹配，没有返回nil
// func (re *Regexp) FindAllIndex(b []byte, n int) [][]int                                     // FindIndex的all版本
// func (re *Regexp) FindAllString(s string, n int) []string                                   // FindString的all版本，返回所有匹配表达式的切片
// func (re *Regexp) FindAllStringIndex(s string, n int) [][]int                               // FindString类似，返回的是匹配内容的index
// func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string                         // FindStringSubmatch的all版本
// func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int                       // FindAllStringSubmatch类似，返回匹配内容的索引
// func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte                               // FindSubmatch的all版本
// func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int                             // FindSubmatchIndex的all版本
// func (re *Regexp) FindIndex(b []byte) (loc []int)                                           // 返回2个int类型元素的切片，其定义了内容b与正则表达式re最左匹配的位置，匹配的位置为b[loc[0]:loc[1]]，不匹配返回nil
// func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)                              // FindIndex的io.RuneReader版本
// func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int                            // FindSubmatchIndex的io.RuneReader版本
// func (re *Regexp) FindString(s string) string                                               // 返回一个内容s与正则表达式re最左匹配的文本，不匹配则返回空字符串
// func (re *Regexp) FindStringIndex(s string) (loc []int)                                     // FindString类似，返回的匹配内容再s的index
// func (re *Regexp) FindStringSubmatch(s string) []string                                     // 返回一字符串切片，其中包含是中正则表达式最左匹配文本和其匹配的子表达式如果有，不匹配则返回nil
// func (re *Regexp) FindStringSubmatchIndex(s string) []int                                   // FindStringSubmatch类似，返回的index
// func (re *Regexp) FindSubmatch(b []byte) [][]byte                                           // 返回b中正则表达式最左匹配的文本和其子表达式的匹配如果有，不匹配则返回nil
// func (re *Regexp) FindSubmatchIndex(b []byte) []int                                         // 与FindSubmatch类似，返回的是index
// func (re *Regexp) LiteralPrefix() (prefix string, complete bool)                            // 返回一个文字字符串，该字符串必须以正则表达式 re 的任何匹配开始。如果文字字符串包含整个正则表达式，则返回布尔值 true。
// func (re *Regexp) Longest()                                                                 // 使以后的搜索更喜欢最左边最长的匹配项
// func (re *Regexp) Match(b []byte) bool                                                      // 报告b是否包含匹配正则表达式re匹配的部分
// func (re *Regexp) MatchReader(r io.RuneReader) bool                                         // Match的io.RuneReader版本
// func (re *Regexp) MatchString(s string) bool                                                // Match的字符串版本
// func (re *Regexp) NumSubexp() int                                                           // 返回正则表达式re中空号子表达式的数量
// func (re *Regexp) ReplaceAll(src, repl []byte) []byte                                       // 替换src中与正则表达式匹配部分为repl，$解释为展示，即$1表示第一个自匹配文本
// func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte               // ReplaceAll类似，从替换为repl到按repl指定函数替换
// func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte                                // 与ReplaceAll类似，但不支持使用Expand
// func (re *Regexp) ReplaceAllLiteralString(src, repl string) string                          // ReplaceAllLiteral的字符串版本
// func (re *Regexp) ReplaceAllString(src, repl string) string                                 // ReplaceAll的字符串版本
// func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string         // ReplaceAllFunc的字符串版本
// func (re *Regexp) Split(s string, n int) []string                                           // 切片s分成由表达式分隔的子字段，并返回这些表达式匹配之间的子串的片。
// func (re *Regexp) String() string                                                           // 返回用于编译正则表达式的源文本
// func (re *Regexp) SubexpIndex(name string) int                                              // 返回给定name的第一个子表达式的index，没有返回-1
// func (re *Regexp) SubexpNames() []string                                                    // 返回此Regexp中括号子表单的名称
