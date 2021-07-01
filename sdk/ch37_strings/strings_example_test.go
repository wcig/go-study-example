package ch37_strings

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"
)

func TestCompare(t *testing.T) {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	// output:
	// -1
	// 0
	// 1
}

func TestContains(t *testing.T) {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	// output:
	// true
	// false
	// true
	// true
}

func TestContainsAny(t *testing.T) {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	// output:
	// false
	// true
	// true
	// true
	// false
	// false
}

func TestContainsRune(t *testing.T) {
	// a - 97
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	// output:
	// true
	// false
}

func TestCount(t *testing.T) {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", ""))
	// output:
	// 3
	// 5
}

func TestEqualFold(t *testing.T) {
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.EqualFold("go", "go"))
	// output:
	// true
	// true
}

func TestFields(t *testing.T) {
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	// Fields are: ["foo" "bar" "baz"]
}

func TestFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	// Fields are: ["foo1" "bar2" "baz3"]
}

func TestHasPrefix(t *testing.T) {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	// output:
	// true
	// false
	// true
}

func TestSuffix(t *testing.T) {
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	// output:
	// true
	// false
	// false
	// true
}

func TestIndex(t *testing.T) {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	// output:
	// 4
	// -1
}

func TestIndexAny(t *testing.T) {
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	// output:
	// 2
	// -1
}

func TestIndexByte(t *testing.T) {
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
	// output:
	// 0
	// 3
	// -1
}

func TestIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
	// output:
	// 7
	// -1
}

func TestIndexRun(t *testing.T) {
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	// output:
	// 4
	// -1
}

func TestJoin(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
}

func TestLastIndex(t *testing.T) {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	// output:
	// 0
	// 3
	// -1
}

func TestLastIndexAny(t *testing.T) {
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
	// output:
	// 4
	// 8
	// -1
}

func TestLastIndexByte(t *testing.T) {
	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
	// output:
	// 10
	// 8
	// -1
}

func TestLastIndexFunc(t *testing.T) {
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
	// output:
	// 5
	// 2
	// -1
}

func TestMap(t *testing.T) {
	mapping := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 32
		}
		return r
	}
	fmt.Println(strings.Map(mapping, "This is a cat.")) // THIS IS A CAT.
}

func TestRepeat(t *testing.T) {
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana
}

func TestReplace(t *testing.T) {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// output:
	// oinky oinky oink
	// moo moo moo
}

func TestReplaceAll(t *testing.T) {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) // moo moo moo
}

func TestSplit(t *testing.T) {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	// output:
	// ["a" "b" "c"]
	// ["" "man " "plan " "canal panama"]
	// [" " "x" "y" "z" " "]
	// [""]
}

func TestSplitAfter(t *testing.T) {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
}

func TestSplitAfterN(t *testing.T) {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
}

func TestSplitN(t *testing.T) {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	// output:
	// ["a" "b,c"]
	// [] (nil = true)
}

func TestTitle(t *testing.T) {
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))
	fmt.Println(strings.Title("хлеб"))
	// output:
	// Her Royal Highness
	// Loud Noises
	// Хлеб
}

func TestToLower(t *testing.T) {
	fmt.Println(strings.ToLower("Gopher")) // gopher
}

func TestToLowerSpecial(t *testing.T) {
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) // önnek iş
}

func TestToTitle(t *testing.T) {
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
	// output:
	// HER ROYAL HIGHNESS
	// LOUD NOISES
	// ХЛЕБ
}

func TestToTitleSpecial(t *testing.T) {
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	// DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
}

func TestToUpper(t *testing.T) {
	fmt.Println(strings.ToUpper("Gopher")) // GOPHER
}

func TestToUpperSpecial(t *testing.T) {
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş")) // ÖRNEK İŞ
}

func TestTrim(t *testing.T) {
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.Trim("aaaHello, Gophersbbb", "ab"))
	// output:
	// Hello, Gophers
	// Hello, Gophers
}

func TestTrimFunc(t *testing.T) {
	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	// Hello, Gophers
}

func TestTrimLeft(t *testing.T) {
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimLeft("aaaHello, Gophersbbb", "ab"))
	// output:
	// Hello, Gophers!!!
	// Hello, Gophersbbb
}

func TestTrimLeftFunc(t *testing.T) {
	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	// Hello, Gophers!!!
}

func TestTrimPrefix(t *testing.T) {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Println(s) // Gophers!!!
}

func TestTrimRight(t *testing.T) {
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) // ¡¡¡Hello, Gophers
}

func TestTrimRightFunc(t *testing.T) {
	fmt.Println(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	// ¡¡¡Hello, Gophers
}

func TestTrimSpace(t *testing.T) {
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers
}

func TestTrimSuffix(t *testing.T) {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Println(s) // ¡¡¡Hello
}

func TestTypeBuilder(t *testing.T) {
	var b strings.Builder

	fmt.Println("---init---")
	fmt.Println("cap:", b.Cap())
	fmt.Println("len:", b.Len())

	n, err := b.Write([]byte("ok"))
	fmt.Println(n, err)
	err = b.WriteByte('-')
	fmt.Println(err)
	n, err = b.WriteRune('好')
	fmt.Println(n, err)
	n, err = b.WriteString("666")
	fmt.Println(n, err)

	fmt.Println("---after write---")
	fmt.Println("content:", b.String())
	fmt.Println("cap:", b.Cap())
	fmt.Println("len:", b.Len())

	fmt.Println("---after reset---")
	b.Reset()
	fmt.Println("content:", b.String())
	fmt.Println("cap:", b.Cap())
	fmt.Println("len:", b.Len())

	fmt.Println("---after grow---")
	b.Grow(5)
	fmt.Println("cap:", b.Cap())
	fmt.Println("len:", b.Len())
	// output:
	// ---init---
	// cap: 0
	// len: 0
	// 2 <nil>
	// <nil>
	// 3 <nil>
	// 3 <nil>
	// ---after write---
	// content: ok-好666
	// cap: 16
	// len: 9
	// ---after reset---
	// content:
	// cap: 0
	// len: 0
	// ---after grow---
	// cap: 5
	// len: 0
}

func TestTypeReader(t *testing.T) {
	r := strings.NewReader("ok\n")
	fmt.Println("len:", r.Len())
	fmt.Println("size:", r.Size())
	n, err := r.WriteTo(os.Stdout)
	fmt.Println(n, err)
	// output:
	// len: 3
	// size: 3
	// ok
	// 3 <nil>
}

func TestReplacer(t *testing.T) {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	n, err := r.WriteString(os.Stdout, "<ok>\n")
	fmt.Println(n, err)
	// output:
	// This is &lt;b&gt;HTML&lt;/b&gt;!
	// &lt;ok&gt;
	// 11 <nil>
}
