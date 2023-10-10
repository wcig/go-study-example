package ch100_redis

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

func TestSingleInstanceLock(t *testing.T) {
	initRedis()

	mu := NewSingleInstanceMutex("111", client)
	if err := mu.Lock(time.Minute); err != nil {
		log.Fatalf(">> lock err: %v", err)
	}

	time.Sleep(30 * time.Second)

	ok, err := mu.Unlock()
	if err != nil {
		log.Fatalf(">> unlock err: %v", err)
	}
	log.Printf(">> unlock result: %v", ok)
}

type SingleInstanceMutex struct {
	Key string
	Val string
	c   *redis.Client
}

func NewSingleInstanceMutex(name string, client *redis.Client) *SingleInstanceMutex {
	return &SingleInstanceMutex{
		Key: name,
		Val: uuid.NewV4().String(),
		c:   client,
	}
}

func (m *SingleInstanceMutex) Lock(expiry time.Duration) error {
	result, err := m.c.SetNX(m.Key, m.Val, expiry).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.New("lock failed err")
	}
	return nil
}

func (m *SingleInstanceMutex) Unlock() (bool, error) {
	const luaScript = `
	if redis.call("get",KEYS[1]) == ARGV[1] then
	    return redis.call("del",KEYS[1])
	else
	    return 0
	end`

	result, err := m.c.Eval(luaScript, []string{m.Key}, []string{m.Val}).Result()
	if err != nil {
		return false, err
	}
	if num := result.(int64); num == 0 {
		return false, nil
	}
	// num == 1
	return true, nil
}
