package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("index=%d, value=%s\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("value=%s\n", v)
	}

	// Output:
	// index=0, value=a
	// index=1, value=b
	// index=2, value=c
	// value=a
	// value=b
	// value=c
}
