package ch14_errors

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// errors：实现操作错误的函数

// 函数
// func As(err error, target interface{}) bool // err的错误链是否有与target匹配的，如果是则将target设置为err值并返回true
// func Is(err, target error) bool             // 报告err错误链中任何错误是否与target匹配
// func New(text string) error                 // 返回给定文本的错误（每一次调用都会返回一不同的错误值）
// func Unwrap(err error) error                // 如果err的类型包含返回错误的Unwrap方法，则Unwrap返回对err调用Unwrap方法的结果。否则Unwrap返回nil。

func TestAs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	err1 := errors.New("file not exists")
	err2 := fmt.Errorf("wrap %w", err1)
	fmt.Println(errors.As(err2, &err1))
	fmt.Printf("err1: %s, err2: %s\n", err1, err2)
	// output:
	// Failed at path: non-existing
	// true
	// err1: wrap file not exists, err2: wrap file not exists
}

func TestIs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

	err1 := errors.New("file not exists")
	err2 := fmt.Errorf("wrap %w", err1)
	fmt.Println(errors.Is(err2, err1))
	fmt.Printf("err1: %s, err2: %s\n", err1, err2)
	// output:
	// file does not exist
	// true
	// err1: file not exists, err2: wrap file not exists
}

func TestNew(t *testing.T) {
	err := errors.New("file not exists")
	fmt.Println(err)

	err2 := errors.New("file not exists")
	fmt.Println(err == err2)

	filename := "tmp.txt"
	err = fmt.Errorf("%s file not exists", filename)
	fmt.Println(err)
	// output:
	// file not exists
	// false
	// tmp.txt file not exists
}

func TestUnwrap(t *testing.T) {
	err1 := errors.New("file not exists")
	err2 := wrapped{"wrap", err1}
	resultErr := errors.Unwrap(err2)
	assert.Equal(t, err1, resultErr)
}

type wrapped struct {
	msg string
	err error
}

func (e wrapped) Error() string { return e.msg }

func (e wrapped) Unwrap() error { return e.err }
