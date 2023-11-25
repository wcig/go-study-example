package main

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	fmt.Println("AA")
}

func TestBB(t *testing.T) {
	fmt.Println("BB")
}

func TestCC(t *testing.T) {
	fmt.Println("CC")
}

// $ go test -v -shuffle=on shuffle_test.go
// -test.shuffle 1700918311507776000
// === RUN   TestCC
// CC
// --- PASS: TestCC (0.00s)
// === RUN   TestBB
// BB
// --- PASS: TestBB (0.00s)
// === RUN   TestAA
// AA
// --- PASS: TestAA (0.00s)
// PASS
// ok      command-line-arguments  0.333s
