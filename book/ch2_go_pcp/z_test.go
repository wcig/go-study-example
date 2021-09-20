package ch2_go_pcp

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	fmt.Println("test1...")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func test2() {
	fmt.Println("test2...")
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func test3() {
	fmt.Println("test3...")
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func test4() {
	fmt.Println("test4...")
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(0)
	}
}
