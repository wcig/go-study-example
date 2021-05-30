package ch4_bytes

import (
	"bytes"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// bytes包 TODO

// 字节切片校验相等 (nil和空切片相等)
func TestEqual(t *testing.T) {
	b1 := []byte("hello")
	b2 := []byte("hello")
	b3 := []byte("ok")
	var b4 []byte

	assert.True(t, bytes.Equal(b1, b2))
	assert.False(t, bytes.Equal(b1, b3))
	assert.True(t, bytes.Equal(b4, nil))
}

// 字节切片比较 (0:a==b, 1:a>b, -1:a<b)
func TestCompare(t *testing.T) {
	b1 := []byte("hello")
	b2 := []byte("hello")
	b3 := []byte("ok")
	var b4 []byte

	println(bytes.Compare(b1, b2))
	println(bytes.Compare(b1, b3))
	println(bytes.Compare(b4, nil))
	// output:
	// 0
	// -1
	// 0
}

// 子切片seq在切片s出现的次数
func TestCount(t *testing.T) {
	b1 := []byte("hello")
	b2 := []byte("hello ok")
	b3 := []byte("hello ok hello")
	b4 := []byte("ok")
	var b5 []byte

	println(bytes.Count(b2, b1))
	println(bytes.Count(b3, b1))
	println(bytes.Count(b4, b1))
	println(bytes.Count(b5, b1))
	// output:
	// 1
	// 2
	// 0
	// 0
}

// b切片是否包含subslice子切片
func TestContains(t *testing.T) {
	b1 := []byte("hello")
	b2 := []byte("hello ok")
	b3 := []byte("ok")
	var b4 []byte

	println(bytes.Contains(b2, b1))
	println(bytes.Contains(b3, b1))
	println(bytes.Contains(b4, b1))
	// output:
	// true
	// false
	// false
}

// 返回b字节切片是否包含chars字符串的任意一个字符
func TestContainsAny(t *testing.T) {
	s := "hello"
	b1 := []byte("hello ok")
	b2 := []byte("ck")
	var b3 []byte

	println(bytes.ContainsAny(b1, s))
	println(bytes.ContainsAny(b2, s))
	println(bytes.ContainsAny(b3, s))
	// output:
	// true
	// false
	// false
}

// 返回b字节切片是否包含字符
func TestContainsRune(t *testing.T) {
	r := '好'
	b1 := []byte("你好")
	b2 := []byte("hello")
	var b3 []byte

	println(bytes.ContainsRune(b1, r))
	println(bytes.ContainsRune(b2, r))
	println(bytes.ContainsRune(b3, r))
	// output:
	// true
	// false
	// false
}

// 返回seq字节切片在b字节切片第一次出现的位置,没有则返回-1
func TestIndex(t *testing.T) {
	seq := []byte("hello")
	b1 := []byte("ok hello")
	b2 := []byte("ok")
	var b3 []byte

	println(bytes.Index(b1, seq))
	println(bytes.Index(b2, seq))
	println(bytes.Index(b3, seq))
}

// 返回字节c在字节切片b第一次出现的位置,没有则返回-1
func TestIndexByte(t *testing.T) {
	var c byte = 'h'
	b1 := []byte("hello")
	b2 := []byte("ok")
	var b3 []byte

	println(bytes.IndexByte(b1, c))
	println(bytes.IndexByte(b2, c))
	println(bytes.IndexByte(b3, c))
	// output:
	// 0
	// -1
	// -1
}

// 返回s字符串任意字符在b字节切片第一次出现位置
func TestIndexAny(t *testing.T) {
	s := "hello好"
	b1 := []byte("hello ok")
	b2 := []byte("你好")
	b3 := []byte("ck")
	var b4 []byte

	println(bytes.IndexAny(b1, s))
	println(bytes.IndexAny(b2, s))
	println(bytes.IndexAny(b3, s))
	println(bytes.IndexAny(b4, s))
	// output:
	// 0
	// 3
	// -1
	// -1
}

// 返回字符r在字节切片b第一次出现位置.没有则返回-1
func TestIndexRune(t *testing.T) {
	r := '好'
	b1 := []byte("你好")
	b2 := []byte("hello")
	var b3 []byte

	println(bytes.IndexRune(b1, r))
	println(bytes.IndexRune(b2, r))
	println(bytes.IndexRune(b3, r))
	// output:
	// 3
	// -1
	// -1
}

// 将 s 解释为一系列 UTF-8 编码的代码点。 它返回满足f(c)的第一个Unicode代码点的s中的字节索引，如果没有，则返回-1。
func TestIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	s1 := []byte("Hello, 世界")
	s2 := []byte("Hello, world")

	println(bytes.IndexFunc(s1, f))
	println(bytes.IndexFunc(s2, f))
	// output:
	// 7
	// -1
}

