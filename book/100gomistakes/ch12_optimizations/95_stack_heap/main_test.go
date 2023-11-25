package main

import "testing"

var (
	globalValue int
	globalPtr   *int
)

func BenchmarkSumValue(b *testing.B) {
	b.ReportAllocs()
	var local int
	for i := 0; i < b.N; i++ {
		local = sumValue(i, i)
	}
	globalValue = local
}

func BenchmarkSumPtr(b *testing.B) {
	b.ReportAllocs()
	var local *int
	for i := 0; i < b.N; i++ {
		local = sumPtr(i, i)
	}
	globalValue = *local
}

// BenchmarkSumValue
// BenchmarkSumValue-8   	1000000000	         0.9862 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSumPtr
// BenchmarkSumPtr-8     	135349081	         8.785 ns/op	       8 B/op	       1 allocs/op
