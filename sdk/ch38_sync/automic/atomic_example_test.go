package automic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for count := 0; count < 2; count++ {
				atomic.AddInt64(&counter, 1)
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()
	fmt.Println("result:", counter, atomic.LoadInt64(&counter)) // result: 4 4
}
