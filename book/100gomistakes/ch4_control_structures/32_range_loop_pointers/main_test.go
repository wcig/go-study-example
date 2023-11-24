package main

import (
	"fmt"
	"testing"
)

func TestRangeLoopPointers(t *testing.T) {
	{
		// bad
		m := make(map[int]*int)
		s := []int{1, 2, 3}
		for _, v := range s {
			m[v] = &v
		}
		for k, v := range m {
			fmt.Println(">> range map-1:", k, *v)
		}
	}

	{
		// good
		m := make(map[int]*int)
		s := []int{1, 2, 3}
		for i := range s {
			m[s[i]] = &s[i]
		}
		for k, v := range m {
			fmt.Println(">> range map-2:", k, *v)
		}
	}

	// Output:
	// >> range map-1: 3 3
	// >> range map-1: 1 3
	// >> range map-1: 2 3
	// >> range map-2: 1 1
	// >> range map-2: 2 2
	// >> range map-2: 3 3
}
