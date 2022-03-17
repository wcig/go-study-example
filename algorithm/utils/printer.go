package utils

import "fmt"

type Printer func(v interface{})

func IntPrinter(v interface{}) {
	fmt.Print(v.(int))
}
