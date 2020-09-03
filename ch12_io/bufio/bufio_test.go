package bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

// ReadSlice
func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadSlice('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadSlice('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
}

// output:
// line1: hello world.|
// line2: It is ok
// line1: It is okrld.|

// ReadBytes
func TestReadBytes(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadBytes('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadBytes('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
}

// output:
// line1: hello world.|
// line2: It is ok
// line1: hello world.|

// ReadString
func TestReadString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadString('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadString('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
}

// output:
// line1: hello world.|
// line2: It is ok
// line1: hello world.|

// ReadLine
func TestReadLine(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.\nIt is ok"))
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(line), isPrefix)
	}
}

// output:
// hello world. false
// It is ok false

// WriteString
func TestWriteString(t *testing.T) {
	rb := &bytes.Buffer{}
	writer := bufio.NewWriter(rb)
	n, err := writer.WriteString("hello world.")
	fmt.Println(n, err) // 12 <nil>

	writer.Flush()
	fmt.Println(rb.String()) // hello world.
}

// Test for golang.org/issue/5947
func TestWriterReadFromWhileFull(t *testing.T) {
	buf := new(bytes.Buffer)
	w := bufio.NewWriterSize(buf, 10)

	// Fill buffer exactly.
	n, err := w.Write([]byte("0123456789"))
	if n != 10 || err != nil {
		t.Fatalf("Write returned (%v, %v), want (10, nil)", n, err)
	}

	// Use ReadFrom to read in some data.
	n2, err := w.ReadFrom(strings.NewReader("abcdef"))
	if n2 != 6 || err != nil {
		t.Fatalf("ReadFrom returned (%v, %v), want (6, nil)", n2, err)
	}

	w.Flush()
	fmt.Println(buf.String()) // 0123456789abcdef
}
