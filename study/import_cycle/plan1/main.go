package main

import (
	"fmt"
	"go-app/study/import_cycle/plan1/A"
	"go-app/study/import_cycle/plan1/B"
)

func main() {
	result := A.Foo("ok")
	fmt.Println(">>", result)

	result = B.Bar(result)
	fmt.Println("<<", result)
	// Output:
	// >> |ok|
	// << ok
}
