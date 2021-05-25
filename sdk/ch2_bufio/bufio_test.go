package ch2_bufio

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

// func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
// 读取直到输入中第一次出现delim为止，并返回一个指向缓冲区中字节的切片。字节在下一次读取时不再有效。(由于前一次读取会被后面覆盖,不建议使用)
func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadSlice('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadSlice('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
	// output:
	// line1: hello world.|
	// line2: It is ok
	// line1: It is okrld.|
}

// func (b *Reader) ReadBytes(delim byte) ([]byte, error)
// 读取直到输入中第一次出现delim为止，并返回一个切片，该切片包含直到定界符（包括定界符）的数据。(不会覆盖前一次读取,简单使用的话建议使用Scanner)
func TestReadBytes(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadBytes('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadBytes('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
	// output:
	// line1: hello world.|
	// line2: It is ok
	// line1: hello world.|
}

// func (b *Reader) ReadString(delim byte) (string, error)
// 读取直到输入中第一次出现delim为止，返回一个字符串，该字符串包含直到定界符（包括定界符）的数据。(不会覆盖前一次读取,简单使用的话建议使用Scanner)
func TestReadString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadString('|')
	fmt.Println("line1:", string(line1))

	line2, _ := reader.ReadString('|')
	fmt.Println("line2:", string(line2))
	fmt.Println("line1:", string(line1))
	// output:
	// line1: hello world.|
	// line2: It is ok
	// line1: hello world.|
}

// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
// 是一个低级别的行读取原语。 大多数调用者应改用ReadBytes('\n')或ReadString('\n')或使用扫描仪。
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
	// output:
	// hello world. false
	// It is ok false
}

// func (b *Reader) Read(p []byte) (n int, err error)
// 读取数据并写入p,返回写入p的字节数和错误.可能存在n<len(p)情况,如需读取准确的len(p)需使用io.ReadFull(b,p)
func TestRead(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))
	buf := make([]byte, 100)
	n, err := reader.Read(buf)
	fmt.Println(n, err)           // 21 <nil>
	fmt.Println(string(buf[0:n])) // hello world.|It is ok
}

// func (b *Reader) ReadByte() (byte, error)
// 读取并返回一个字节,如果没有字节将返回错误 (建议使用bufio.Scanner更方便)
func TestReadByte(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}
	// output:
	// h
	// e
	// l
	// l
	// o
}

// func (b *Reader) ReadRune() (r rune, size int, err error)
// 读取单个UTF-8编码的Unicode字符，并返回符文及其大小（以字节为单位）。 (建议使用bufio.Scanner更方便)
func TestReadRune(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	for {
		r, n, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Println(r, n)
	}
	// output:
	// h
	// e
	// l
	// l
	// o
}

// func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
// 写入数据到w
func TestWriteTo(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))
	var buf bytes.Buffer
	num, err := reader.WriteTo(&buf)
	fmt.Println(num, err)     // 21 <nil>
	fmt.Println(buf.String()) // hello world.|It is ok
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

// func (b *Writer) WriteString(s string) (int, error)
// 写入一个字符串。 它返回写入的字节数。 如果计数小于len，它还会返回一个错误，解释为什么写时间短。
func TestWriteString(t *testing.T) {
	rb := &bytes.Buffer{}
	writer := bufio.NewWriter(rb)
	n, err := writer.WriteString("hello world.")
	fmt.Println(n, err)      // 12 <nil>
	fmt.Println(rb.String()) //
}

// func (b *Writer) Flush() error
// 缓冲区数据写入底层io.Writer中
func TestFlush(t *testing.T) {
	rb := &bytes.Buffer{}
	writer := bufio.NewWriter(rb)
	_, err := writer.WriteString("hello world.")
	assert.Nil(t, err)

	err = writer.Flush()
	assert.Nil(t, err)
	fmt.Println(rb.String()) // hello world.
}

// func NewReadWriter(r *Reader, w *Writer) *ReadWriter
// 分配一个新的ReadWriter，该ReadWriter调度到r和w。(使用较少)
func TestNewReadWriter(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	readWriter := bufio.NewReadWriter(reader, writer)
	n, err := readWriter.WriteString("hello")
	fmt.Println(n, err)
	_ = readWriter.Flush()
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

// func NewScanner(r io.Reader) *Scanner
// 返回一个新的Scanner以从r读取。 拆分功能默认为ScanLines。
// 分割函数bufio.ScanLines: 按换行符'\n'分割
func TestScannerWithScanLines(t *testing.T) {
	reader := strings.NewReader("hello world.\nit is ok\n666")
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	// output:
	// hello world.
	// it is ok
	// 666
}

// 分割函数bufio.ScanWords: 按单词分割
func TestScannerWithScanWords(t *testing.T) {
	reader := strings.NewReader("hello world. 666")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	// output:
	// hello
	// world.
	// 666
}

// 分割函数bufio.ScanBytes: 按字节分割
func TestScannerWithScanBytes(t *testing.T) {
	reader := strings.NewReader("hello")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	// output:
	// h
	// e
	// l
	// l
	// o
}

// 分割函数bufio.ScanRunes: 按字符分割
func TestScannerWithScanRunes(t *testing.T) {
	reader := strings.NewReader("天气ok")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	// output:
	// 天
	// 气
	// o
	// k
}
