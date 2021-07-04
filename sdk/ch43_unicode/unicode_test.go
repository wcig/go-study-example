package ch43_unicode

import (
	"testing"
	"unicode"
)

// unicode：提供数据和函数来测试unicode代码点的一些属性

// 常量
func TestConst(t *testing.T) {
	_ = unicode.MaxRune
	_ = unicode.ReplacementChar
	_ = unicode.MaxASCII
	_ = unicode.MaxLatin1

	_ = unicode.UpperCase
	_ = unicode.LowerCase
	_ = unicode.TitleCase
	_ = unicode.MaxCase

	_ = unicode.UpperLower
}

// 函数
// func In(r rune, ranges ...*RangeTable) bool     // 报告r是否在ranges范围内
// func Is(rangeTab *RangeTable, r rune) bool      // 报告r是否在rangeTab范围内
// func IsControl(r rune) bool                     // 报告r是否是控制字符
// func IsDigit(r rune) bool                       // 报告r是否是十进制数字
// func IsGraphic(r rune) bool                     // 报告r是否被 Unicode 定义为图形。此类字符包括来自 L、M、N、P、S、Zs 类别的字母、标记、数字、标点符号、符号和空格。
// func IsLetter(r rune) bool                      // 报告r是否为字母（L类）
// func IsLower(r rune) bool                       // 报告r是否为小写字母
// func IsMark(r rune) bool                        // 报告r是否为标记字符（M类）
// func IsNumber(r rune) bool                      // 报告r是否为数字（N类）
// func IsOneOf(ranges []*RangeTable, r rune) bool // 报告r是否是ranges范围之一的成员（函数In提供更好的签名优先于InOneOf使用）
// func IsPrint(r rune) bool                       // 报告r是否被Go定义为可打印字符
// func IsPunct(r rune) bool                       // 报告r是否为unicode标点字符（P类）
// func IsSpace(r rune) bool                       // 报告r是否为空格字符
// func IsSymbol(r rune) bool                      // 报告r是否为符号字符
// func IsTitle(r rune) bool                       // 报告r是否为标题大小写字母
// func IsUpper(r rune) bool                       // 报告r是否为大写字母
// func SimpleFold(r rune) rune                    // 在 Unicode 定义的简单大小写折叠下迭代等效的 Unicode 代码点
// func To(_case int, r rune) rune                 // 将r转换为指定case字符：UpperCase、LowerCase、TitleCase
// func ToLower(r rune) rune                       // 将r转换为小写
// func ToTitle(r rune) rune                       // 将r转换为标题大小写
// func ToUpper(r rune) rune                       // 将r转换为大写

// 类型
// 1.CaseRange：大小写转换的unicode范围
// type CaseRange struct {
//    Lo    uint32
//    Hi    uint32
//    Delta d
// }

// 2.Range16：表示 16 位 Unicode 代码点的范围。范围从 Lo 到 Hi 包含在内，并具有指定的步幅。
// type Range16 struct {
//    Lo     uint16
//    Hi     uint16
//    Stride uint16
// }

// 3.Range32：表示 32 位 Unicode 代码点的范围。
// type Range32

// 4.RangeTable：通过列出集合中的代码点范围来定义一组 Unicode 代码点。
// type RangeTable struct {
//    R16         []Range16
//    R32         []Range32
//    LatinOffset int // number of entries in R16 with Hi <= MaxLatin1; added in Go 1.1
// }

// 5.SpecialCase：表示特定于语言的大小写映射，例如土耳其语。 SpecialCase 的方法自定义（通过覆盖）标准映射
// type SpecialCase
// func (special SpecialCase) ToLower(r rune) rune
// func (special SpecialCase) ToTitle(r rune) rune
// func (special SpecialCase) ToUpper(r rune) rune
