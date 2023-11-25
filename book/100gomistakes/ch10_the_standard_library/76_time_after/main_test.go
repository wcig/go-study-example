package main

import (
	"context"
	"go-app/util/xruntime"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// 76.time.After是等到timer触发后才会被垃圾回收, select每次有channel到达都会触发time.After创建资源(200byte左右)
func TestTimeAfterMemoryLeak(t *testing.T) {
	e := make(chan Event, 1)
	go consumer1(e)
	// go consumer2(e)
	// go consumer3(e)

	go func() {
		for {
			e <- Event{}
			time.Sleep(time.Millisecond * 1)
		}
	}()

	go func() {
		for {
			xruntime.PrintAlloc()
			time.Sleep(time.Second)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// Output:
	// 1.Consumer1:
	// 2023/11/12 16:04:07 mem alloc: 148 KB
	// 2023/11/12 16:04:08 mem alloc: 344 KB
	// 2023/11/12 16:04:09 mem alloc: 548 KB
	// 2023/11/12 16:04:10 mem alloc: 740 KB
	// 2023/11/12 16:04:11 mem alloc: 927 KB
	// 2023/11/12 16:04:12 mem alloc: 1130 KB
	// 2023/11/12 16:04:13 mem alloc: 1327 KB
	// 2023/11/12 16:04:14 mem alloc: 1519 KB
	// 2023/11/12 16:04:15 mem alloc: 1756 KB
	// 2023/11/12 16:04:16 mem alloc: 1934 KB
	// 2023/11/12 16:04:17 mem alloc: 2122 KB
	// 2023/11/12 16:04:18 mem alloc: 2308 KB
	// 2023/11/12 16:04:19 mem alloc: 2534 KB
	// 2023/11/12 16:04:20 mem alloc: 2745 KB
	// 2023/11/12 16:04:21 mem alloc: 2916 KB
	// 2023/11/12 16:04:22 mem alloc: 3088 KB

	// 2.Consumer2:
	// 2023/11/12 16:02:35 mem alloc: 150 KB
	// 2023/11/12 16:02:36 mem alloc: 445 KB
	// 2023/11/12 16:02:37 mem alloc: 730 KB
	// 2023/11/12 16:02:38 mem alloc: 1017 KB
	// 2023/11/12 16:02:39 mem alloc: 1303 KB
	// 2023/11/12 16:02:40 mem alloc: 1589 KB
	// 2023/11/12 16:02:41 mem alloc: 1874 KB
	// 2023/11/12 16:02:42 mem alloc: 2161 KB
	// 2023/11/12 16:02:43 mem alloc: 2447 KB
	// 2023/11/12 16:02:44 mem alloc: 2735 KB
	// 2023/11/12 16:02:45 mem alloc: 3021 KB
	// 2023/11/12 16:02:46 mem alloc: 3307 KB
	// 2023/11/12 16:02:47 mem alloc: 3595 KB
	// 2023/11/12 16:02:48 mem alloc: 203 KB
	// 2023/11/12 16:02:49 mem alloc: 492 KB
	// 2023/11/12 16:02:50 mem alloc: 778 KB
	// 2023/11/12 16:02:51 mem alloc: 1064 KB

	// 3.Consumer3:
	// 2023/11/12 16:03:30 mem alloc: 150 KB
	// 2023/11/12 16:03:31 mem alloc: 156 KB
	// 2023/11/12 16:03:32 mem alloc: 156 KB
	// 2023/11/12 16:03:33 mem alloc: 156 KB
	// 2023/11/12 16:03:34 mem alloc: 156 KB
	// 2023/11/12 16:03:35 mem alloc: 156 KB
	// 2023/11/12 16:03:36 mem alloc: 156 KB
	// 2023/11/12 16:03:37 mem alloc: 156 KB
	// 2023/11/12 16:03:38 mem alloc: 156 KB
	// 2023/11/12 16:03:39 mem alloc: 156 KB
}

type Event struct{}

func handle(event Event) {}

// memory leak
func consumer1(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

// fixed
func consumer2(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			handle(event)
		case <-ctx.Done():
			cancel()
			log.Println("warning: no messages received")
		}
	}
}

// fixed
func consumer3(ch <-chan Event) {
	timer := time.NewTimer(time.Hour)
	for {
		timer.Reset(time.Hour)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}
