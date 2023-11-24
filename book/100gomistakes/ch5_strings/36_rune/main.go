package main

import (
	"fmt"
	"unicode/utf8"
)

// type byte = uint8
// type rune = int32
func main() {
	s1 := "hello"
	fmt.Println(len(s1), len([]rune(s1)), utf8.RuneCountInString(s1))

	s2 := "好"
	fmt.Println(len(s2), len([]rune(s2)), utf8.RuneCountInString(s2))

	// Output:
	// 5 5 5
	// 3 1 1
}
