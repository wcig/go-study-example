package utf16

// unicode/utf16：实现了 UTF-16 序列的编码和解码。

// 函数
// func Decode(s []uint16) []rune        // 返回由 UTF-16 编码表示的 Unicode 代码点序列
// func DecodeRune(r1, r2 rune) rune     // 返回代理对的 UTF-16 解码。如果该对不是有效的 UTF-16 代理对，DecodeRune 将返回 Unicode 替换代码点 U+FFFD。
// func Encode(s []rune) []uint16        // 返回 Unicode 代码点序列 s 的 UTF-16 编码
// func EncodeRune(r rune) (r1, r2 rune) // 返回给定符文的 UTF-16 代理对 r1、r2。如果符文不是有效的 Unicode 代码点或不需要编码，则 EncodeRune 返回 U+FFFD、U+FFFD。
// func IsSurrogate(r rune) bool         // 报告指定的 Unicode 代码点是否可以出现在代理项对中
