package ch36_strconv

import (
	"strconv"
	"testing"
)

// strconv: 字符串其他类型转换

// 错误
func TestErr(t *testing.T) {
	_ = strconv.ErrRange  // value out of range
	_ = strconv.ErrSyntax // invalid syntax
}

// 函数

// 1.基础类型->字符串
// func Itoa(i int) string                                              // int
// func FormatBool(b bool) string                                       // bool
// func FormatComplex(c complex128, fmt byte, prec, bitSize int) string // complex
// func FormatFloat(f float64, fmt byte, prec, bitSize int) string      // float
// func FormatInt(i int64, base int) string                             // int
// func FormatUint(i uint64, base int) string                           // uint

// 2.字符串->基础类型
// func Atoi(s string) (int, error)                                    // int
// func ParseBool(str string) (bool, error)                            // bool
// func ParseComplex(s string, bitSize int) (complex128, error)        // complex
// func ParseFloat(s string, bitSize int) (float64, error)             // float
// func ParseInt(s string, base int, bitSize int) (i int64, err error) // int
// func ParseUint(s string, base int, bitSize int) (uint64, error)     // uint

// 3.引号相关
// func Quote(s string) string                                                                 // 字符串添加双引号
// func QuoteRune(r rune) string                                                               // 字符添加单引号
// func QuoteRuneToASCII(r rune) string                                                        // 字符先ascii转换再添加单引号
// func QuoteRuneToGraphic(r rune) string                                                      // 字符先对IsGraphic函数定义的非ascii或不可打印字符转义再添加单引号
// func QuoteToASCII(s string) string                                                          // 字符串先ascii转换再添加双引号
// func QuoteToGraphic(s string) string                                                        // 字符串先对IsGraphic函数定义的非ascii或不可打印字符转义再添加双引号
// func Unquote(s string) (string, error)                                                      // 去除双引号
// func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) // 去除单引号

// 4.返回字节切片dst添加基础类型字符串后的结果
// func AppendBool(dst []byte, b bool) []byte
// func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
// func AppendInt(dst []byte, i int64, base int) []byte
// func AppendQuote(dst []byte, s string) []byte
// func AppendQuoteRune(dst []byte, r rune) []byte
// func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
// func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
// func AppendQuoteToASCII(dst []byte, s string) []byte
// func AppendQuoteToGraphic(dst []byte, s string) []byte
// func AppendUint(dst []byte, i uint64, base int) []byte

// 5.其他
// func CanBackquote(s string) bool // 报告字符串s是否可以不变地表示为单行反引号字符串,且没有制表符以外的控制字符
// func IsGraphic(r rune) bool      // 报告是否通过Unicode将符文定义为“图形”
// func IsPrint(r rune) bool        // 判断是否为go可打印字符

// 结构体
// type NumError: 格式转换错误
// func (e *NumError) Error() string
// func (e *NumError) Unwrap() error
