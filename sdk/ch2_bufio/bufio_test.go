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
func TestReader_ReadSlice(t *testing.T) {
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
func TestReader_ReadBytes(t *testing.T) {
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
func TestReader_ReadString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))

	line1, _ := reader.ReadString('|')
	fmt.Println("line1:", line1)

	line2, _ := reader.ReadString('|')
	fmt.Println("line2:", line2)
	fmt.Println("line1:", line1)
	// output:
	// line1: hello world.|
	// line2: It is ok
	// line1: hello world.|
}

// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
// 是一个低级别的行读取原语。 大多数调用者应改用ReadBytes('\n')或ReadString('\n')或使用扫描仪。
func TestReader_ReadLine(t *testing.T) {
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
func TestReader_Read(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("hello world.|It is ok"))
	buf := make([]byte, 100)
	n, err := reader.Read(buf)
	fmt.Println(n, err)           // 21 <nil>
	fmt.Println(string(buf[0:n])) // hello world.|It is ok
}

// func (b *Reader) ReadByte() (byte, error)
// 读取并返回一个字节,如果没有字节将返回错误 (建议使用bufio.Scanner更方便)
func TestReader_ReadByte(t *testing.T) {
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
func TestReader_ReadRune(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("天气ok"))
	for {
		r, n, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Println(string(r), n)
	}
	// output:
	// 天 3
	// 气 3
	// o 1
	// k 1
}

// func (b *Reader) UnreadByte() error
// 取消读取最后一个字节。 只有最近读取的字节可以不被读取。
func TestReader_UnreadByte(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	b, err := r.ReadByte()
	fmt.Println(string(b), err) // h <nil>

	err = r.UnreadByte()
	assert.Nil(t, err)

	b, err = r.ReadByte()
	fmt.Println(string(b), err) // h <nil>
}

// func (b *Reader) UnreadRune() error
// 取消读取最后一个字符。 如果在Reader上调用的最新方法不是ReadRune，则UnreadRune返回错误。
func TestReader_UnreadRune(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("天气ok"))
	rr, n, err := r.ReadRune()
	fmt.Println(string(rr), n, err) // 天 3 <nil>

	err = r.UnreadRune()
	assert.Nil(t, err)

	rr, n, err = r.ReadRune()
	fmt.Println(string(rr), n, err) // 天 3 <nil>
}

// func (b *Reader) Buffered() int
// 返回可从当前缓冲区读取的字节数。
func TestReader_Buffered(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	fmt.Println(r.Buffered())

	b, err := r.ReadByte()
	fmt.Println(string(b), err) // h <nil>

	fmt.Println(r.Buffered())
	// output:
	// 0
	// h <nil>
	// 4
}

// func (b *Reader) Discard(n int) (discarded int, err error)
// 跳过n个字节,返回跳过的字节数
func TestReader_Discard(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	discard, err := r.Discard(2)
	fmt.Println(discard, err) // 2 <nil>

	b, err := r.ReadByte()
	fmt.Println(string(b), err) // l <nil>
}

// func (b *Reader) Peek(n int) ([]byte, error)
// 返回下一个n个字节，而不会使阅读器前进。 字节在下一个读取调用时不再有效。 如果Peek返回的字节数少于n个字节，则它还会返回一个错误，说明读取短的原因。 如果n大于b的缓冲区大小，则错误为ErrBufferFull。
func TestReader_Peek(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	b, err := r.Peek(2)
	fmt.Println(string(b), err) // he <nil>

	b2, err2 := r.ReadByte()
	fmt.Println(string(b2), err2) // h <nil>
}

// func (b *Reader) Reset(r io.Reader)
// 丢弃所有缓冲的数据，重置所有状态，并将缓冲的读取器切换为从r读取。
func TestReader_Reset(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	b, err := r.ReadByte()
	fmt.Println(string(b), err) // h <nil>

	r.Reset(strings.NewReader("world"))

	b, err = r.ReadByte()
	fmt.Println(string(b), err) // w <nil>
}

// func (b *Reader) Size() int
// 返回底层缓冲区的大小（以字节为单位）。
func TestReader_Size(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello"))
	fmt.Println(r.Size()) // 4096
}

