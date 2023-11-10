package ch38_sync

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	f := func() {
		fmt.Println("only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(f)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// output:
	// only once
}

func TestMap(t *testing.T) {
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
	// output:
	// 1 true
	// 1 1
	// 2 2
	// 4 4
	// 6 6
	// 8 8
	// 3 3
	// 5 5
	// 7 7
	// 9 9
	// 10 10
}

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
	fmt.Println("result:", result) // result: [1 10 9 8 7 6 5 4 3 2]
}

func TestPoolWithNew(t *testing.T) {
	const num = 10
	p := &sync.Pool{
		New: func() any {
			return -rand.Intn(10)
		},
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		v := i
		go func() {
			p.Put(v)
			wg.Done()
		}()
	}

	wg.Wait()
	var result []int
	for {
		val := p.Get()
		if val == nil || len(result) == 20 {
			break
		}
		result = append(result, val.(int))
	}
	fmt.Println("result:", result) // result: [0 4 3 1 2 8 7 6 5 9 -2 -8 -7 -4 -7 -8 -1 -3 -3 -5]
}

type counter struct {
	num int
	sync.Mutex
}

func (c *counter) incr(n int) {
	c.Lock()
	c.num += n
	c.Unlock()
}

func (c *counter) value() int {
	c.Lock()
	defer c.Unlock()
	return c.num
}

func TestMutex(t *testing.T) {
	var c counter
	for i := 0; i < 10000; i++ {
		go c.incr(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(c.value()) // 49995000
}

type rwCounter struct {
	num int
	sync.RWMutex
}

func (c *rwCounter) incr(n int) {
	c.Lock()
	c.num += n
	c.Unlock()
}

func (c *rwCounter) value() int {
	c.RLock()
	defer c.RUnlock()
	return c.num
}

func TestRWMutex(t *testing.T) {
	var c counter
	for i := 0; i < 10000; i++ {
		go c.incr(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(c.value()) // 49995000
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("1 done")
		wg.Done()
	}()

	go func() {
		fmt.Println("2 done")
		wg.Done()
	}()

	wg.Wait()
	// output:
	// 2 done
	// 1 done
}
