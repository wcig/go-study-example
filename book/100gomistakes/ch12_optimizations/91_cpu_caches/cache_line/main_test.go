package main

import "testing"

var global int64

func BenchmarkSum2(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum2(s)
	}
	global = local
}

func BenchmarkSum8(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum8(s)
	}
	global = local
}

// BenchmarkSum2
// BenchmarkSum2-8   	    3639	    286184 ns/op
// BenchmarkSum8
// BenchmarkSum8-8   	    5934	    196166 ns/op
