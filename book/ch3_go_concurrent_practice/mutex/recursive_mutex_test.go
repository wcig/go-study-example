package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/petermattis/goid"
)

func TestRecursiveMutex(t *testing.T) {
	mu := &RecursiveMutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo(mu)
		}()
	}
	wg.Wait()
	fmt.Println("over")
}

func foo(mu *RecursiveMutex) {
	mu.Lock()
	defer mu.Unlock()
	bar(mu)
}

func bar(mu *RecursiveMutex) {
	mu.Lock()
	defer mu.Unlock()
}

type RecursiveMutex struct {
	mu    sync.Mutex // 锁
	id    int64      // 持有者
	count int32      // 重入次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.id) == gid {
		m.count++
		return
	}

	m.mu.Lock()
	atomic.StoreInt64(&m.id, gid)
	m.count = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.id) != gid {
		panic(fmt.Sprintf("current goroutine [%d] not mutex owner: [%d]", m.id, gid))
	}

	m.count--
	if atomic.LoadInt32(&m.count) == 0 {
		atomic.StoreInt64(&m.id, -1)
		m.mu.Unlock()
	}
}
