package utf8

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestDecodeLastRune(t *testing.T) {
	b := []byte("Hello 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %d\n", r, size)

		b = b[:len(b)-size]
	}
	// output:
	// 界 3
	// 世 3
	//  1
	// o 1
	// l 1
	// l 1
	// e 1
	// H 1
}

func TestDecodeLastRuneInString(t *testing.T) {
	s := "Hello 世界"
	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		fmt.Printf("%c %d\n", r, size)

		s = s[:len(s)-size]
	}
	// output:
	// 界 3
	// 世 3
	//  1
	// o 1
	// l 1
	// l 1
	// e 1
	// H 1
}

func TestDecodeRune(t *testing.T) {
	b := []byte("Hello 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %d\n", r, size)

		b = b[size:]
	}
	// output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	//  1
	// 世 3
	// 界 3
}

func TestDecodeRuneLnString(t *testing.T) {
	s := "Hello 世界"
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("%c %d\n", r, size)

		s = s[size:]
	}
	// output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	//  1
	// 世 3
	// 界 3
}

func TestEncodeRune(t *testing.T) {
	r := '好'
	b1 := make([]byte, 3)
	n1 := utf8.EncodeRune(b1, r)
	fmt.Println(n1) // 3

	// b2 := make([]byte, 2)
	// n2 := utf8.EncodeRune(b2, r)
	// fmt.Println(n2) // panic
}

func TestFullRune(t *testing.T) {
	b := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(b))
	fmt.Println(utf8.FullRune(b[:2]))
	// output:
	// true
	// false
}

func TestFullRuneInString(t *testing.T) {
	s := "世"
	fmt.Println(utf8.FullRuneInString(s))
	fmt.Println(utf8.FullRuneInString(s[:2]))
	// output:
	// true
	// false
}

func TestRuneCount(t *testing.T) {
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
	// output:
	// bytes = 13
	// runes = 9
}

func TestRuneCountInString(t *testing.T) {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))
	// output:
	// bytes = 13
	// runes = 9
}

func TestRuneLen(t *testing.T) {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('世'))
	// output:
	// 1
	// 3
}

func TestRuneStart(t *testing.T) {
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
	// output:
	// true
	// true
	// false
}

func TestValid(t *testing.T) {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
	// output:
	// true
	// false
}

func TestValidRune(t *testing.T) {
	valid := 'a'
	invalid := rune(0xfffffff)

	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
	// output:
	// true
	// false
}

func TestValidString(t *testing.T) {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
	// output:
	// true
	// false
}
