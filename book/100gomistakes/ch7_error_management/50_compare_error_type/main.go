package main

import (
	"errors"
	"fmt"
)

func main() {
	{
		// func As(err error, target interface{}) bool: err的错误链是否是target类型，如果是则将target设置为err值并返回true
		err1 := &FileNotExistsError{Path: "/root/non-existing"}
		err2 := fmt.Errorf("wrap %w", err1)
		var err3 *FileNotExistsError
		fmt.Println(errors.As(err2, &err3))
		fmt.Printf("err1: %s, err2: %s, err3: %s\n", err1, err2, err3)
		// Output:
		// true
		// err1: file not exists: /root/non-existing, err2: wrap file not exists: /root/non-existing, err3: file not exists: /root/non-existing
	}

	{
		// func Is(err, target error) bool: 报告err错误链中任何错误是否与target匹配
		err1 := &FileNotExistsError{Path: "/root/non-existing"}
		err2 := fmt.Errorf("wrap %w", err1)
		fmt.Println(errors.Is(err2, err1))
		fmt.Printf("err1: %s, err2: %s\n", err1, err2)
		// Output:
		// true
		// err1: file not exists: /root/non-existing, err2: wrap file not exists: /root/non-existing
	}
}

type FileNotExistsError struct {
	Path string
}

func (e *FileNotExistsError) Error() string {
	return fmt.Sprintf("file not exists: %s", e.Path)
}
