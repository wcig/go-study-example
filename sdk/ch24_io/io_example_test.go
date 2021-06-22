package ch24_io

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Copy
func TestCopy(t *testing.T) {
	rb := &bytes.Buffer{}
	wb := &bytes.Buffer{}

	rb.WriteString("hello, world.")
	written, err := io.Copy(wb, rb)
	if err != nil || written != 13 {
		t.Fatal(err)
	}
	assert.True(t, wb.String() == "hello, world.")
}

// CopyBuffer
func TestCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
	// output:
	// first reader
	// second reader
}

// CopyN
func TestCopyN(t *testing.T) {
	rb := &bytes.Buffer{}
	wb := &bytes.Buffer{}

	rb.WriteString("hello, world.")
	written, err := io.CopyN(wb, rb, 5)
	if err != nil || written != 5 {
		t.Fatal(err)
	}
	assert.True(t, wb.String() == "hello")
}

// Pipe
func TestPipe(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		n, err := w.Write([]byte("hello world.\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("成功写入字节数:", n)
		_ = w.Close()
	}()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("成功读取字节数:", n)
	// output:
	// hello world.
	// 成功写入字节数: 13
	// 成功读取字节数: 13
}

// ReadAll
func TestReadAll(t *testing.T) {
	src := "hello world."
	b, err := io.ReadAll(strings.NewReader(src))
	assert.Nil(t, err)
	assert.Equal(t, src, string(b))
}

// ReadAtLeast
func TestReadAtLeast(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	n, err := io.ReadAtLeast(r, buf, 4)
	assert.Nil(t, err)
	fmt.Println(n, string(buf))

	shortBuf := make([]byte, 3)
	_, err = io.ReadAtLeast(r, shortBuf, 4)
	fmt.Println(err == io.ErrShortBuffer, err)

	longBuf := make([]byte, 64)
	_, err = io.ReadAtLeast(r, longBuf, 64)
	fmt.Println(err == io.ErrUnexpectedEOF, err)
	// output:
	// 14 some io.Reader
	// true short buffer
	// true unexpected EOF
}

// ReadFull
func TestReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	n, err := io.ReadFull(r, buf)
	assert.Nil(t, err)
	fmt.Println(n, string(buf))

	longBuf := make([]byte, 64)
	_, err = io.ReadFull(r, longBuf)
	fmt.Println(err == io.ErrUnexpectedEOF, err)
	// output:
	// 4 some
	// true unexpected EOF
}

// WriteString
func TestWriteString(t *testing.T) {
	n, err := io.WriteString(os.Stdout, "hello world.\n")
	assert.Nil(t, err)
	fmt.Println(n)
	// output:
	// hello world.
	// 13
}

// ByteReader接口
func TestByteReader(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hello world."))
	b, err := buf.ReadByte()
	assert.Nil(t, err)
	fmt.Printf("%q\n", b) // 'h'
}

// ByteScanner接口
func TestByteScanner(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hello world."))
	b, err := buf.ReadByte()
	assert.Nil(t, err)
	fmt.Printf("%q\n", b)

	err = buf.UnreadByte()
	assert.Nil(t, err)

	b, err = buf.ReadByte()
	assert.Nil(t, err)
	fmt.Printf("%q\n", b)
	// output:
	// 'h'
	// 'h'
}

// ByteWriter接口
func TestByteWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	err := buf.WriteByte('h')
	assert.Nil(t, err)
	fmt.Println(buf.String()) // h
}

// Closer接口
func TestCloser(t *testing.T) {
	file, err := os.Open("io_example_test.go")
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)
}

// LimitedReader接口
func TestLimitedReader(t *testing.T) {
	r := strings.NewReader("hello world.")
	lr := io.LimitReader(r, 5)

	buf1 := make([]byte, 5)
	n, err := lr.Read(buf1)
	fmt.Println(n, err, string(buf1))

	buf2 := make([]byte, 5)
	n, err = lr.Read(buf2)
	fmt.Println(n, err)
	// output:
	// 5 <nil> hello
	// 0 EOF
}

