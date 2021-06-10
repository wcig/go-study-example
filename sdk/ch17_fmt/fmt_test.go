package ch17_fmt

import (
	"fmt"
	"testing"
)

// fmt

func TestPrintln(t *testing.T) {
	n, err := fmt.Println("ok")
	fmt.Println(n, err)
}
