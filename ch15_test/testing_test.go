package ch15_test

import (
	"math/rand"
	"testing"
)

/* 单元测试: package testing */

// 单元测试简单样例
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

// 表组测试
func TestTableDriven(t *testing.T) {
	dataList := []struct {
		Input  int
		Except int
	}{
		{
			-1,
			1,
		},
		{
			1,
			1,
		},
		{
			0,
			0,
		},
	}

	for _, data := range dataList {
		result := Abs(data.Input)
		if result != data.Except {
			t.Errorf("Abs(%d) = %d; want %d", data.Input, result, data.Except)
		}
	}
}

// 基准测试
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
