package ch24_io

import (
	"io"
	"testing"
)

// io：提供IO原语的基础接口

// 变量
func TestVar(t *testing.T) {
	_ = io.Discard // 任何写入调用将不执行任何操作
}

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
// type Closer interface {
//    Close() error
// }

// 5.type LimitedReader: 限制从R读取数据最多返回的数据量为N个字节，每次调用Read方法后N都会变化
// type LimitedReader struct {
//    R Reader // underlying reader
//    N int64  // max bytes remaining
// }
// 方法：
// func (l *LimitedReader) Read(p []byte) (n int, err error)

// 6.type PipeReader: 管道读取部分
// type PipeReader struct {
//    // contains filtered or unexported fields
// }
// 方法：
// func (r *PipeReader) Close() error                        // 关闭Reader，后续写入管道的写入部分将返回ErrClosedPipe错误
// func (r *PipeReader) CloseWithError(err error) error      // 关闭Reader，后续写入管道的写入部分将返回ErrClosedPipe错误，不会覆盖先前的错误并且会返回nil
// func (r *PipeReader) Read(data []byte) (n int, err error) // 实现标准的Read接口，阻塞知道writer到达或写入端关闭

// 7.type PipeWriter: 管道写入部分
// type PipeWriter struct {
//    // contains filtered or unexported fields
// }
// 方法：
// func (w *PipeWriter) Close() error                         // 关闭writer，后续的从管道读取部分读取将不返回字节，同时返回错误
// func (w *PipeWriter) CloseWithError(err error) error       // 功能跟Close方法类似，但如果先前存在错误则不会覆盖，并总返回nil
// func (w *PipeWriter) Write(data []byte) (n int, err error) // 实现标准的Write接口，写入数据到管道，阻塞直到一个或多个reader消费完所有的数据或读取端关闭

// 8.type ReadCloser: 包含Read和Close方法的接口
// type ReadCloser interface {
//    Reader
//    Closer
// }
// 函数：
// func NopCloser(r Reader) ReadCloser: 将Reader r包装成ReadCloser，其Close方法没有逻辑

// 9.type ReadSeekCloser: 包含Read、Seek、Close3个方法的接口
// type ReadSeekCloser interface {
//    Reader
//    Seeker
//    Closer
// }

// 10.type ReadSeeker: 包含Read、Seek方法的接口
// type ReadSeeker interface {
//    Reader
//    Seeker
// }

// 11.type ReadWriteCloser: 包含Read、Write、Close3个方法的接口
// type ReadWriteCloser interface {
//    Reader
//    Writer
//    Closer
// }

// 12.type ReadWriteSeeker: 包含Read、Write、Seek3个方法的接口
// type ReadWriteSeeker interface {
//    Reader
//    Writer
//    Seeker
// }

// 13.type ReadWriter: 包含Read、Write方法接口
// type ReadWriter interface {
//    Reader
//    Writer
// }

// 14.type Reader: 包含一个Read方法的接口
// type Reader interface {
//    Read(p []byte) (n int, err error)
// }
// 函数：
// func LimitReader(r Reader, n int64) Reader // 包装r为一LimitedReader
// func MultiReader(readers ...Reader) Reader // 将多个reader串联成一个reader，按顺序读取
// func TeeReader(r Reader, w Writer) Reader  // 包装r为一个reader，同时读取的内容会写入w

// 15.type ReaderAt: 包含一ReadAt方法接口（与Reader接口区别在于可以以偏移量读取）
// type ReaderAt interface {
//    ReadAt(p []byte, off int64) (n int, err error)
// }

// 16.type ReaderFrom：包含一ReadFrom方法接口
// type ReaderFrom interface {
//    ReadFrom(r Reader) (n int64, err error)
// }

// 17.type RuneReader：包含一ReadRun方法接口
// type RuneReader interface {
//    ReadRune() (r rune, size int, err error)
// }

// 18.type RuneScanner: 在RuneReader接口基础上增加一UnreadRune方法接口（UnreadRune的后一次和前一次ReadRune方法返回数据相同，不能连续两次调用UnreadRune而中间没有ReadRune调用）
// type RuneScanner interface {
//    RuneReader
//    UnreadRune() error
// }

// 19.type SectionReader：在底层 ReaderAt 的一部分上实现 Read、Seek 和 ReadAt。
// type SectionReader struct {
//    // contains filtered or unexported fields
// }
// 函数：
// func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader // 以r、偏移量off和读取n个字节或EOF结束构造一SectionReader
// 方法：
// func (s *SectionReader) Read(p []byte) (n int, err error)
// func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
// func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
// func (s *SectionReader) Size() int64 // 返回section字节大小

// 20.type Seeker：包含一Seek方法的接口
// type Seeker interface {
//    Seek(offset int64, whence int) (int64, error)
// }

// 21.type StringWriter：包含一WriteString方法的接口
// type StringWriter interface {
//    WriteString(s string) (n int, err error)
// }

// 22.type WriteCloser：包含Write和Close方法的接口
// type WriteCloser interface {
//    Writer
//    Closer
// }

// 23.type WriteSeeker：包含Write、Seek方法接口
// type WriteSeeker interface {
//    Writer
//    Seeker
// }

// 24.type Writer：包含一Write方法接口
// type Writer interface {
//    Write(p []byte) (n int, err error)
// }
// 函数：
// func MultiWriter(writers ...Writer) Writer：组合多个writer为一个，其每次写入将复制到所有writer

// 25.type WriterAt：包含一WriteAt方法的接口
// type WriterAt interface {
//    WriteAt(p []byte, off int64) (n int, err error)
// }

// 26.type WriterTo：包含一WriteTo方法接口
// type WriterTo interface {
//    WriteTo(w Writer) (n int64, err error)
// }
