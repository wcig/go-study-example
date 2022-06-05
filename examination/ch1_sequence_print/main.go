package main

import (
	"fmt"
	"sync"
)

// 题目: 3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上3次。
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	count := 3

	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	catChan <- struct{}{}
	go printCat(wg, count, catChan, dogChan)
	go printDog(wg, count, dogChan, fishChan)
	go printFish(wg, count, fishChan, catChan)

	wg.Wait()
}

func printCat(wg *sync.WaitGroup, count int, catChan chan struct{}, dogChan chan struct{}) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		<-catChan
		fmt.Println(i, "cat")
		dogChan <- struct{}{}
	}
}

func printDog(wg *sync.WaitGroup, count int, dogChan chan struct{}, fishChan chan struct{}) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		<-dogChan
		fmt.Println(i, "dog")
		fishChan <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, count int, fishChan chan struct{}, catChan chan struct{}) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		<-fishChan
		fmt.Println(i, "fish")
		catChan <- struct{}{}
	}
}