// PipeReader接口
func TestPipeReader(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		n, err := w.Write([]byte("hello world.\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("成功写入字节数:", n)
		_ = w.Close()
	}()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("成功读取字节数:", n)
	// output:
	// hello world.
	// 成功写入字节数: 13
	// 成功读取字节数: 13
}

// PipeWriter接口
func TestPipeWriter(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		n, err := w.Write([]byte("hello world.\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("成功写入字节数:", n)
		_ = w.Close()
	}()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("成功读取字节数:", n)
	// output:
	// hello world.
	// 成功写入字节数: 13
	// 成功读取字节数: 13
}

// ReadCloser接口
func TestReadCloser(t *testing.T) {
	r := strings.NewReader("hello world.")
	rc := io.NopCloser(r)
	err := rc.Close()
	assert.Nil(t, err)
}

// LimitReader方法
func TestLimitReader(t *testing.T) {
	TestLimitedReader(t)
}

// MultiReader方法
func TestMultiReader(t *testing.T) {
	readers := []io.Reader{
		strings.NewReader("from strings reader..."),
		bytes.NewBufferString("from bytes buffer..."),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 128)
	buf := make([]byte, 10)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}

		data = append(data, buf[:n]...)
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%s\n", data) // from strings reader...from bytes buffer...
}

// TeeReader方法
func TestTeeReader(t *testing.T) {
	var r io.Reader = strings.NewReader("hello world.\n")
	r = io.TeeReader(r, os.Stdout)
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// output:
	// hello world.
	// hello world.
	//
}

// ReaderAt接口
func TestReaderAt(t *testing.T) {
	reader := strings.NewReader("ok-hello world.")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s, %d\n", p, n) // -hello, 6
}

// ReaderFrom接口
func TestReaderFrom(t *testing.T) {
	_ = ioutil.WriteFile("temp1", []byte("hello world."), os.ModePerm)
	defer os.Remove("temp1")

	file, err := os.Open("temp1")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	num, err := buf.ReadFrom(file)
	fmt.Println(num, err)
	fmt.Println(buf.String())
	// output:
	// 12 <nil>
	// hello world.
}

// RuneReader接口
func TestRuneReader(t *testing.T) {
	src := "tom你好"
	buf := bytes.NewBuffer([]byte(src))

	for {
		r, size, err := buf.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(r), size)
	}
	// output:
	// t 1
	// o 1
	// m 1
	// 你 3
	// 好 3
}

// RuneScanner接口
func TestRuneScanner(t *testing.T) {
	buf := bytes.NewBufferString("你好")
	r, size, _ := buf.ReadRune()
	fmt.Println(string(r), size)

	err := buf.UnreadRune()
	assert.Nil(t, err)

	r, size, _ = buf.ReadRune()
	fmt.Println(string(r), size)
	// output:
	// 你 3
	// 你 3
}

// SectionReader接口
func TestSectionReader(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	// output:
	// some io.Reader stream
}

// io.Discard
func TestDiscard(t *testing.T) {
	w := io.Discard
	n, err := w.Write([]byte("hello world."))
	fmt.Println(n, err) // 12 <nil>
}

// MultiWriter
func TestMultiWriter(t *testing.T) {
	file, err := os.Create("tmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	n, err := writer.Write([]byte("hello world.\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("成功写入字节数:", n)
	// output:
	// hello world.
	// 成功写入字节数: 13
}

// WriterAt接口
func TestWriterAt(t *testing.T) {
	file, err := os.Create("temp3")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	file.WriteString("ok-hello world.")
	n, err := file.WriteAt([]byte("666"), 8)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n) // 3
	// temp3: ok-hello666rld.
}

// WriterTo接口
func TestWriterTo(t *testing.T) {
	file, err := os.Create("temp2")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	buf := bytes.NewBufferString("hello ok")
	num, err := buf.WriteTo(file)
	fmt.Println(num, err)     // 8 <nil>
	fmt.Println(buf.String()) // hello ok
}
