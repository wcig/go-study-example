package ch21_generic

import (
	"fmt"
	"testing"
)

// 泛型slice
type Slice[T int | float64 | string] []T

func TestSlice(t *testing.T) {
	var s1 Slice[int] = []int{1, 2, 3}
	var s2 Slice[float64] = []float64{1.1, 2.2, 3.3}
	var s3 Slice[string] = []string{"a", "b", "c"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// error
	// var s4 Slice[int32] = []int32{1, 2, 3}
	// s1 = append(s1, 1.1)
	// s2 = append(s2, "11")
	// s3 = append(s3, 1)
}
