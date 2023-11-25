package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	i := 0
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		i++
		wg.Done()
	}()

	go func() {
		i++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(i)
}

// $ go test -race .
// ==================
// WARNING: DATA RACE
// Read at 0x00c00012e158 by goroutine 8:
//  go-app/book/100gomistakes/ch11_testing/83_race.TestRace.func2()
//      /Users/yangbo/Documents/workspace/myproject/go-study-example/book/100gomistakes/ch11_testing/83_race/race_test.go:20 +0x34
//
// Previous write at 0x00c00012e158 by goroutine 7:
//  go-app/book/100gomistakes/ch11_testing/83_race.TestRace.func1()
//      /Users/yangbo/Documents/workspace/myproject/go-study-example/book/100gomistakes/ch11_testing/83_race/race_test.go:15 +0x44
//
// Goroutine 8 (running) created at:
//  go-app/book/100gomistakes/ch11_testing/83_race.TestRace()
//      /Users/yangbo/Documents/workspace/myproject/go-study-example/book/100gomistakes/ch11_testing/83_race/race_test.go:19 +0x198
//  testing.tRunner()
//      /Users/yangbo/go/go1.20/src/testing/testing.go:1576 +0x180
//  testing.(*T).Run.func1()
//      /Users/yangbo/go/go1.20/src/testing/testing.go:1629 +0x40
//
// Goroutine 7 (finished) created at:
//  go-app/book/100gomistakes/ch11_testing/83_race.TestRace()
//      /Users/yangbo/Documents/workspace/myproject/go-study-example/book/100gomistakes/ch11_testing/83_race/race_test.go:14 +0x100
//  testing.tRunner()
//      /Users/yangbo/go/go1.20/src/testing/testing.go:1576 +0x180
//  testing.(*T).Run.func1()
//      /Users/yangbo/go/go1.20/src/testing/testing.go:1629 +0x40
// ==================
// 2
// --- FAIL: TestRace (0.00s)
//    testing.go:1446: race detected during execution of test
// FAIL
// FAIL    go-app/book/100gomistakes/ch11_testing/83_race  0.416s
// FAIL
