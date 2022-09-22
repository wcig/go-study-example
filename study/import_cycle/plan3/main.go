package main

import (
	"fmt"
	"go-app/study/import_cycle/plan3/A"
	"go-app/study/import_cycle/plan3/B"
	"go-app/study/import_cycle/plan3/C"
)

func main() {
	a := &A.PackageA{}
	b := &B.PackageB{}
	c := &C.PackageC{
		A: a,
		B: b,
	}

	result := c.FooAdd("ok")
	fmt.Println(">>", result)

	result = c.BarTrim(result)
	fmt.Println("<<", result)

	// Output:
	// >> |ok|
	// << ok
}
