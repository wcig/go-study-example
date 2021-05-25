package ch100_redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func TestString(t *testing.T) {
	initRedis()

	key, val := "hello", "world"
	result, err := client.Set(key, val, -1).Result()
	fmt.Println(result, err) // OK <nil>

	result, err = client.Get(key).Result()
	fmt.Println(result, err) // world <nil>

	num, err := client.Del(key).Result()
	fmt.Println(num, err) // 1 <nil>
}

func TestMSet(t *testing.T) {
	initRedis()

	result, err := client.MSet("one", "1", "two", "2", "three", "3", "four", "4").Result()
	fmt.Println(result, err) // OK <nil>
}

func TestMGet(t *testing.T) {
	initRedis()

	result, err := client.MGet("one", "two+", "three").Result()
	fmt.Println(result, err) // [1 <nil> 3] <nil>

	for i, val := range result {
		fmt.Println(i, val, val == nil)
	}
}

func TestExpire(t *testing.T) {
	initRedis()

	key, val := "hello", "world"
	ok, err := client.Expire(key, time.Second*1000).Result()
	fmt.Println(ok, err) // false <nil>

	client.Set(key, val, -1)
	ok, err = client.Expire(key, time.Second*1000).Result()
	fmt.Println(ok, err) // true <nil>

	client.Del(key)
}
