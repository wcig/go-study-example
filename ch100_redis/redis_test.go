package ch100_redis

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func TestString(t *testing.T) {
	key, val := "hello", "world"
	result, err := client.Set(key, val, -1).Result()
	fmt.Println(result, err)

	result, err = client.Get(key).Result()
	fmt.Println(result, err)
	// assert.True(t, err == nil)
	// assert.True(t, result == val)
}

func TestMSet(t *testing.T) {
	result, err := client.MSet("one", "1", "two", "2", "three", "3", "four", "4").Result()
	fmt.Println(result, err) // OK <nil>
}

func TestMGet(t *testing.T) {
	result, err := client.MGet("one", "two+", "three").Result()
	fmt.Println(result, err) // [1 <nil> 3] <nil>

	for i, val := range result {
		fmt.Println(i, val, val == nil)
	}
}

func TestMGetMSet(t *testing.T) {
	// var pairs []interface{}
	// pairs = append(pairs, "one", "1")
	// pairs = append(pairs, "two", "2")
	// pairs = append(pairs, "three", "3")
	// pairs = append(pairs, "four", "4")
	// result, err := client.MSet(pairs...).Result()
	// fmt.Println(result, err)

	keys := []string{"one", "two", "three", "four"}
	result, err := client.MGet(keys...).Result()
	fmt.Println(result, err)
}
