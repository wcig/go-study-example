package main

import "testing"

const n = 1_000_000

var global []Bar

func BenchmarkConvert_EmptySlice(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertEmptySlice(foos)
	}
	global = local
}

func BenchmarkConvert_GivenCapacity(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenCapacity(foos)
	}
	global = local
}

func BenchmarkConvert_GivenLength(b *testing.B) {
	var local []Bar
	foos := make([]Foo, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = convertGivenLength(foos)
	}
	global = local
}

// Output:
// BenchmarkConvert_EmptySlice
// BenchmarkConvert_EmptySlice-8      	     528	   2114952 ns/op
// BenchmarkConvert_GivenCapacity
// BenchmarkConvert_GivenCapacity-8   	    3288	    359636 ns/op
// BenchmarkConvert_GivenLength
// BenchmarkConvert_GivenLength-8     	    3676	    320680 ns/op
