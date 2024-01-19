package main

import (
	"fmt"
	"os"
	"strings"
)

// 修改 echo 程序输出 os.Args[0]，即命令的名字

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
