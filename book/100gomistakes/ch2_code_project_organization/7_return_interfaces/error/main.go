package main

import (
	"fmt"
	"reflect"
)

// error interface nil 问题 (可以直接返回error类型)
func main() {
	{
		// example1
		var err error
		err = Handle1()
		fmt.Println(err == nil, reflect.TypeOf(err), reflect.ValueOf(err)) // true <nil> <invalid reflect.Value>
		err = Handle2()
		fmt.Println(err == nil, err.(*CustomError) == nil, reflect.TypeOf(err), reflect.ValueOf(err)) // false true *main.CustomError <nil>
	}

	{
		// example2
		var err error
		err = Handle2()
		fmt.Println(err == nil, err.(*CustomError) == nil) // false true
		err = Handle1()
		fmt.Println(err == nil) // true
	}

	{
		// example3
		err1 := Handle1()
		fmt.Println(err1 == nil) // true
		err2 := Handle2()
		fmt.Println(err2 == nil) // true
	}
}

type CustomError struct {
	Msg string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("custom error: %s", c.Msg)
}

func Handle1() error {
	return nil
}

func Handle2() *CustomError {
	return nil
}
