package main

import "fmt"

func main() {
	sum := 100 + 010
	fmt.Println(sum) // 108

	// 二进制
	n1 := 0b10
	fmt.Println(n1) // 2

	// 八进制
	n2 := 010
	fmt.Println(n2) // 8

	// 十六进制
	n3 := 0x10
	fmt.Println(n3) // 16
}
