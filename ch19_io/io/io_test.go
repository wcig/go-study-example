package io

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("aaa")
}

// ReaderFrom接口
func TestReaderFrom(t *testing.T) {
	file, err := os.Open("temp1") // 文本: hello world.
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	buf := bytes.NewBufferString("ok")
	num, err := buf.ReadFrom(file)
	fmt.Println(num, err)     // 12 <nil>
	fmt.Println(buf.String()) // okhello world.
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

// CopyBuffer
func TestCopyBuffer(t *testing.T) {
	rb := &bytes.Buffer{}
	wb := &bytes.Buffer{}

	rb.WriteString("hello, world.")
	io.CopyBuffer(wb, rb, make([]byte, 1))
	if wb.String() != "hello, world." {
		t.Errorf("CopyBuffer did not work properly")
	}
}

// MultiReader
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

// MultiWriter
func TestMultiWriter(t *testing.T) {
	file, err := os.Create("temp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("hello world.\n"))
}

// TeeReader
func TestTeeReader(t *testing.T) {
	reader := io.TeeReader(strings.NewReader("hello world.\n"), os.Stdout)
	reader.Read(make([]byte, 20))
}
