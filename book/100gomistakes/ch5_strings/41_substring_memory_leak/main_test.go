package main

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// bad
func TestSubString1(t *testing.T) {
	printAlloc()

	s := genLongStr()
	printAlloc()

	sub := s[:36]

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(sub)

	// Output:
	// 141 KB
	// 1165 KB
	// 1163 KB
}

// good
func TestSubString2(t *testing.T) {
	printAlloc()

	s := genLongStr()
	printAlloc()

	sub := string([]byte(s[:36]))

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(sub)

	// Output:
	// 142 KB
	// 1166 KB
	// 139 KB
}

// good
func TestSubString3(t *testing.T) {
	printAlloc()

	s := genLongStr()
	printAlloc()

	sub := strings.Clone(s[:36])

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(sub)

	// Output:
	// 142 KB
	// 1166 KB
	// 139 KB
}

func genLongStr() string {
	const size = 1024 * 1024
	sb := strings.Builder{}
	sb.Grow(size)
	for i := 0; i < size; i++ {
		sb.WriteByte('a')
	}
	return sb.String()
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
