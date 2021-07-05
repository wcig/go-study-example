package scanner

import (
	"fmt"
	"strings"
	"testing"
	"text/scanner"
	"unicode"
)

func TestScanner(t *testing.T) {
	const src = `
// This is scanned code.
if a > 10 {
    someParsable = text
}`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	// output:
	// example:3:1: if
	//example:3:4: a
	//example:3:6: >
	//example:3:8: 10
	//example:3:11: {
	//example:4:5: someParsable
	//example:4:18: =
	//example:4:20: text
	//example:5:1: }
}

func TestIsIdentRune(t *testing.T) {
	const src = "%var1 var2%"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	fmt.Println()
	s.Init(strings.NewReader(src))
	s.Filename = "percent"

	// treat leading '%' as part of an identifier
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	// output:
	// default:1:1: %
	//default:1:2: var1
	//default:1:7: var2
	//default:1:11: %
	//
	//percent:1:1: %var1
	//percent:1:7: var2
	//percent:1:11: %
}

func TestMode(t *testing.T) {
	const src = `
    // Comment begins at column 5.

This line should not be included in the output.

/*
This multiline comment
should be extracted in
its entirety.
*/
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if strings.HasPrefix(txt, "//") || strings.HasPrefix(txt, "/*") {
			fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}
	// output:
	// comments:2:5: // Comment begins at column 5.
	// comments:6:1: /*
	// This multiline comment
	// should be extracted in
	// its entirety.
	// */
}

func TestWhitespace(t *testing.T) {
	// tab-separated values
	const src = `aa	ab	ac	ad
ba	bb	bc	bd
ca	cb	cc	cd
da	db	dc	dd`

	var (
		col, row int
		s        scanner.Scanner
		tsv      [4][4]string // large enough for example above
	)
	s.Init(strings.NewReader(src))
	s.Whitespace ^= 1<<'\t' | 1<<'\n' // don't skip tabs and new lines

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '\n':
			row++
			col = 0
		case '\t':
			col++
		default:
			tsv[row][col] = s.TokenText()
		}
	}

	fmt.Print(tsv) // [[aa ab ac ad] [ba bb bc bd] [ca cb cc cd] [da db dc dd]]
}