// func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
// 写入数据到w
func TestReader_WriteTo(t *testing.T) {
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
func TestWriter_WriteString(t *testing.T) {
	rb := &bytes.Buffer{}
	writer := bufio.NewWriter(rb)
	n, err := writer.WriteString("hello world.")
	fmt.Println(n, err)      // 12 <nil>
	fmt.Println(rb.String()) //
}

// func (b *Writer) Flush() error
// 缓冲区数据写入底层io.Writer中
func TestWriter_Flush(t *testing.T) {
	rb := &bytes.Buffer{}
	w := bufio.NewWriter(rb)
	_, err := w.WriteString("hello world.")
	assert.Nil(t, err)

	err = w.Flush()
	assert.Nil(t, err)
	fmt.Println(rb.String()) // hello world.
}

// func (b *Writer) Write(p []byte) (nn int, err error)
// 写操作将p的内容写入缓冲区。 它返回写入的字节数。如果nn<len(p)，它还会返回一个错误，解释为什么写入被截断。
func TestWriter_Write(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	n, err := w.Write([]byte("hello"))
	fmt.Println(n, err) // 5 <nil>

	fmt.Println(buf.String()) //
	_ = w.Flush()
	fmt.Println(buf.String()) // hello
}

// func (b *Writer) WriteByte(c byte) error
// 写入一个字节
func TestWriter_WriteByte(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err := w.WriteByte('h')
	assert.Nil(t, err)

	fmt.Println(buf.String()) //
	_ = w.Flush()
	fmt.Println(buf.String()) // h
}

// func (b *Writer) WriteRune(r rune) (size int, err error)
// 写入一个unicode字符，返回写入的字节数和错误
func TestWriter_WriteRune(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	size, err := w.WriteRune('好')
	fmt.Println(size, err) // 3 <nil>

	fmt.Println(buf.String()) //
	_ = w.Flush()
	fmt.Println(buf.String()) // h
}

// func (b *Writer) Available() int
// 返回缓冲区未被使用的字节数
func TestWriter_Available(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	fmt.Println(w.Available()) // 4096

	n, err := w.Write([]byte("hello"))
	fmt.Println(n, err) // 5 <nil>

	fmt.Println(w.Available()) // 4091
}

// func (b *Writer) Buffered() int
// 返回已经写入当前缓冲区的字节数
func TestWriter_Buffered(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	fmt.Println(w.Buffered()) // 0

	n, err := w.Write([]byte("hello"))
	fmt.Println(n, err) // 5 <nil>

	fmt.Println(w.Buffered()) // 5
}

// func (b *Writer) Reset(w io.Writer)
// 丢弃所有未刷新的缓冲数据，清除所有错误，然后复位b将其输出写入w。
func TestWriter_Reset(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	n, err := w.WriteString("hello")
	fmt.Println(n, err) // 5 <nil>
	_ = w.Flush()
	fmt.Println(buf.String()) // hello

	var buf2 bytes.Buffer
	w.Reset(&buf2)
	n2, err2 := w.WriteString("你好")
	fmt.Println(n2, err2) // 6 <nil>
	_ = w.Flush()
	fmt.Println(buf2.String()) // 你好
}

// func (b *Writer) Size() int
// 返回底层缓冲区字节大小
func TestWriter_Size(t *testing.T) {
	w1 := bufio.NewWriter(os.Stdout)
	fmt.Println(w1.Size()) // 4096

	w2 := bufio.NewWriterSize(os.Stdout, 16)
	fmt.Println(w2.Size()) // 16
}

// func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
// 实现io.ReaderFrom。 如果底层Writer支持ReadFrom方法，并且b尚无缓冲数据，则这将调用基础ReadFrom而不进行缓冲。
func TestWriter_ReadFrom(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	n, err := w.WriteString("hello")
	fmt.Println(n, err) // 5 <nil>
	_ = w.Flush()
	fmt.Println(buf.String()) // hello

	nn, err := w.ReadFrom(strings.NewReader("你好"))
	fmt.Println(nn, err)      // 6 <nil>
	fmt.Println(buf.String()) // hello你好
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

// func NewScanner(r io.Reader) *Scanner
// 返回一个新的Scanner以从r读取。 拆分功能默认为ScanLines。(分割函数bufio.ScanLines: 按换行符'\n'分割)
// func (s *Scanner) Scan() bool: 扫描后面是否还有内容
// func (s *Scanner) Text() string: 返回扫描的文本
// func (s *Scanner) Err() error: 返回扫描中出现的第一个非EOF错误
func TestScannerWithScanLines(t *testing.T) {
	reader := strings.NewReader("hello world.\nit is ok\n666")
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	assert.Nil(t, err)
	// output:
	// hello world.
	// it is ok
	// 666
}

// func (s *Scanner) Split(split SplitFunc): 使用指定的分割函数
// 分割函数bufio.ScanWords: 按单词分割
func TestScannerWithScanWords(t *testing.T) {
	reader := strings.NewReader("hello world. 666")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	assert.Nil(t, err)
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
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	assert.Nil(t, err)
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
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	assert.Nil(t, err)
	// output:
	// 天
	// 气
	// o
	// k
}

// func (s *Scanner) Bytes() []byte
// 返回扫描的字节切片
func TestScanner_Bytes(t *testing.T) {
	reader := strings.NewReader("hello world.\nit is ok\n666")
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		b := scanner.Bytes()
		fmt.Println(len(b), string(b))
	}
	err := scanner.Err()
	assert.Nil(t, err)
	// output:
	// 12 hello world.
	// 8 it is ok
	// 3 666
}

// func (s *Scanner) Buffer(buf []byte, max int)
// 设置初始扫描的缓冲区和扫描时可分配最大缓冲区大小
func TestScanner_Buffer(t *testing.T) {
	text := strings.Repeat("x", 2*bufio.MaxScanTokenSize)
	s := bufio.NewScanner(strings.NewReader(text + "\n"))
	s.Buffer(make([]byte, 100), 3*bufio.MaxScanTokenSize)
	for s.Scan() {
		token := s.Text()
		if token != text {
			t.Errorf("scan got incorrect token of length %d", len(token))
		}
	}
	if s.Err() != nil {
		t.Fatal("after scan:", s.Err())
	}
}

// 测试读取或写入数据超过缓冲区大小
func TestExceedBuffedSize(t *testing.T) {
	src := "012345678901234567890123456789"
	r := bufio.NewReaderSize(strings.NewReader(src), 16)
	val, err := r.ReadString('\n')
	assert.Equal(t, val, src)
	assert.Equal(t, io.EOF, err)

	buf := &bytes.Buffer{}
	w := bufio.NewWriterSize(buf, 16)
	n, err := w.WriteString(src)
	fmt.Println(n, err) // 30 <nil>
	_ = w.Flush()
	assert.Equal(t, src, buf.String())
}
