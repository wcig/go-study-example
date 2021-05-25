package main

import (
	"fmt"
)

var num = a()

func init() {
	fmt.Println("calling init() 1")
}

func init() {
	fmt.Println("calling init() 2")
}

func a() int {
	fmt.Println("calling a()")
	return 0
}

// Go 初始化顺序:(1)引入的包;(2)当前包的变量常量;(3)当前包的init函数;(4)main函数.
func main() {
	fmt.Println("calling main()")
}

// 输出:
// calling a()
// calling init() 1
// calling init() 2
// calling main()
