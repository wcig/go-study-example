package ch36_strconv

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Itoa(i int) string: int -> string
func TestItoa(t *testing.T) {
	type itoaTest struct {
		num      int
		expected string
	}
	its := []*itoaTest{
		{123, "123"},
		{-123, "-123"},
		{0, "0"},
	}
	for _, it := range its {
		assert.Equal(t, it.expected, strconv.Itoa(it.num))
	}
}

// func Atoi(s string) (int, error): string -> int
func TestAtoi(t *testing.T) {
	type atoiTest struct {
		str      string
		expected int
	}
	ats := []*atoiTest{
		{"123", 123},
		{"-123", -123},
		{"0", 0},
		{"01", 1},
	}
	for _, at := range ats {
		val, err := strconv.Atoi(at.str)
		assert.Nil(t, err)
		assert.Equal(t, at.expected, val)
	}
	fmt.Println(strconv.Atoi("01a")) // 0 strconv.Atoi: parsing "01a": invalid syntax
}

// func FormatInt(i int64, base int) string
// func FormatUint(i uint64, base int) string
// 以指定进制转换数字为字符串 (2<=base<=36)
func TestFormatInt(t *testing.T) {
	type formatIntTest struct {
		num      int64
		base     int
		expected string
	}
	fts := []*formatIntTest{
		{123, 10, "123"},
		{-123, 10, "-123"},
		{0, 10, "0"},
		{15, 2, "1111"},
		{15, 8, "17"},
		{31, 16, "1f"},
		{1e3, 10, "1000"},
	}
	for _, ft := range fts {
		assert.Equal(t, ft.expected, strconv.FormatInt(ft.num, ft.base))
	}
}

// func ParseInt(s string, base int, bitSize int) (i int64, err error)
// func ParseUint(s string, base int, bitSize int) (uint64, error)
// 以指定进制指定位数转换字符串为数字 (0,2<=base<=36, 0<=bitSize<=64)
func TestParseInt(t *testing.T) {
	type parseIntTest struct {
		str      string
		base     int
		bitSize  int
		Expected int64
	}
	pts := []*parseIntTest{
		{"-354634382", 10, 32, -354634382},
		{"-3546343826724305832", 10, 64, -3546343826724305832},
		{"1111", 2, 32, 15},
		{"17", 8, 32, 15},
		{"1f", 16, 32, 31},
	}
	for _, pt := range pts {
		val, err := strconv.ParseInt(pt.str, pt.base, pt.bitSize)
		assert.Nil(t, err)
		assert.Equal(t, pt.Expected, val)
	}
	fmt.Println(strconv.ParseInt("-3546343826724305832", 10, 32))
	fmt.Println(strconv.ParseInt("1e3", 10, 32))
	// output:
	// -2147483648 strconv.ParseInt: parsing "-3546343826724305832": value out of range
	// 0 strconv.ParseInt: parsing "1e3": invalid syntax
}

// func FormatBool(b bool) string: bool -> string
func TestFormatBool(t *testing.T) {
	assert.Equal(t, "true", strconv.FormatBool(true))
	assert.Equal(t, "false", strconv.FormatBool(false))
}

// func ParseBool(str string) (bool, error): string -> bool
func TestParseBool(t *testing.T) {
	var (
		val bool
		err error
	)

	val, err = strconv.ParseBool("true")
	assert.Nil(t, err)
	assert.Equal(t, true, val)
	val, err = strconv.ParseBool("false")
	assert.Nil(t, err)
	assert.Equal(t, false, val)
}

// func FormatFloat(f float64, fmt byte, prec, bitSize int) string: float64 -> string
func TestFormatFloat(t *testing.T) {
	fmt.Println(strconv.FormatFloat(0.001, 'e', 3, 32))
	fmt.Println(strconv.FormatFloat(0.001, 'E', 3, 32))

	fmt.Println(strconv.FormatFloat(1.23456, 'e', 2, 32))
	fmt.Println(strconv.FormatFloat(1.23456, 'f', 2, 32))

	fmt.Println(strconv.FormatFloat(1.23456, 'e', -1, 32))
	fmt.Println(strconv.FormatFloat(1.23456, 'f', -1, 32))

	v := 3.1415926535
	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
	// output:
	// 1.000e-03
	// 1.000E-03
	// 1.23e+00
	// 1.23
	// 1.23456e+00
	// 1.23456
	// string, 3.1415927E+00
	// string, 3.1415926535E+00
}

// func ParseFloat(s string, bitSize int) (float64, error): string -> float
func TestParseFloat(t *testing.T) {
	fmt.Println(strconv.ParseFloat("1.23456", 64))
	fmt.Println(strconv.ParseFloat("3.1415926535", 32))
	fmt.Println(strconv.ParseFloat("3.1415926535", 64))
	// output:
	// 1.23456 <nil>
	// 3.1415927410125732 <nil>
	// 3.1415926535 <nil>
}

