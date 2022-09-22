package main

import (
	"fmt"
	"go-app/study/import_cycle/question/A"
	"go-app/study/import_cycle/question/B"
)

func main() {
	result := A.Foo("ok")
	fmt.Println(">>", result)

	result = B.Bar(result)
	fmt.Println("<<", result)
	// Output:
	// package go-app/study/import_cycle/question
	//	 imports go-app/study/import_cycle/question/A
	//	 imports go-app/study/import_cycle/question/B
	//	 imports go-app/study/import_cycle/question/A: import cycle not allowed
}
