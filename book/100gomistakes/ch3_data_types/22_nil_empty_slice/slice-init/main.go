package main

import (
	"fmt"
)

func main() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	// Output:
	// 1: empty=true   nil=true
	// 2: empty=true   nil=true
	// 3: empty=true   nil=false
	// 4: empty=true   nil=false
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}
