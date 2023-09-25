package ch14_concurrent

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

// Go 程序可创建最大 goroutine 依赖系统内存 (4GB内存即可创建一百万goroutine)

const (
	n int = 1e5 // Number of goroutines to create
)

var ch = make(chan byte)
var counter = 0

func f() {
	counter++
	<-ch // Block this goroutine
}

func TestMaxGoroutines(t *testing.T) {
	time.Sleep(10 * time.Second)
	start := time.Now()
	fmt.Println(">> start:", start)

	// Limit the number of spare OS threads to just 1
	runtime.GOMAXPROCS(1)

	// Make a copy of MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	t0 := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		go f()
	}
	runtime.Gosched()
	t1 := time.Now().UnixNano()
	runtime.GC()

	// Make a copy of MemStats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	if counter != n {
		fmt.Fprintf(os.Stderr, "failed to begin execution of all goroutines")
		os.Exit(1)
	}

	fmt.Printf("Number of goroutines: %d\n", n)
	fmt.Printf("Per goroutine:\n")
	fmt.Printf("  Memory: %.2f bytes\n", float64(m1.Sys-m0.Sys)/float64(n))
	fmt.Printf("  Time:   %f µs\n", float64(t1-t0)/float64(n)/1e3)

	fmt.Println(">> end:", time.Now(), time.Since(start))
	time.Sleep(time.Minute)

	// Output:
	// >> start: 2023-09-25 21:36:05.324375 +0800 CST m=+10.002434668
	// Number of goroutines: 100000
	// Per goroutine:
	//  Memory: 2716.54 bytes
	//  Time:   0.958670 µs
	// >> end: 2023-09-25 21:36:05.493801 +0800 CST m=+10.171863209 169.428708ms
}
