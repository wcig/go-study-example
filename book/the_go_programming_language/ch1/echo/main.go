package main

import (
	"fmt"
	"os"
	"strings"
)

// echo: 输出命令行参数

func main() {
	echo1()
	// echo2()
	// echo3()
}

func echo1() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}

func echo2() {
	var s, step string
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
