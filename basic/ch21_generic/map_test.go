package ch21_generic

import (
	"fmt"
	"testing"
)

// 泛型map
type Map[K string, V int | float64] map[K]V

func TestMap(t *testing.T) {
	m1 := Map[string, int]{"a": 1, "b": 2, "c": 3}
	m1["d"] = 4
	fmt.Println(m1)

	m2 := Map[string, float64]{"aa": 1.1, "bb": 2.2, "cc": 3.3}
	m2["dd"] = 4.4
	fmt.Println(m2)
}
