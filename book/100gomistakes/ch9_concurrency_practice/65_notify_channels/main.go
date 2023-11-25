package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bad
	ch1 := make(chan bool, 1)
	go func() {
		ch1 <- true
	}()
	v1 := <-ch1
	fmt.Println(v1, unsafe.Sizeof(v1)) // true 1

	// good
	ch2 := make(chan struct{}, 1)
	go func() {
		ch2 <- struct{}{}
	}()
	v2 := <-ch2
	fmt.Println(v2, unsafe.Sizeof(v2)) // {} 0
}
