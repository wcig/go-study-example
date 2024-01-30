package main

import (
	"os"
	"runtime/trace"
)

// 1.generate trace file: go run main.go 2> trace.out
// 2.view trace file: go tool trace trace.out
func main() {
	_ = trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "ok"
	}()
	<-ch
}
