package ch100_redis

import (
	"log"
	"testing"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

var rs *redsync.Redsync

// 注意: 这里只有初始化多个client, pool并传入pool列表时,才会按redlock算法往往多个示例写入锁
func initRedsync() {
	pool := goredis.NewPool(client)
	rs = redsync.New(pool)
}

func TestRedlock(t *testing.T) {
	initRedis()
	initRedsync()

	mutexName := "my-global-mutex"
	mutex := rs.NewMutex(mutexName, redsync.WithExpiry(time.Minute), redsync.WithTries(3))

	if err := mutex.Lock(); err != nil {
		log.Fatalf("redlock lock err: %v", err)
	} else {
		log.Println("redlock lock success")
	}

	time.Sleep(30 * time.Second)

	ok, err := mutex.Unlock()
	if err != nil {
		log.Fatalf("redlock unlock err: %v", err)
	}
	log.Printf("redlock unlock result: %v", ok)
}
