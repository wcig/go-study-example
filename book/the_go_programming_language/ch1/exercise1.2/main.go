package main

import (
	"fmt"
	"os"
)

// 修改 echo 程序，输出参数的索引和值，每行一个

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
