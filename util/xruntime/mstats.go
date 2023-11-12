package xruntime

import (
	"log"
	"runtime"
)

func PrintNumGoroutine() {
	log.Printf("goroutine num: %d\n", NumGoroutine())
}

func NumGoroutine() int {
	return runtime.NumGoroutine()
}

func PrintAlloc() {
	log.Printf("mem alloc: %d KB\n", MemAlloc()/1024)
}

func MemAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
