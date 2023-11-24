package main

import "fmt"

func main() {
	listing1()
	// Output:
	// map[0:true 1:false 2:true 10:true 12:true 20:true 22:true]
	// 或
	// map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true]
	// ...

	listing2()
	// Output:
	// map[0:true 1:false 2:true 10:true 12:true]
}

func listing1() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}

	fmt.Println(m)
}

func listing2() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)

	for k, v := range m {
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	res := make(map[int]bool, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
