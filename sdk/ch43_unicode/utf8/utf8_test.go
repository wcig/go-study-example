package utf8

import (
	"testing"
	"unicode/utf8"
)

// unicode/utf8：实现函数和常量以支持以 UTF-8 编码的文本。它包括在符文和 UTF-8 字节序列之间进行转换的函数。

// 常量
func TestConst(t *testing.T) {
	_ = utf8.RuneError
	_ = utf8.RuneSelf
	_ = utf8.MaxRune
	_ = utf8.UTFMax
}

// 函数
// func DecodeLastRune(p []byte) (r rune, size int)         // 解析字节切片p，返回最后一个字符和其宽度
// func DecodeLastRuneInString(s string) (r rune, size int) // 解析字符串s，返回最后一个字符和其宽度
// func DecodeRune(p []byte) (r rune, size int)             // 解析字节切片p，返回第一个字符和其宽度
// func DecodeRuneInString(s string) (r rune, size int)     // 解析字符串s，返回第一个字符和其宽度
// func EncodeRune(p []byte, r rune) int                    // 将字符r写入到p中，返回写入的字节数，如果超出范围将返回RuneError错误
// func FullRune(p []byte) bool                             // 报告字节切片p是否以完整utf8编码字符开头
// func FullRuneInString(s string) bool                     // 报告字符串s是否以完整utf8编码字符开头
// func RuneCount(p []byte) int                             // 返回字节切片p中的符文数。错误和短编码被视为宽度为1字节的单个符文。
// func RuneCountInString(s string) (n int)                 // 返回字符串s中的符文数。错误和短编码被视为宽度为1字节的单个符文。
// func RuneLen(r rune) int                                 // 返回字符r所需的字节数，如果符文不是以 UTF-8 编码的有效值，则返回 -1。
// func RuneStart(b byte) bool                              // 报告该字节是否可能是编码的、可能无效的符文的第一个字节。第二个和随后的字节总是将前两位设置为 10。（是否是字符第一个字节？）
// func Valid(p []byte) bool                                // 报告字节切片p是否全部由有效的utf8编码字符组成
// func ValidRune(r rune) bool                              // 报告字符r是否可以被合法编码成utf8，超过或替代一半都是非法
// func ValidString(s string) bool                          // 报告字符串s是否全部由有效的utf8编码字符组成
