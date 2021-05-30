package ch4_bytes

import (
	"bytes"
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// bytes包

/* -----------------------------字节切片比较-------------------------------------- */

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

// 不区分大小写是否相等
func TestEqualFold(t *testing.T) {
	println(bytes.EqualFold([]byte("Go"), []byte("go"))) // true
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

/* -----------------------------字节切片包含,计算-------------------------------------- */

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

// s是否有指定prefix前缀
func TestHasPrefix(t *testing.T) {
	prefix := []byte("hello")
	s1 := []byte("hello ok")
	s2 := []byte("ok")

	fmt.Println(bytes.HasPrefix(s1, prefix)) // true
	fmt.Println(bytes.HasPrefix(s2, prefix)) // false
}

// s是否有指定prefix前缀
func TestHasSuffix(t *testing.T) {
	suffix := []byte("hello")
	s1 := []byte("ok hello")
	s2 := []byte("ok")

	fmt.Println(bytes.HasSuffix(s1, suffix)) // true
	fmt.Println(bytes.HasSuffix(s2, suffix)) // false
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

/* -----------------------------字节切片修剪-------------------------------------- */

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

/* -----------------------------字节切片大小写转换-------------------------------------- */

// 转小写
func TestToLower(t *testing.T) {
	println(string(bytes.ToLower([]byte("Gopher")))) // gopher
}

// 转大写
func TestToUpper(t *testing.T) {
	println(string(bytes.ToUpper([]byte("Gopher")))) // GOPHER
}

// 转小写+特殊字符处理
func TestToLowerSpecial(t *testing.T) {
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	println("Original : " + string(str))
	println("ToLower : " + string(totitle))
	// output:
	// Original : AHOJ VÝVOJÁRİ GOLANG
	// ToLower : ahoj vývojári golang
}

// 转大写+特殊字符处理
func TestToUpperSpecial(t *testing.T) {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToUpperSpecial(unicode.AzeriCase, str)
	println("Original : " + string(str))
	println("ToUpper : " + string(totitle))
	// output:
	// Original : ahoj vývojári golang
	// ToUpper : AHOJ VÝVOJÁRİ GOLANG
}

// 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有 Unicode 字母都映射到它们的标题大小写。
func TestToTitle(t *testing.T) {
	println(string(bytes.ToTitle([]byte("loud noises"))))
	println(string(bytes.ToTitle([]byte("хлеб"))))
	// output:
	// LOUD NOISES
	// ХЛЕБ
}

// 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中每个字节代表无效的 UTF-8 替换为替换中的字节，该字节可能为空。
func TestToValidUTF8(t *testing.T) {
	println(string(bytes.ToValidUTF8([]byte("hello你好"), []byte("ok")))) // hello你好
}

/* -----------------------------字节切片分割合并-------------------------------------- */

// 以指定字节切片合并
func TestJoin(t *testing.T) {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	println(string(bytes.Join(s, []byte(", ")))) // foo, bar, baz
}

// 以seq字节切片分割,如果seq为空则划分每个utf8字符 (等价于SplitN() n=-1)
func TestSplit(t *testing.T) {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c,好"), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
	// output:
	// ["a" "b" "c"]
	// ["a" "," "b" "," "c" "," "好"]
	// ["" "man " "plan " "canal panama"]
	// [" " "x" "y" "z" " "]
	// [""]
}

// 基本同Split()一样,区别: n>0返回n个子切片,超过的不再划分; n==0返回空切片; n<0返回所有子切片
func TestSplitN(t *testing.T) {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	// output:
	// ["a" "b,c"]
	// [] (nil = true)
}

// 切片划分
func TestSplitAfter(t *testing.T) {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c,好"), []byte("")))
	// output:
	// ["a," "b," "c"]
	// ["a" "," "b" "," "c" "," "好"]
}

// 切片划分
func TestSplitAfterN(t *testing.T) {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2)) // ["a," "b,c"]
}

// 按单个或多个空白符划分
func TestFields(t *testing.T) {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   "))) // Fields are: ["foo" "bar" "baz"]
}

// 按指定函数划分
func TestFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f)) // Fields are: ["foo1" "bar2" "baz3"]
}

// 将 s 解释为一系列 UTF-8 编码的代码点。 它返回一段相当于 s 的符文（Unicode 代码点）。
func TestRunes(t *testing.T) {
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}
	// output:
	// U+0067 'g'
	// U+006F 'o'
	// U+0020 ' '
	// U+0067 'g'
	// U+006F 'o'
	// U+0070 'p'
	// U+0068 'h'
	// U+0065 'e'
	// U+0072 'r'
}

/* -----------------------------字节切片替换-------------------------------------- */

// 重复字节切片指定次数
func TestRepeat(t *testing.T) {
	println(string(bytes.Repeat([]byte("ab"), 3))) // ababab
}

// 替换字节切片最多指定个数,-1表示所有
func TestReplace(t *testing.T) {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
	// output:
	// oinky oinky oink
	// moo moo moo
}

// 替换所有字节切片
func TestReplaceAll(t *testing.T) {
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo"))) // moo moo moo
}

// 返回字节切片 s 的副本，其中所有字符都根据映射函数进行了修改。 如果映射返回负值，则字符从字节切片中删除而没有替换。 s 中的字符和输出被解释为 UTF-8 编码的代码点。
func TestMap(t *testing.T) {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
	// 'Gjnf oevyyvt naq gur fyvgul tbcure...
}

// 根据输入字节切片初始化一buffer (buffer集成多个reader、writer接口)
func TestNewBuffer(t *testing.T) {
	buf1 := bytes.NewBuffer([]byte("hello"))
	assert.NotNil(t, buf1)

	buf2 := bytes.NewBufferString("hello")
	assert.NotNil(t, buf2)

	var buf3 bytes.Buffer
	assert.NotNil(t, buf3)
}

// bytes.Buffer方法
func TestBufferMethod(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hello"))
	fmt.Println(buf.String())

	n, err := buf.Write([]byte(" world"))
	fmt.Println(n, err, buf.String())

	s := make([]byte, 2)
	nn, err := buf.Read(s)
	fmt.Println(nn, err, buf.String())
	// output:
	// hello
	// 6 <nil> hello world
	// 2 <nil> llo world
}

// 创建bytes.Reader (buffer集成多个reader接口)
func TestNewReader(t *testing.T) {
	reader := bytes.NewReader([]byte("hello"))
	s1 := make([]byte, 5)
	n, err := reader.Read(s1)
	fmt.Println(n, err)
	fmt.Println(string(s1))
	fmt.Println(reader.Size(), reader.Len())

	reader = bytes.NewReader([]byte("hello world."))
	fmt.Println(reader.Size(), reader.Len())
	var buf bytes.Buffer
	nn, err := reader.WriteTo(&buf)
	fmt.Println(nn, err)
	fmt.Println(buf.String())
	// output:
	// 5 <nil>
	// hello
	// 5 0
	// 12 12
	// 12 <nil>
	// hello world.
}
