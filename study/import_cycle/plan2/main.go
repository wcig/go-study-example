package main

import (
	"fmt"
	"go-app/study/import_cycle/plan2/A"
	"go-app/study/import_cycle/plan2/B"
)

func main() {
	a := &A.PackageA{}
	b := &B.PackageB{}
	a.B = b
	b.A = a

	result := a.Foo("ok")
	fmt.Println(">>", result)

	result = b.Bar(result)
	fmt.Println("<<", result)
	// Output:
	// >> |ok|
	// << ok
}
