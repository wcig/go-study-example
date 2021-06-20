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