// func AppendBool(dst []byte, b bool) []byte: 添加bool类型字符串到目标字符切片后
func TestAppendBool(t *testing.T) {
	type appendBoolTest struct {
		dst      string
		b        bool
		expected string
	}
	ats := []*appendBoolTest{
		{"ok-", true, "ok-true"},
		{"ok-", false, "ok-false"},
	}
	for _, at := range ats {
		val := strconv.AppendBool([]byte(at.dst), at.b)
		assert.Equal(t, at.expected, string(val))
	}
}

// func AppendInt(dst []byte, i int64, base int) []byte
// func AppendUint(dst []byte, i uint64, base int) []byte
// 添加int类型以指定进制转换的字符串到最后
func TestAppendInt(t *testing.T) {
	type appendIntTest struct {
		dst      string
		i        int64
		base     int
		expected string
	}
	ats := []*appendIntTest{
		{"ok", 123, 10, "ok123"},
		{"ok", -123, 10, "ok-123"},
		{"ok", 7, 2, "ok111"},
		{"ok", 15, 8, "ok17"},
		{"ok", 31, 16, "ok1f"},
	}
	for _, at := range ats {
		val := strconv.AppendInt([]byte(at.dst), at.i, at.base)
		assert.Equal(t, at.expected, string(val))
	}
}

// func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte: 添加float类型转换后字符串到最后
func TestAppendFloat(t *testing.T) {
	fmt.Println(string(strconv.AppendFloat([]byte("ok"), 1.23456, 'f', -1, 64)))
	fmt.Println(string(strconv.AppendFloat([]byte("ok"), 1.23456, 'e', 2, 64)))
	fmt.Println(string(strconv.AppendFloat([]byte("ok"), 1.23456, 'E', 2, 64)))
	// output:
	// ok1.23456
	// ok1.23e+00
	// ok1.23E+00
}

// Quote: 添加引号
func TestQuote(t *testing.T) {
	// func Quote(s string) string: 字符串添加双引号
	fmt.Println(strconv.Quote("ok"))
	// func QuoteRune(r rune) string: 字符添加单引号
	fmt.Println(strconv.QuoteRune('a'))
	// func QuoteRuneToASCII(r rune) string: 字符先ascii转换再添加单引号
	fmt.Println(strconv.QuoteRuneToASCII('好'))
	// func QuoteToASCII(s string) string: 字符串先ascii转换再添加双引号
	fmt.Println(strconv.QuoteToASCII("好"))
	// func QuoteRuneToGraphic(r rune) string: 字符先对IsGraphic函数定义的非ascii或不可打印字符转义再添加单引号
	fmt.Println(strconv.QuoteRuneToGraphic('好'))
	// func QuoteToGraphic(s string) string: 字符串先对IsGraphic函数定义的非ascii或不可打印字符转义再添加双引号
	fmt.Println(strconv.QuoteToGraphic("好"))
}

// func Unquote(s string) (string, error): 将s解释为单引号/双引号/反引号字符串
func TestUnquote(t *testing.T) {
	fmt.Println(strconv.Unquote(`"ok"`))
	fmt.Println(strconv.Unquote("`ok`"))
	fmt.Println(strconv.Unquote(`'o'`))
	fmt.Println(strconv.Unquote("ok"))
	// output:
	// ok <nil>
	// ok <nil>
	// o <nil>
	// invalid syntax
}

// func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error):
// 解码转义的字符串或由字符串s表示的字符文字中的第一个字符或字节
func TestUnquoteChar(t *testing.T) {
	v, mb, tt, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", tt)
	// output:
	// value: "
	// multibyte: false
	// tail: Fran & Freddie's Diner\"
}

// AppendQuote: 添加引号包裹的字符/字符串到最后
func TestAppendQuote(t *testing.T) {
	fmt.Println(string(strconv.AppendQuote([]byte("ok"), "tom")))
	fmt.Println(string(strconv.AppendQuoteToASCII([]byte("ok"), "好")))
	fmt.Println(string(strconv.AppendQuoteToGraphic([]byte("ok"), "好")))

	fmt.Println(string(strconv.AppendQuoteRune([]byte("ok"), '好')))
	fmt.Println(string(strconv.AppendQuoteRuneToASCII([]byte("ok"), '好')))
	fmt.Println(string(strconv.AppendQuoteRuneToGraphic([]byte("ok"), '好')))
	// output:
	// ok"tom"
	// ok"\u597d"
	// ok"好"
	// ok'好'
	// ok'\u597d'
	// ok'好'
}

