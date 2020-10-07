package ch14_concurrent

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 设定CPU使用个数
func TestGOPROCES(t *testing.T) {
	num := runtime.NumCPU()
	fmt.Println("this computer cpu num:", num) // 输出当前电脑的CPU个数
	runtime.GOMAXPROCS(1)                      // 设定当前使用1个CPU
}

// 创建goroutine
func TestCreateGoroutine(t *testing.T) {
	go printChar('a', 'z')
	go printChar('A', 'Z')
	time.Sleep(2 * time.Second)
}

func printChar(start byte, end byte) {
	b := start
	for {
		fmt.Printf("%c", b)
		b++
		if b > end {
			break
		}
	}
}

// sync.WaitGroup
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printChar('a', 'z')
	}()
	go func() {
		defer wg.Done()
		printChar('A', 'Z')
	}()

	wg.Wait()
}

// 共享资源没有加锁
type unsafeCounter struct {
	counter int
}

func (c *unsafeCounter) Incr(n int) {
	c.counter += n
}

func (c *unsafeCounter) Value() int {
	return c.counter
}

func TestWithoutMutex(t *testing.T) {
	c := unsafeCounter{counter: 0}
	for i := 0; i < 1000; i++ {
		go c.Incr(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("result value:", c.Value()) // result value: 478684
}

// sync.Mutex
type safeCounter struct {
	counter int
	mux     sync.Mutex
}

func (c *safeCounter) Incr(n int) {
	c.mux.Lock()
	c.counter += n
	c.mux.Unlock()
}

func (c *safeCounter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.counter
}

func TestMutex(t *testing.T) {
	c := safeCounter{counter: 0}
	for i := 0; i < 1000; i++ {
		go c.Incr(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("result value:", c.Value()) // result value: 499500
}

// sync.RWmutex
type rwCounter struct {
	counter int
	mux     sync.RWMutex
}

func (c *rwCounter) Incr(n int) {
	c.mux.Lock()
	c.counter += n
	c.mux.Unlock()
}

func (c *rwCounter) Value() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.counter
}

func TestRWmutex(t *testing.T) {
	c := rwCounter{counter: 0}
	for i := 0; i < 1000; i++ {
		go c.Incr(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("result value:", c.Value()) // result value: 499500
}

// sync.Once
func TestOnce(t *testing.T) {
	var one sync.Once
	for i := 0; i < 100; i++ {
		one.Do(func() {
			fmt.Println("one done")
		})
	}
}

// output:
// one done

// atomic
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
	fmt.Println("result counter:", counter, atomic.LoadInt64(&counter))

	var num int64
	fmt.Println(atomic.SwapInt64(&num, 1))
	fmt.Println(num)

	fmt.Println(atomic.CompareAndSwapInt64(&num, 1, 20))
	fmt.Println(num)
}

// output:
// result counter: 4 4

// 并发读写操作map将导致错误
func TestMap1(t *testing.T) {
	m := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				m[strconv.Itoa(j)] = j
			}
		}()
	}

	wg.Wait()
	fmt.Println("map:", m)
}

// output:
// fatal error: concurrent map writes

// sync.Map
func TestMap2(t *testing.T) {
	var m sync.Map
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				m.Store(strconv.Itoa(j), j)
			}
		}()
	}

	wg.Wait()
	val, ok := m.Load("1")
	fmt.Println(val, ok)

	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

// output:
// 1 true
// 1 1
// 2 2
// 3 3
// 6 6
// 7 7
// 8 8
// 4 4
// 5 5
// 9 9
// 10 10

// sync.Cond
func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go listen(c)
	}

	time.Sleep(3 * time.Second)
	go broadcast(c)

	time.Sleep(1 * time.Second)
	fmt.Println("over...")
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	c.Broadcast()
	fmt.Println("broadcast...")
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println("listen...")
	c.L.Unlock()
}

// output:
// broadcast...
// listen...
// listen...
// listen...
// listen...
// listen...
// listen...
// listen...
// listen...
// listen...
// listen...
// over...

// sync.Pool
func TestPool(t *testing.T) {
	var p sync.Pool
	for i := 1; i <= 10; i++ {
		p.Put(i)
	}

	var result []int
	for {
		val := p.Get()
		if val == nil {
			break
		}
		result = append(result, val.(int))
	}
	fmt.Println("result:", result)
}

// output:
// result: [1 10 9 8 7 6 5 4 3 2]
