package cond

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestRight(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	const num = 10

	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员 [%d] 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != num {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")

	// Output:
	// 2022/04/23 23:42:34 运动员 [9] 已准备就绪
	// 2022/04/23 23:42:34 运动员 [5] 已准备就绪
	// 2022/04/23 23:42:34 裁判员被唤醒一次
	// 2022/04/23 23:42:35 运动员 [1] 已准备就绪
	// 2022/04/23 23:42:35 裁判员被唤醒一次
	// 2022/04/23 23:42:35 运动员 [0] 已准备就绪
	// 2022/04/23 23:42:35 裁判员被唤醒一次
	// 2022/04/23 23:42:35 运动员 [4] 已准备就绪
	// 2022/04/23 23:42:35 裁判员被唤醒一次
	// 2022/04/23 23:42:40 运动员 [8] 已准备就绪
	// 2022/04/23 23:42:40 裁判员被唤醒一次
	// 2022/04/23 23:42:41 运动员 [2] 已准备就绪
	// 2022/04/23 23:42:41 裁判员被唤醒一次
	// 2022/04/23 23:42:42 运动员 [7] 已准备就绪
	// 2022/04/23 23:42:42 运动员 [3] 已准备就绪
	// 2022/04/23 23:42:42 裁判员被唤醒一次
	// 2022/04/23 23:42:43 运动员 [6] 已准备就绪
	// 2022/04/23 23:42:43 裁判员被唤醒一次
	// 2022/04/23 23:42:43 所有运动员都准备就绪。比赛开始，3，2，1, ......
}

func TestWrong1(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	const num = 10

	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员 [%d] 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	// c.L.Lock()
	for ready != num {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	// c.L.Unlock()

	// 所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")

	// Output:
	// fatal error: sync: unlock of unlocked mutex
}

func TestWrong2(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	const num = 10

	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员 [%d] 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	// for ready != num {
	c.Wait()
	log.Println("裁判员被唤醒一次")
	// }
	c.L.Unlock()

	// 所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")

	// Output:
	// 2022/04/23 23:45:24 运动员 [9] 已准备就绪
	// 2022/04/23 23:45:24 裁判员被唤醒一次
	// 2022/04/23 23:45:24 所有运动员都准备就绪。比赛开始，3，2，1, ......
}
