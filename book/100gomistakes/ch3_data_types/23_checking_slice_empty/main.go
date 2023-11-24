package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int(nil)
	s3 := []int{}
	s4 := make([]int, 0)
	fmt.Println(isEmpty(s1), isEmpty(s2), isEmpty(s3), isEmpty(s4)) // true true true true
}

func isEmpty(s []int) bool {
	return len(s) == 0
}
