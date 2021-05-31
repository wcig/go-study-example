package bzip2

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// compress/bzip2: 实现了bzip2的解压缩,但没有实现压缩.

// func NewReader(r io.Reader) io.Reader
// 返回一个io.Reader，它从r解压缩bzip2数据。 如果r还没有实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。
func TestNewReader(t *testing.T) {
	reader := bzip2.NewReader(strings.NewReader("hello world."))
	assert.NotNil(t, reader)
}

// 测试bunzip2
func TestBunzip2(t *testing.T) {
	// 准备: bzip tmp.txt
	srcFile := "tmp.txt.bz2"
	dstFile := "tmp.txt"

	fr, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	br := bzip2.NewReader(fr)

	fw, err := os.Create(dstFile)
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	n, err := io.Copy(fw, br)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解压文件成功, 拷贝字节数:%d\n", n)
}

/******************************bzip command**********************************/
// demo:
// bzip -k tmp.txt
// bunzip2 -k tmp.txt.bz2

// $ bzip2 -h
// bzip2, a block-sorting file compressor.  Version 1.0.6, 6-Sept-2010.
//
// usage: bzip2 [flags and input files in any order]
//
// -h --help           print this message
// -d --decompress     force decompression
// -z --compress       force compression
// -k --keep           keep (don't delete) input files
// -f --force          overwrite existing output files
// -t --test           test compressed file integrity
// -c --stdout         output to standard out
// -q --quiet          suppress noncritical error messages
// -v --verbose        be verbose (a 2nd -v gives more)
// -L --license        display software version & license
// -V --version        display software version & license
// -s --small          use less memory (at most 2500k)
// -1 .. -9            set block size to 100k .. 900k
// --fast              alias for -1
// --best              alias for -9
//
// If invoked as `bzip2', default action is to compress.
//               as `bunzip2',  default action is to decompress.
// as `bzcat', default action is to decompress to stdout.
//
//    If no file names are given, bzip2 compresses or decompresses
//    from standard input to standard output.  You can combine
//    short flags, so `-v -4' means the same as -v4 or -4v, &c.
