package ch100_redis

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
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
	client.Set(key, val, -1)
	result, err := client.Get(key).Result()
	assert.True(t, err == nil)
	assert.True(t, result == val)
}
