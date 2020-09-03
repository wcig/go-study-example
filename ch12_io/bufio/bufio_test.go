package bufio

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

// ReadSlice
func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.\nIt is ok"))

	line1, _ := reader.ReadSlice('\n')
	fmt.Println("1st line:", string(line1))

	line2, _ := reader.ReadSlice('\n')
	fmt.Println("2nd line:", string(line2))
	fmt.Println(string(line1))
}

// 1st line: hello world.
//
// 2nd line: It is ok
// It is okrld.
//
