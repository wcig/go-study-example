package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	t.Parallel()
	fmt.Println("A")
}

func TestB(t *testing.T) {
	t.Parallel()
	fmt.Println("B")
}

func TestC(t *testing.T) {
	t.Parallel()
	fmt.Println("C")
}

// $ go test -v -parallel 3 parallel_test.go
// === RUN   TestA
// === PAUSE TestA
// === RUN   TestB
// === PAUSE TestB
// === RUN   TestC
// === PAUSE TestC
// === CONT  TestA
// A
// === CONT  TestC
// C
// --- PASS: TestA (0.00s)
// === CONT  TestB
// B
// --- PASS: TestC (0.00s)
// --- PASS: TestB (0.00s)
// PASS
// ok      command-line-arguments  0.103s
