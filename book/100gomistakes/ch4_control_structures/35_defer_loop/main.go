package main

import (
	"log"
	"time"
)

func main() {
	example1()

	log.Println("----------------------")

	example2()

	// Output:
	// 2023/11/24 19:55:01 >> range: 0
	// 2023/11/24 19:55:02 >> range: 1
	// 2023/11/24 19:55:03 >> range: 2
	// 2023/11/24 19:55:04 >> defer: 2
	// 2023/11/24 19:55:04 >> defer: 1
	// 2023/11/24 19:55:04 >> defer: 0
	// 2023/11/24 19:55:04 ----------------------
	// 2023/11/24 19:55:04 >> range: 0
	// 2023/11/24 19:55:05 >> defer: 0
	// 2023/11/24 19:55:05 >> range: 1
	// 2023/11/24 19:55:06 >> defer: 1
	// 2023/11/24 19:55:06 >> range: 2
	// 2023/11/24 19:55:07 >> defer: 2
}

// bad
func example1() {
	for i := 0; i < 3; i++ {
		v := i
		defer log.Println(">> defer:", v)
		log.Println(">> range:", v)
		time.Sleep(time.Second)
	}
}

// good
func example2() {
	for i := 0; i < 3; i++ {
		v := i
		handle(v)
	}
}

func handle(v int) {
	defer log.Println(">> defer:", v)
	log.Println(">> range:", v)
	time.Sleep(time.Second)
}
