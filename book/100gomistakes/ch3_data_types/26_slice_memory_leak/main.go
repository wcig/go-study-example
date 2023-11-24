package main

import (
	"fmt"
	"runtime"
)

func keepFirstTwoElementsOnly(s []byte) []byte {
	return s[:5]
}

func keepFirstTwoElementsOnlyCopy(s []byte) []byte {
	dst := make([]byte, 5)
	copy(dst, s)
	return dst
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
