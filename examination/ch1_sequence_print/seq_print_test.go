package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// another solution
func TestSolution(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	count := 3

	animalList := []string{"cat", "dog", "fish"}
	chanList := []chan struct{}{make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)}
	size := len(chanList)

	for i := range animalList {
		go printAnimal(wg, count, animalList[i], chanList[i], chanList[(i+1)%size])
	}
	chanList[0] <- struct{}{}

	wg.Wait()
}

func printAnimal(wg *sync.WaitGroup, count int, animal string, c1 chan struct{}, c2 chan struct{}) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		<-c1
		fmt.Println(i, animal)
		c2 <- struct{}{}
	}
}

func TestSeqPrint1234(t *testing.T) {
	const num = 4
	chanList := []chan struct{}{make(chan struct{}, 1), make(chan struct{}, 1),
		make(chan struct{}, 1), make(chan struct{}, 1)}
	printNum := func(id int, recv chan struct{}, send chan struct{}) {
		for {
			token := <-recv
			log.Println(">>", id)
			time.Sleep(time.Second)
			send <- token
		}
	}
	for i := 0; i < num; i++ {
		id := i + 1
		go printNum(id, chanList[i], chanList[(i+1)%num])
	}
	chanList[0] <- struct{}{}
	select {}
}
