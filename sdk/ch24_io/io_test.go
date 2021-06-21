package ch24_io

import (
	"io"
	"testing"
)

// io：提供IO原语的基础接口

// 常量
func TestConst(t *testing.T) {
	_ = io.SeekStart   // seek相对于文件起始位置
	_ = io.SeekCurrent // seek相对于文件当前偏移量
	_ = io.SeekEnd     // seek相对于文件结尾
}

// 错误
func TestErr(t *testing.T) {
	_ = io.EOF
	_ = io.ErrClosedPipe
	_ = io.ErrNoProgress
	_ = io.ErrShortBuffer
	_ = io.ErrShortWrite
	_ = io.ErrUnexpectedEOF
}

// 函数
// func Copy(dst Writer, src Reader) (written int64, err error)                   // 从src拷贝数据至dst,直至遇到EOF或错误
// func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) // 与Copy类似,区别在使用指定的buf而不是分配的临时缓冲区
// func CopyN(dst Writer, src Reader, n int64) (written int64, err error) // 从src拷贝n个字节到dst
// func Pipe() (*PipeReader, *PipeWriter) // 创建一同步的内存管道,用于连接io.Reader和io.Writer
// func ReadAll(r Reader) ([]byte, error) // 读取r所有数据直至遇到EOF或错误,成功返回err=nil
// func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) // 从r读取至少min个字节到buf,超过min则读取最多len(buf)个字节
// func ReadFull(r Reader, buf []byte) (n int, err error) // 从r读取准备的len(buf)个字节
// func WriteString(w Writer, s string) (n int, err error) // 往w写入字符串s

// 接口
// 1.type ByteReader: 包装唯一ReadByte方法接口
// type ByteReader interface {
//    ReadByte() (byte, error)
// }

// 2.type ByteScanner: ByteReader接口基础上增加一UnreadByte方法 (UnreadByte会导致下一次ReadByte和上一次返回同一个byte,两次UnreadByte之间没有ReadByte调用将返回错误)
// type ByteScanner interface {
//    ByteReader
//    UnreadByte() error
// }

// 3.type ByteWriter: 包装唯一的WriteByte方法接口
// type ByteWriter interface {
//    WriteByte(c byte) error
// }

// 4.type Closer: 包装唯一Close方法接口