// 返回s字符串任意字符在b字节切片最后一次出现位置
func TestLastIndex(t *testing.T) {
	seq := []byte("hello")
	b1 := []byte("hello ok hello")
	b2 := []byte("ok")
	var b3 []byte

	println(bytes.LastIndex(b1, seq))
	println(bytes.LastIndex(b2, seq))
	println(bytes.LastIndex(b3, seq))
	// output:
	// 9
	// -1
	// -1
}

// 返回字节c在字节切片b最后一次出现的位置,没有则返回-1
func TestLastIndexByte(t *testing.T) {
	var c byte = 'l'
	b1 := []byte("hello")
	b2 := []byte("ok")
	var b3 []byte

	println(bytes.LastIndexByte(b1, c))
	println(bytes.LastIndexByte(b2, c))
	println(bytes.LastIndexByte(b3, c))
	// output:
	// 3
	// -1
	// -1
}

// 返回s字符串任意字符在b字节切片最后一次出现位置
func TestLastIndexAny(t *testing.T) {
	s := "hello"
	b1 := []byte("hello ok hello")
	b2 := []byte("你好")
	b3 := []byte("ck")
	var b4 []byte

	println(bytes.LastIndexAny(b1, s))
	println(bytes.LastIndexAny(b2, s))
	println(bytes.LastIndexAny(b3, s))
	println(bytes.LastIndexAny(b4, s))
	// output:
	// 13
	// -1
	// -1
	// -1
}

// 将 s 解释为一系列 UTF-8 编码的代码点。 它返回满足 f(c) 的最后一个 Unicode 代码点在 s 中的字节索引，如果没有，则返回 -1。
func TestLastIndexFunc(t *testing.T) {
	println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))
	// output:
	// 8
	// 9
	// -1
}

// 返回字节切片s去除所有前后cutset子串出现的字符
func TestTrim(t *testing.T) {
	cutset := "0123456789"
	s := []byte("453gopher8257")
	println(string(bytes.Trim(s, cutset))) // gopher
}

// 返回字节切片s去除所有左侧cutset子串出现的字符
func TestTrimLeft(t *testing.T) {
	cutset := "0123456789"
	s := []byte("453gopher8257")
	println(string(bytes.TrimLeft(s, cutset))) // gopher8257
}

// 返回字节切片s去除所有右侧cutset子串出现的字符
func TestTrimRight(t *testing.T) {
	cutset := "0123456789"
	s := []byte("453gopher8257")
	println(string(bytes.TrimRight(s, cutset))) // 453gopher
}

// 去除字节切片s的单个前缀prefix字节切片
func TestTrimPrefix(t *testing.T) {
	prefix := []byte("<")
	s := []byte("<<<ok<<<")
	println(string(bytes.TrimPrefix(s, prefix))) // <<ok<<<
}

// 去除字节切片s的单个后缀prefix字节切片
func TestTrimSuffix(t *testing.T) {
	suffix := []byte("<")
	s := []byte("<<<ok<<<")
	println(string(bytes.TrimSuffix(s, suffix))) // <<<ok<<
}

// 去除字节切片s中所有符合f(c)的左右两侧的字符
func TestTrimFunc(t *testing.T) {
	println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter)))
	println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter)))
	println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
	// output:
	// -gopher!
	// "go-gopher!"
	// go-gopher
	// go-gopher
}

// 去除字节切片s中所有符合f(c)的左侧的字符
func TestTrimLeftFunc(t *testing.T) {
	println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter)))
	println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct)))
	println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
	// output:
	// -gopher
	// go-gopher!
	// go-gopher!567
}

// 去除字节切片s中所有符合f(c)的右侧的字符
func TestTrimRightFunc(t *testing.T) {
	println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
	// output:
	// go-
	// go-gopher
	// 1234go-gopher!
}

// 去除字节切片s所有前后空白 (包括换行符等)
func TestTrimSpace(t *testing.T) {
	println("|" + string(bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n"))) + "|")
	// |a lone gopher|
}

// TODO
