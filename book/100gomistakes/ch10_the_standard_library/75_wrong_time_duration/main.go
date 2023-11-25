package main

import (
	"fmt"
	"time"
)

func main() {
	listing1()

	// listing2()
}

// bad
func listing1() {
	ticker := time.NewTicker(1000)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}

// good
func listing2() {
	ticker := time.NewTicker(time.Microsecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
