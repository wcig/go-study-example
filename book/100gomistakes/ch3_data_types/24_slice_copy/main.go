package main

import "fmt"

func main() {
	{
		// bad
		s1 := []int{1, 2, 3}
		s2 := make([]int, len(s1)+1)
		copy(s2, s1)
		fmt.Println(s2)
	}

	{
		// bad
		s1 := []int{1, 2, 3}
		s2 := make([]int, len(s1)-1)
		copy(s2, s1)
		fmt.Println(s2)
	}

	{
		// bad
		s1 := []int{1, 2, 3}
		var s2 []int
		copy(s2, s1)
		fmt.Println(s2)
	}

	{
		// good
		s1 := []int{1, 2, 3}
		s2 := make([]int, len(s1))
		copy(s2, s1)
		fmt.Println(s2)
	}

	// Output:
	// [1 2 3 0]
	// [1 2]
	// []
	// [1 2 3]
}
