package main

import (
	"fmt"
)

// bad
func listing1() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func() {
			fmt.Print(i)
		}()
	}
	// Output:
	// 333
}

// good
func listing2() {
	s := []int{1, 2, 3}

	for _, i := range s {
		val := i
		go func() {
			fmt.Print(val)
		}()
	}
}

// good
func listing3() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func(val int) {
			fmt.Print(val)
		}(i)
	}
}
