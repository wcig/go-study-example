package single_flight

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

type User struct {
	ID   string
	Name string
}

var (
	cacheUser atomic.Value

	cacheTotal atomic.Int64
	cacheHint  atomic.Int64
	cacheMiss  atomic.Int64
)

func GetUser(id string) *User {
	cacheTotal.Add(1)
	if user := GetUserFromCache(id); user != nil {
		cacheHint.Add(1)
		return user
	}

	cacheMiss.Add(1)
	user := GetUserFromDB(id)
	SetUserCache(user)
	return user
}

func SetUserCache(user *User) {
	time.Sleep(10 * time.Millisecond)
	cacheUser.Store(user)
}

func GetUserFromCache(id string) *User {
	time.Sleep(10 * time.Millisecond)
	if val := cacheUser.Load(); val != nil {
		user, _ := val.(*User)
		return user
	}
	return nil
}

func GetUserFromDB(id string) *User {
	time.Sleep(100 * time.Millisecond)
	return &User{ID: "123", Name: "tom"}
}

func TestConcurrentNoSingleFlight(t *testing.T) {
	const concurrentNum = 10000

	start := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(concurrentNum)

	for i := 0; i < concurrentNum; i++ {
		go func() {
			_ = GetUser("123")
			wg.Done()
		}()
		// time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
	latency := time.Since(start)
	log.Printf(">> over, time cost: %s, query total: %d, cache stat total: %d, hint: %d, miss: %d",
		latency, concurrentNum, cacheTotal.Load(), cacheHint.Load(), cacheMiss.Load())

	// Output:
	// 2023/10/14 12:52:52 >> over, time cost: 133.614917ms, query total: 10000, cache stat total: 10000, hint: 0, miss: 10000
}

func TestConcurrentWithSingleFlight(t *testing.T) {
	const concurrentNum = 10000

	start := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(concurrentNum)
	sg := &singleflight.Group{}

	for i := 0; i < concurrentNum; i++ {
		go func() {
			_, err, _ := sg.Do("GetUser", func() (interface{}, error) {
				user := GetUser("123")
				return user, nil
			})
			if err != nil {
				log.Fatalf(">> get user err: %v", err)
			}
			wg.Done()
		}()
		// time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
	latency := time.Since(start)
	log.Printf(">> over, time cost: %s, query total: %d, cache stat total: %d, hint: %d, miss: %d",
		latency, concurrentNum, cacheTotal.Load(), cacheHint.Load(), cacheMiss.Load())

	// Output:
	// 2023/10/14 12:53:05 >> over, time cost: 125.019792ms, query total: 10000, cache stat total: 1, hint: 0, miss: 1
}
