package main

import "fmt"

func main() {
	{
		// bad
		s1 := []int{1, 2, 3, 4}
		_ = append(s1, 10)
		fmt.Println(s1) // [1 2 3 4]

		// good
		s2 := []int{1, 2, 3, 4}
		s2 = append(s2, 10)
		fmt.Println(s2) // [1 2 3 4 10]
	}

	{
		s1 := []int{1, 2, 3}
		s2 := s1[1:2]
		fmt.Println(s2) // [2]
		s3 := append(s2, 10)
		fmt.Println(s1, s2, s3) // [1 2 10] [2] [2 10]
	}
}
