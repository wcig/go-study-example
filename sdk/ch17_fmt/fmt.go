package main

import (
	"fmt"
	"io"
)

// test scan
func main() {
	var (
		name string
		age  int
	)
	n, err := fmt.Scan(&name, &age)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(name, age)
	// 终端输入：tom 20
	// output:
	// 2
	// tom 20
}
