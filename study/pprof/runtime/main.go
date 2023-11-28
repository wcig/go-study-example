package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	file, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatalf("create cpu profile file err: %v", err)
	}

	if err = pprof.StartCPUProfile(file); err != nil {
		log.Fatalf("start cpu profile err: %v", err)
	}

	n := 0
	for i := 0; i < 1e6; i++ {
		n++
	}

	pprof.StopCPUProfile()
}
