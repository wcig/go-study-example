package once

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	cacheOnce = sync.Once{}
	cacheVal  = ""
)

func TestOnce(t *testing.T) {
	initCache := func() string {
		return "cache"
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cacheOnce.Do(func() {
				cacheVal = initCache()
				fmt.Println("init cache success..")
			})
		}()
	}
	wg.Wait()
	fmt.Println(cacheVal)
	// Output:
	// init cache success..
	// cache
}

type ErrOnce struct {
	done uint32
	m    sync.Mutex
}

func (o *ErrOnce) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 0 {
		return o.doSlow(f)
	}
	return nil
}

func (o *ErrOnce) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		err := f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
		return err
	}
	return nil
}

func TestErrOnce(t *testing.T) {
	var n uint32
	fn := func() error {
		var tmp uint32
		if tmp = atomic.AddUint32(&n, 1); tmp > 3 {
			return nil
		}
		return fmt.Errorf("number not reached: %d", tmp)
	}

	once := &ErrOnce{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := once.Do(fn)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("do once success..")
			}
		}()
	}
	wg.Wait()
	// Output:
	// number not reached: 1
	// number not reached: 2
	// number not reached: 3
	// do once success..
	// do once success..
	// do once success..
	// do once success..
	// do once success..
	// do once success..
	// do once success..
}

// 错误使用1: 嵌套调用Do方法
func TestOnceWrong1(t *testing.T) {
	f := func() {
		fmt.Println("ok")
	}
	once := &sync.Once{}
	once.Do(func() {
		once.Do(f)
	})
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// 错误使用2: 调用Do方法初始化资源时未判断失败情况
func TestOnceWrong2(t *testing.T) {
	initConn := func() (conn interface{}, err error) {
		return nil, errors.New("init conn failed")
	}

	var once sync.Once
	var conn interface{}
	once.Do(func() {
		conn, _ = initConn()
	})
	// conn未初始化成功就使用
	_ = conn
}

// 错误使用3: Once使用时进行重置为一个新的Once导致解锁失败错误
func TestOnceWrong3(t *testing.T) {
	once := &sync.Once{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				*once = sync.Once{}
				fmt.Println("ok")
			})
		}()
	}
	wg.Wait()
	// Output:
	// fatal error: sync: unlock of unlocked mutex
}
