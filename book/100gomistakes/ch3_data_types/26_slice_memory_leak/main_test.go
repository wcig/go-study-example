package main

import (
	"runtime"
	"testing"
)

func TestSliceMemoryLeak(t *testing.T) {
	printAlloc()

	s1 := make([]byte, 1024*1024)
	printAlloc()

	s2 := keepFirstTwoElementsOnly(s1)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(s2)

	// Output:
	// 142 KB
	// 1166 KB
	// 1163 KB
}

func TestSliceMemoryLeakCopy(t *testing.T) {
	printAlloc()

	s1 := make([]byte, 1024*1024)
	printAlloc()

	s2 := keepFirstTwoElementsOnlyCopy(s1)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(s2)

	// Output:
	// 143 KB
	// 1167 KB
	// 140 KB
}
