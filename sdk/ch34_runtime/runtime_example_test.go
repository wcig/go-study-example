package ch34_runtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGC(t *testing.T) {
	runtime.GC()
}

func TestGOMAXPROCS(t *testing.T) {
	cpuNum := runtime.NumCPU()
	fmt.Println("device cpu num:", cpuNum)

	oldNum := runtime.GOMAXPROCS(cpuNum)
	fmt.Println("previous setting max process num:", oldNum)
	// output:
	// device cpu num: 8
	// previous setting max process num: 8
}

func TestGOROOT(t *testing.T) {
	root := runtime.GOROOT()
	fmt.Println("go root:", root) // go root: /usr/local/go
}

func TestGoexit(t *testing.T) {
	go func() {
		defer fmt.Println("defer A")
		func() {
			defer fmt.Println("defer B")
			runtime.Goexit() // 退出子协程
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	time.Sleep(3 * time.Second)
	// output:
	// defer B
	// defer A
}

func TestGosched(t *testing.T) {
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("hello")
		}
	}()

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("ok")
	}
	// output:
	// hello
	// hello
	// ok
	// ok
}

func TestNumCPU(t *testing.T) {
	fmt.Println(runtime.NumCPU()) // 8
}

func TestNumCgoCall(t *testing.T) {
	fmt.Println(runtime.NumCgoCall()) // 0
}

func TestNumGoroutine(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // 2
}

func TestReadMemStats(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc) // 155944
}

func TestVersion(t *testing.T) {
	fmt.Println(runtime.Version()) // go1.16.4
}
