package ch14_concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// time.After
func TestTimeAfter(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
		close(c1)
	}()

	for {
		select {
		case res, ok := <-c1:
			if ok {
				fmt.Println(res)
			} else {
				fmt.Println("empty")
				goto End
			}
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		}
	}
End:
	fmt.Println("over...")
}

// output:
// timeout 1
// result 1
// empty
// over...

// time.NewTimer
func TestTimer(t *testing.T) {
	n := 2
	timer := time.NewTimer(time.Duration(n) * time.Second)
	tt := <-timer.C
	fmt.Printf("%d seconds expire, tt: %s\n", n, tt.Format("15:04:05"))
	fmt.Println("over...")
}

// output:
// 2 seconds expire, tt: 22:38:59
// over...

// Time.Ticker
func TestTicker(t *testing.T) {
	timer := time.NewTicker(time.Second)

	n := 0
	for {
		tt := <-timer.C
		fmt.Println(tt.Format("15:04:05"))

		n++
		if n >= 3 {
			timer.Stop()
			break
		}
	}
	fmt.Println("over...")
}

// output:
// 23:06:28
// 23:06:29
// 23:06:30
// over...

func Test1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调  //整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()
}
