package main

import "testing"

var global [2]int64

func BenchmarkAdd(b *testing.B) {
	a := [2]int64{}
	var local [2]int64
	for i := 0; i < b.N; i++ {
		local = add(a)
	}
	global = local
}

func BenchmarkAdd2(b *testing.B) {
	a := [2]int64{}
	var local [2]int64
	for i := 0; i < b.N; i++ {
		local = add2(a)
	}
	global = local
}

// BenchmarkAdd
// BenchmarkAdd-8    	     505	   2248629 ns/op
// BenchmarkAdd2
// BenchmarkAdd2-8   	     562	   2121188 ns/op