// func IsGraphic(r rune) bool:
// IsGraphic报告是否通过Unicode将符文定义为“图形”。 此类字符包括字母，标记，数字，标点符号，符号和空格，来自类别L，M，N，P，S和Zs。
func TestIsGraphic(t *testing.T) {
	assert.Equal(t, true, strconv.IsGraphic('*'))
	assert.Equal(t, true, strconv.IsGraphic('a'))
	assert.Equal(t, true, strconv.IsGraphic('\u597d'))
	assert.Equal(t, false, strconv.IsGraphic('\007'))
}

// func IsPrint(r rune) bool: 判断是否为go可打印字符
func TestIsPrint(t *testing.T) {
	assert.Equal(t, true, strconv.IsPrint('*'))
	assert.Equal(t, true, strconv.IsPrint('a'))
	assert.Equal(t, true, strconv.IsPrint('\u597d'))
	assert.Equal(t, false, strconv.IsPrint('\007'))
}

// func CanBackquote(s string) bool:
// CanBackquote报告字符串s是否可以不变地表示为单行反引号字符串,且没有制表符以外的控制字符
func TestCanBackquote(t *testing.T) {
	assert.Equal(t, true, strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	assert.Equal(t, false, strconv.CanBackquote("`can't backquote this`"))
}

// ------------------------------------------------------------------- //

// 基础类型->字符串
func TestTypeToStr(t *testing.T) {
	s1 := strconv.Itoa(123)
	fmt.Println(s1)

	s2 := strconv.FormatInt(123, 10)
	fmt.Println(s2)

	s3 := strconv.FormatUint(uint64(132), 10)
	fmt.Println(s3)

	s4 := strconv.FormatBool(true)
	fmt.Println(s4)

	s5 := strconv.FormatFloat(1.23, 'f', -1, 64)
	fmt.Println(s5)

	s6 := strconv.FormatComplex(1+2i, 'g', -1, 128)
	fmt.Println(s6)
	// output:
	// 123
	// 123
	// 132
	// true
	// 1.23
	// (1+2i)
}

// 字符串->基础类型
func TestStrToType(t *testing.T) {
	n, err := strconv.Atoi("123")
	fmt.Println(n, err)

	m, err := strconv.ParseInt("123", 10, 64)
	fmt.Println(m, err)

	u, err := strconv.ParseUint("123", 10, 64)
	fmt.Println(u, err)

	f, err := strconv.ParseFloat("1.23", 64)
	fmt.Println(f, err)

	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)

	c, err := strconv.ParseComplex("1+2i", 128)
	fmt.Println(c, err)
	// output:
	// 123 <nil>
	// 123 <nil>
	// 123 <nil>
	// 1.23 <nil>
	// true <nil>
	// (1+2i) <nil>
}

// 引号
func TestAboutQuote(t *testing.T) {
	fmt.Println(strconv.Quote("ok"))
	fmt.Println(strconv.QuoteRune('a'))
	fmt.Println(strconv.QuoteRuneToASCII('好'))
	fmt.Println(strconv.QuoteToASCII("好"))
	fmt.Println(strconv.QuoteRuneToGraphic('好'))
	fmt.Println(strconv.QuoteToGraphic("好"))

	s1, err := strconv.Unquote(`"ok"`)
	fmt.Println(s1, err)
	v, m, tail, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	fmt.Println(string(v), m, tail, err)
	// output:
	// "ok"
	// 'a'
	// '\u597d'
	// "\u597d"
	// '好'
	// "好"
	// ok <nil>
	// " false Fran & Freddie's Diner\" <nil>
}

// 拼接基础类型字符串
func TestAppendType(t *testing.T) {
	dst := []byte("ok-")
	b1 := strconv.AppendBool(dst, true)
	fmt.Println(string(b1))

	b2 := strconv.AppendInt(dst, 123, 10)
	fmt.Println(string(b2))

	b3 := strconv.AppendUint(dst, 123, 10)
	fmt.Println(string(b3))

	b4 := strconv.AppendFloat(dst, 1.23, 'f', -1, 64)
	fmt.Println(string(b4))

	b5 := strconv.AppendQuote(dst, "tom")
	fmt.Println(string(b5))

	b6 := strconv.AppendQuoteToASCII(dst, "好")
	fmt.Println(string(b6))

	b7 := strconv.AppendQuoteToGraphic(dst, "好")
	fmt.Println(string(b7))

	b8 := strconv.AppendQuoteRune(dst, '好')
	fmt.Println(string(b8))

	b9 := strconv.AppendQuoteRuneToASCII(dst, '好')
	fmt.Println(string(b9))

	b10 := strconv.AppendQuoteRuneToGraphic(dst, '好')
	fmt.Println(string(b10))
	// output:
	// ok-true
	// ok-123
	// ok-123
	// ok-1.23
	// ok-"tom"
	// ok-"\u597d"
	// ok-"好"
	// ok-'好'
	// ok-'\u597d'
	// ok-'好'
}

// 其他
func TestOther(t *testing.T) {
	TestCanBackquote(t)
	TestIsGraphic(t)
	TestIsPrint(t)
}
