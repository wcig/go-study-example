package ch14_concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 无缓冲channel: 将会阻塞程序
func TestUnbufferedChannel(t *testing.T) {
	msg := make(chan string)
	go func() {
		fmt.Println("start...")
		time.Sleep(1 * time.Second)
		msg <- "ping"
	}()
	val, ok := <-msg
	fmt.Println(val, ok) // ping true (将等待发送goroutine发送完成)
}

// 缓冲channel
func TestBufferedChannel1(t *testing.T) {
	msg := make(chan string, 2)
	go func() {
		time.Sleep(1 * time.Second)
		msg <- "ping"
		msg <- "pong"
	}()
	fmt.Println(<-msg) // ping
	fmt.Println(<-msg) // pong
}

// 已关闭的channel，不能再给其发送消息，会导致panic
func TestChannel1(t *testing.T) {
	c := make(chan int)
	close(c)
	c <- 1
}

// output:
// panic: send on closed channel [recovered]
//	panic: send on closed channel

// 已关闭的channel，可以接收消息，只是接收的消息为零值
func TestChannel2(t *testing.T) {
	c := make(chan int)
	close(c)
	val, ok := <-c
	fmt.Println(val, ok) // 0 false
}

// 往nil channel发送消息或从它接受消息会一直阻塞，同时close关闭会panic
func TestChannel3(t *testing.T) {
	var c chan int
	// c <- 1
	// <-c
	close(c)
}

// channel的3种类型
func TestChannelType(t *testing.T) {
	// chan: 双向channel，可接收和发送消息
	c1 := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- 1
	}()
	fmt.Println(<-c1)

	// chan<-: 单向channel，只能发送消息
	c2 := make(chan<- int)
	go func() {
		c2 <- 2
	}()
	// fmt.Println(<-c2) // Invalid operation: <-c2 (receive from send-only type chan<- int)
	time.Sleep(100 * time.Millisecond)

	// <-chan: 单向channel，只能接受消息
	c3 := make(<-chan int)
	// c3 <- 3 // invalid operation: c3 <- 3 (send to receive-only type <-chan int)
	go func() {
		<-c3
	}()
	time.Sleep(100 * time.Millisecond)
}

//
func TestChannelError1(t *testing.T) {
	// 错误
	// c1 := make(chan int)
	// c1 <- 1
	// fmt.Println(<-c1) // fatal error: all goroutines are asleep - deadlock!

	// 正确
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()
	fmt.Println(<-c1)

	// 错误
	// c2 := make(chan int)
	// fmt.Println(<-c2) // fatal error: all goroutines are asleep - deadlock!
	// c2 <- 1

	// 正确
	c2 := make(chan int)
	go func() {
		fmt.Println(<-c2)
	}()
	c2 <- 1
}

// 发送、接收顺序
func TestChannelBlock(t *testing.T) {
	// 先发送再接收
	c1 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("c1 set val...")
		c1 <- 1
	}()
	fmt.Println(<-c1)

	// 先接收再发送
	c2 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("c2 get val...")
		fmt.Println(<-c2)
	}()
	c2 <- 2
	close(c2)
}

// 已关闭Channel依然可以接收数据
func TestChannelReceive(t *testing.T) {
	// 已关闭无缓冲Channel，关闭后接收数据为零值 + false
	c1 := make(chan int)
	close(c1)
	val1, ok1 := <-c1
	fmt.Println(val1, ok1)

	// 已关闭无缓冲Channel且Channel没有数据，关闭后接收数据为零值 + false
	c2 := make(chan int, 2)
	close(c2)
	val2, ok2 := <-c2
	fmt.Println(val2, ok2)

	// 已关闭无缓冲Channel且Channel没有数据，关闭后接收数据发送数据，全部数据接收完时为零值 + false
	c3 := make(chan int, 2)
	c3 <- 3
	close(c3)
	val3, ok3 := <-c3
	fmt.Println(val3, ok3)
	val3, ok3 = <-c3
	fmt.Println(val3, ok3)
}

// output:
// 0 false
// 0 false
// 3 true
// 0 false

// 缓冲通道: 一组worker完成一组task
func TestWorkerBufferedChannel(t *testing.T) {
	const (
		workerNum = 4
		taskNum   = 10
	)

	var wg sync.WaitGroup
	wg.Add(workerNum)
	tasks := make(chan int, taskNum)

	for i := 1; i <= workerNum; i++ {
		go worker(tasks, i, &wg)
	}

	for i := 1; i <= taskNum; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	fmt.Println("over...")
}

func worker(tasks chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-tasks
		if !ok {
			fmt.Printf("worker: %d, shutdown\n", id)
			break
		}

		fmt.Printf("worker: %d, start    task :%d\n", id, val)
		time.Sleep(time.Duration(rand.Int63n(500)) * time.Millisecond)
		fmt.Printf("worker: %d, complete task :%d\n", id, val)
	}
}

// range
func TestRange(t *testing.T) {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	for val := range c { // 如果channel没有关闭，在取完所有数据后会产生panic错误：fatal error: all goroutines are asleep - deadlock!
		fmt.Println(val)
	}
	fmt.Println("over...")
}

// select
func TestSelect(t *testing.T) {
	c1 := make(chan int, 2)
	go func() {
		for i := 1; i <= 5; i++ {
			c1 <- i
		}
		close(c1)
	}()

	c2 := make(chan int, 2)
	go func() {
		for c := 101; c <= 105; c++ {
			c2 <- c
		}
		close(c2)
	}()

	f1, f2 := false, false
	for {
		if f1 && f2 {
			break
		}

		select {
		case val, ok := <-c1:
			if !ok {
				fmt.Println("c1 shutdown")
				f1 = true
			} else {
				fmt.Println("c1:", val, ok)
			}
		case val, ok := <-c2:
			if !ok {
				f2 = true
				fmt.Println("c2 shutdown")
			} else {
				fmt.Println("c2:", val, ok)
			}
		default:
			fmt.Println("wait")
		}
	}
	fmt.Println("over...")
}
