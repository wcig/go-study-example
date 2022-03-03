package array

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var arr [5]int
	fmt.Printf("array size: %d, value: %v\n", len(arr), arr)

	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	fmt.Printf("array size: %d, value: %v\n", len(arr), arr)
	fmt.Println("array index 0 value:", arr[0])
	fmt.Println("array index 3 value:", arr[3])
	// Output:
	// array size: 5, value: [0 0 0 0 0]
	// array size: 5, value: [0 1 2 3 4]
	// array index 0 value: 0
	// array index 3 value: 3
}
