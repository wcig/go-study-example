package ch2_package

import (
	"fmt"
	"go-app/ch2_package/hello"

	// _ "go-app/ch2_package/hello"
	"testing"
)

// 下划线引入
// 1.引入"go-app/ch2_package/hello", 则hello包下的所有go文件的init函数都会执行, 同时hello包下其他函数可调用
// 2.引入_ "go-app/ch2_package/hello", 则hello包下的所有go文件的init函数都会执行, 同时hello包下其他函数不可调用
func TestUnderlinePackage(t *testing.T) {
	fmt.Println("underline package import")

	hello.Print()
}
