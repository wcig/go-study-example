package main

import (
	"fmt"
	"go-app/study/import_cycle/plan4/A"
	"go-app/study/import_cycle/plan4/B"
	"go-app/study/import_cycle/plan4/C"
)

func init() {
	C.TrimFunc = A.Trim
	C.AddFunc = B.Add
}

func main() {
	result := A.Foo("ok")
	fmt.Println(">>", result)

	result = B.Bar(result)
	fmt.Println("<<", result)
	// Output:
	// >> |ok|
	// << ok
}
