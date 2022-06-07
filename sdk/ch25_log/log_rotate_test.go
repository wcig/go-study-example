package ch25_log

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

type MyLogger struct {
	Filename string
	MaxSize  int // megabytes

	size int64
	file *os.File
	mu   sync.Mutex
}

var (
	megabyte = 1024 * 1024
)

func (l *MyLogger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		if err = l.openExistOrNew(); err != nil {
			return 0, err
		}
	}

	writeLen := int64(len(p))
	if l.size+writeLen > l.max() {
		if err = l.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)
	l.size += int64(n)
	return n, err
}

func (l *MyLogger) openExistOrNew() error {
	info, err := os.Stat(l.Filename)
	if err == nil {
		size := info.Size()
		file, err := os.OpenFile(l.Filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return l.openNew()
		}
		l.file = file
		l.size = size
		return nil
	}

	if os.IsNotExist(err) {
		return l.openNew()
	}

	return fmt.Errorf("get log file info err: %v", err)
}

func (l *MyLogger) openNew() error {
	file, err := os.OpenFile(l.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open new log file err: %v", err)
	}
	l.file = file
	l.size = 0
	return nil
}

// TODO
func (l *MyLogger) rotate() error {
	return nil
}

func (l *MyLogger) max() int64 {
	if l.MaxSize > 0 {
		return int64(l.MaxSize) * int64(megabyte)
	}
	return 100 * int64(megabyte)
}

var (
	str = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func TestLogRotate(t *testing.T) {
	start := time.Now()

	logger := log.New(ioutil.Discard, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l := &MyLogger{
		Filename: "info.log",
		MaxSize:  1,
	}
	logger.SetOutput(l)
	for i := 0; i < 2000; i++ {
		logger.Println(str)
	}

	fmt.Println(time.Since(start).Milliseconds(), "ms")
}
