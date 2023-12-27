package single_call

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	s := &SingleCall{}
	wg := sync.WaitGroup{}
	const num = 10000

	wg.Add(num)
	var n int32
	for i := 0; i < num; i++ {
		go func() {
			s.Do(func() {
				atomic.AddInt32(&n, 1)
				time.Sleep(time.Millisecond * 10)
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n) // 1

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			s.Do(func() {
				atomic.AddInt32(&n, 1)
				time.Sleep(time.Millisecond * 10)
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n) // 2
}
