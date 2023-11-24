package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMapMemoryLeak128Bytes(t *testing.T) {
	// Init
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = rand128Bytes()
	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)

	// Output:
	// 0 MB
	// 461 MB
	// 293 MB
}

func TestMapMemoryLeak129Bytes(t *testing.T) {
	// Init
	n := 1_000_000
	m := make(map[int][129]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = rand129Bytes()
	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)

	// Output:
	// 0 MB
	// 197 MB
	// 38 MB
}

func TestMapMemoryLeak128BytesPointer(t *testing.T) {
	// Init
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = rand128BytesPointer()
	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)

	// Output:
	// 0 MB
	// 182 MB
	// 38 MB
}

func rand128Bytes() [128]byte {
	return [128]byte{}
}

func rand129Bytes() [129]byte {
	return [129]byte{}
}

func rand128BytesPointer() *[128]byte {
	return &[128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
