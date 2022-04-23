package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestTokenRecursiveMutex(t *testing.T) {
	mu := &TokenRecursiveMutex{}
	wg := &sync.WaitGroup{}
	for i := int64(0); i < 100; i++ {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			foo2(mu, n)
		}(i)
	}
	wg.Wait()
	fmt.Println("over")
}

func foo2(mu *TokenRecursiveMutex, token int64) {
	mu.Lock(token)
	defer mu.Unlock(token)
	bar2(mu, token)
}

func bar2(mu *TokenRecursiveMutex, token int64) {
	mu.Lock(token)
	defer mu.Unlock(token)
}

type TokenRecursiveMutex struct {
	mu    sync.Mutex // 锁
	token int64      // 持有者
	count int32      // 重入次数
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.count++
		return
	}

	m.mu.Lock()
	atomic.StoreInt64(&m.token, token)
	m.count = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("current goroutine [%d] not mutex owner: [%d]", m.token, token))
	}

	m.count--
	if atomic.LoadInt32(&m.count) == 0 {
		atomic.StoreInt64(&m.token, -1)
		m.mu.Unlock()
	}
}
