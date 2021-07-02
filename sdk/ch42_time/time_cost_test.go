package ch42_time

import (
	"fmt"
	"testing"
	"time"
)

func Sum(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		sum += i
	}
	return sum
}

// 方式一
func TestTimeConst1(t *testing.T) {
	start := time.Now()
	Sum(100000000)
	fmt.Printf("TestTimeConst1 time cost:%v\n", time.Since(start))
}

// 方式二
func TestTimeConst2(t *testing.T) {
	defer printTimeCost(time.Now(), "TestTimeConst2")
	Sum(100000000)
}

func printTimeCost(start time.Time, funcName string) {
	fmt.Printf("%s time cost:%v\n", funcName, time.Since(start))
}

// 方式三
func TestTimeConst3(t *testing.T) {
	defer printTimeCostFunc("TestTimeConst3")()
	Sum(100000000)
}

func printTimeCostFunc(funcName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s time cost:%v\n", funcName, time.Since(start))
	}
}
