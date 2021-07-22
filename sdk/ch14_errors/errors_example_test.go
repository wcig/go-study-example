package ch14_errors

import (
	"fmt"
	"testing"
)

// 注意包装error后的struct,创建时必须error类型接收才能正确进行nil判断
type AErr struct {
	C int
	E error
}

func NewAErr(code int, err error) *AErr {
	if err == nil {
		return nil
	}
	return &AErr{C: code, E: err}
}

func (e AErr) Error() string {
	return e.E.Error()
}

type BErr struct {
	C int
	E error
}

func NewBErr(code int, err error) error {
	if err == nil {
		return nil
	}
	return &AErr{C: code, E: err}
}

func (e BErr) Error() string {
	return e.E.Error()
}

func TestWrapStructErr(t *testing.T) {
	{
		var e1 error
		fmt.Println(e1 == nil) // true
		e1 = NewAErr(0, e1)
		fmt.Println(e1 == nil) // false
		e2 := NewAErr(0, e1)
		fmt.Println(e2 == nil) // false
	}
	{
		var e1 error
		fmt.Println(e1 == nil) // true
		e1 = NewBErr(0, e1)
		fmt.Println(e1 == nil) // true
		e2 := NewBErr(0, e1)
		fmt.Println(e2 == nil) // true
	}
}
