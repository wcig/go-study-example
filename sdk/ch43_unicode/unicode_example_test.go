package ch43_unicode

import (
	"fmt"
	"testing"
	"unicode"
)

func TestSimpleFold(t *testing.T) {
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'
	// output:
	// U+0061 'a'
	// U+0041 'A'
	// U+006B 'k'
	// U+212A 'K'
	// U+004B 'K'
	// U+0031 '1'
}

func TestTo(t *testing.T) {
	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))

	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))
	// output:
	// U+0047 'G'
	// U+0067 'g'
	// U+0047 'G'
	// U+0047 'G'
	// U+0067 'g'
	// U+0047 'G'
}

func TestToLower(t *testing.T) {
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG)) // U+0067 'g'
}

func TestToTitle(t *testing.T) {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG)) // U+0047 'G'
}

func TestToUpper(t *testing.T) {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG)) // U+0047 'G'
}
