package ch30_strconv

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// strconv: 字符串其他类型转换

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
}

// output:
// -2147483648 strconv.ParseInt: parsing "-3546343826724305832": value out of range
// 0 strconv.ParseInt: parsing "1e3": invalid syntax

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

// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
// float64 -> string
func Test(t *testing.T) {
	fmt.Println(strconv.FormatFloat(1.231, 'e', 2, 32))
	fmt.Println(strconv.FormatFloat(1.231, 'f', 2, 32))
	fmt.Println(strconv.FormatFloat(0.000002, 'e', 10, 32))
	fmt.Println(strconv.FormatFloat(0.001, 'e', 3, 32))
}
