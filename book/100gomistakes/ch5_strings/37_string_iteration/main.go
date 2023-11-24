package main

import "fmt"

func main() {
	s1 := "tom"
	for i, v := range s1 {
		fmt.Printf("%d, %c\n", i, v)
	}
	fmt.Println("-------------------")
	for i, v := range []rune(s1) {
		fmt.Printf("%d, %c\n", i, v)
	}
	fmt.Println("-------------------")

	s2 := "你好"
	for i, v := range s2 {
		fmt.Printf("%d, %c\n", i, v)
	}
	fmt.Println("-------------------")
	for i, v := range []rune(s2) {
		fmt.Printf("%d, %c\n", i, v)
	}
	fmt.Println("-------------------")

	// Output:
	// 0, t
	// 1, o
	// 2, m
	// -------------------
	// 0, t
	// 1, o
	// 2, m
	// -------------------
	// 0, 你
	// 3, 好
	// -------------------
	// 0, 你
	// 1, 好
	// -------------------
}
