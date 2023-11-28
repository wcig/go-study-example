package main

import "testing"

const n = 1e6

func BenchmarkCpuProfileByEmptySlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkCpuProfileByGiveLength(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		for j := 0; j < n; j++ {
			s[j] = j
		}
	}
}

// $ go test -bench=. -cpuprofile=cpu.profile
// goos: darwin
// goarch: arm64
// pkg: go-app/study/pprof/test
// BenchmarkCpuProfileByEmptySlice-8            387           2876780 ns/op
// BenchmarkCpuProfileByGiveLength-8           2370            503772 ns/op
