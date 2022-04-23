package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int64 = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// race detector
// $ go run -race main.go
// ==================
// WARNING: DATA RACE
// Read at 0x00c0000bc018 by goroutine 8:
//  main.main.func1()
//      /Users/yangbo/Documents/MyGithub/go-study-example/basic/ch20_race/main.go:16 +0x84
//
// Previous write at 0x00c0000bc018 by goroutine 7:
//  main.main.func1()
//      /Users/yangbo/Documents/MyGithub/go-study-example/basic/ch20_race/main.go:16 +0x98
//
// Goroutine 8 (running) created at:
//  main.main()
//      /Users/yangbo/Documents/MyGithub/go-study-example/basic/ch20_race/main.go:13 +0xc4
//
// Goroutine 7 (finished) created at:
//  main.main()
//      /Users/yangbo/Documents/MyGithub/go-study-example/basic/ch20_race/main.go:13 +0xc4
// ==================
// 35936
// Found 1 data race(s)
// exit status 66
