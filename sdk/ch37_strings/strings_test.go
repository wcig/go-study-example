package ch37_strings

// strings: 处理UTF-8编码字符串函数

// 函数
// func Compare(a, b string) int                               // 比较字符串a和b（a==b返回0，a>b返回1，a<b返回-1）
// func Contains(s, substr string) bool                        // s是否包含子串substr
// func ContainsAny(s, chars string) bool                      // s是否包含字符串chars中任意字符
// func ContainsRune(s string, r rune) bool                    // s是否包含字符r
// func Count(s, substr string) int                            // s中substr子串出现数量（substr空则返回s的长度+1）
// func EqualFold(s, t string) bool                            // s和t忽略大小写是否相等
// func Fields(s string) []string                              // 根据一个或多个空白字符拆分字符串s为子串切片
// func FieldsFunc(s string, f func(rune) bool) []string       // 根据函数f是否返回true来拆分字符串s
// func HasPrefix(s, prefix string) bool                       // s是否有前缀prefix
// func HasSuffix(s, suffix string) bool                       // s是否有后缀suffix
// func Index(s, substr string) int                            // s中子串substr第一次出现索引（不存在返回-1）
// func IndexAny(s, chars string) int                          // s中字符串chars中任意字符第一次出现索引（不存在返回-1）
// func IndexByte(s string, c byte) int                        // s中字节c第一次出现索引（不存在返回-1）
// func IndexFunc(s string, f func(rune) bool) int             // s中第一个满足f函数为true的索引（不存在返回-1）
// func IndexRune(s string, r rune) int                        // s中第一次出现r的索引（不存在返回-1）
// func Join(elems []string, sep string) string                // 字符串切片elems以seq为分隔符合并为单个字符串
// func LastIndex(s, substr string) int                        // s中substr最后一次出现的索引（不存在返回-1）
// func LastIndexAny(s, chars string) int                      // s中字符串chars中任意字符最后一次出现索引（不存在返回-1）
// func LastIndexByte(s string, c byte) int                    // s中字节c最后一次出现索引（不存在返回-1）
// func LastIndexFunc(s string, f func(rune) bool) int         // s中最后一个满足f函数为true的索引（不存在返回-1）
// func Map(mapping func(rune) rune, s string) string          // 返回s中所有字符根据映射函数替换后的字符串（映射返回复制则删除该字符不替换）
// func Repeat(s string, count int) string                     // s重复count次数组合后的字符串（count为负数将panic）
// func Replace(s, old, new string, n int) string              // 替换s中前n个old子串为new（如果n小于0则替换所有）
// func ReplaceAll(s, old, new string) string                  // 替换s中所有old子串为new
// func Split(s, sep string) []string                          // seq为分隔符将s拆分为字符串切片，如果s不包含seq则返回长度为1元素为s的切片；如果seq为空则在每个UTF8序列之后拆分；如果s和seq都为空则返回一个空切片
// func SplitAfter(s, sep string) []string                     // 在seq的每个实例之后将s拆分为字符串切片，相当于SplitN(s,seq,-1)
// func SplitAfterN(s, sep string, n int) []string             // 在seq的每个实例之后将s拆分为字符串切片，n>0: 最后n个子串，最后一个为拆分的；n==0：零子串；n<0：所有子串
// func SplitN(s, sep string, n int) []string                  // seq作为分隔符拆分s为字符串切片，n取值效果类似SplitAfterN
// func Title(s string) string                                 // 字符串s所有unicode字符替换为大小（BUG(rsc)：标题用于单词边界的规则不能正确处理 Unicode 标点符号。）
// func ToLower(s string) string                               // s所有unicode字母转为小写
// func ToLowerSpecial(c unicode.SpecialCase, s string) string // 使用指定映射替换s的unicode字母为小写
// func ToTitle(s string) string                               // s所有unicode字母替换为大小
// func ToTitleSpecial(c unicode.SpecialCase, s string) string // 使用指定映射替换s的unicode字母为大写
// func ToUpper(s string) string                               // s所有unicode字母转大写
// func ToUpperSpecial(c unicode.SpecialCase, s string) string // 使用指定映射替换s的unicode字母为大写
// func ToValidUTF8(s, replacement string) string              // 将s中无效的UTF8替换为replacement
// func Trim(s, cutset string) string                          // 去除s中所有的前后cutset出现的unicode字符
// func TrimFunc(s string, f func(rune) bool) string           // 删除s中前后所有函数f为true的unicode字符
// func TrimLeft(s, cutset string) string                      // 去除s中所有的前cutset出现的unicode字符
// func TrimLeftFunc(s string, f func(rune) bool) string       // 删除s中前面所有函数f为true的unicode字符
// func TrimPrefix(s, prefix string) string                    // 去除s的前缀prefix
// func TrimRight(s, cutset string) string                     // 去除s中所有的后cutset出现的unicode字符
// func TrimRightFunc(s string, f func(rune) bool) string      // 删除s中后面所有函数f为true的unicode字符
// func TrimSpace(s string) string                             // 去除s前后空格
// func TrimSuffix(s, suffix string) string                    // 去除s的后缀suffix
