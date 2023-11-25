package main

import (
	"fmt"
	"time"
)

func main() {
	example1()

	example2()
}

// good
func example1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()

	panic("foo")
}

// bad
func example2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()

	go func() {
		panic("bar")
	}()
	time.Sleep(time.Second)
}
