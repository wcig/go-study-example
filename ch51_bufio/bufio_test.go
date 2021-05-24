package ch51_bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// bufio
// 包bufio实现了缓冲的I/O。
// 它包装了一个io.Reader或io.Writer对象，创建了另一个对象（Reader或Writer），该对象也实现了该接口，但提供了缓冲和一些有关文本I/O的帮助。

// func NewReader(rd io.Reader) *Reader
// 创建默认缓冲区大小的Reader (其缓冲区大小为默认值4096字节)
func TestNewReader(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello"))
	assert.NotNil(t, reader)
}

// func NewReaderSize(rd io.Reader, size int) *Reader
// 返回一个新的Reader，其缓冲区至少具有指定的大小(16)。如果参数io.Reader已经是具有足够大大小的Reader，则它将返回基础Reader。
func TestNewReaderSize(t *testing.T) {
	reader := bufio.NewReaderSize(strings.NewReader("hello"), 16)
	assert.NotNil(t, reader)
}

// func NewWriter(w io.Writer) *Writer
// 返回默认缓冲区大小的Writer (4096字节)
func TestNewWriter(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	num, err := writer.WriteString("hello")
	fmt.Println(num, err) // 5 <nil>
}

// func NewWriterSize(w io.Writer, size int) *Writer
// 返回一个新的Writer，其缓冲区至少具有指定的大小。如果参数io.Writer已经是一个足够大的Writer，它将返回基础Writer
func TestNewWriterSize(t *testing.T) {
	writer := bufio.NewWriterSize(os.Stdout, 16)
	num, err := writer.WriteString("hello")
	fmt.Println(num, err) // 5 <nil>
}

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

func TestScannerWithScanLines(t *testing.T) {
	reader := strings.NewReader("hello world.\nit is ok\n666")
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

// output:
// hello world.
// it is ok
// 666

func TestScannerWithScanWords(t *testing.T) {
	reader := strings.NewReader("hello world. 666")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

// output:
// hello
// world.
// 666

func TestScannerWithScanBytes(t *testing.T) {
	reader := strings.NewReader("hello")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

// output:
// h
// e
// l
// l
// o

func TestScannerWithScanRunes(t *testing.T) {
	reader := strings.NewReader("天气ok")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

// output:
// 天
// 气
// o
// k
