package main

import (
	"fmt"
	"go-app/study/import_cycle/plan5/A"
	_ "go-app/study/import_cycle/plan5/A"
	"go-app/study/import_cycle/plan5/B"
	_ "go-app/study/import_cycle/plan5/B"
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
