package main

import (
	"fmt"
	"reflect"
)

type customer1 struct {
	id string
}

type customer2 struct {
	id         string
	operations []float64
}

func main() {
	var a any = 3
	var b any = 3
	fmt.Println(a == b) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// Doesn't compile
	// fmt.Println(s1 == s2)
	fmt.Println(reflect.DeepEqual(s1, s2)) // true

	cust11 := customer1{id: "x"}
	cust12 := customer1{id: "x"}
	fmt.Println(cust11 == cust12) // true

	cust21 := customer2{id: "x", operations: []float64{1.}}
	cust22 := customer2{id: "x", operations: []float64{1.}}
	// Doesn't compile
	// fmt.Println(cust21 == cust22)
	_ = cust21
	_ = cust22
	fmt.Println(cust21.equal(cust22))              // true
	fmt.Println(reflect.DeepEqual(cust21, cust22)) // true
}

func (a customer2) equal(b customer2) bool {
	if a.id != b.id {
		return false
	}
	if len(a.operations) != len(b.operations) {
		return false
	}
	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}
