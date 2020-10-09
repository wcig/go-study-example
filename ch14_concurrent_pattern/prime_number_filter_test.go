package ch14_concurrent_pattern

import (
	"fmt"
	"testing"
)

// 素数筛

func Test1(t *testing.T) {
	limit := 100
	num := 0

	for i := 2; ; i++ {
		if isPrimeNumber(i) {
			num++
			fmt.Println(num, ":", i)
			if num >= limit {
				break
			}
		}
	}
}

func isPrimeNumber(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	f := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			f = false
			break
		}
	}
	return f
}

func TestPrimeNumber(t *testing.T) {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
		fmt.Println("")
	}
}

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
