package ch100_redis

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

func initRedis() {
	// redis 7.x
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func TestString(t *testing.T) {
	initRedis()

	key, val := "hello", "world"
	result, err := client.Set(ctx, key, val, -1).Result()
	fmt.Println(result, err) // OK <nil>

	result, err = client.Get(ctx, key).Result()
	fmt.Println(result, err) // world <nil>

	num, err := client.Del(ctx, key).Result()
	fmt.Println(num, err) // 1 <nil>
}

func TestMSet(t *testing.T) {
	initRedis()

	result, err := client.MSet(ctx, "one", "1", "two", "2", "three", "3", "four", "4").Result()
	fmt.Println(result, err) // OK <nil>
}

func TestMGet(t *testing.T) {
	initRedis()

	result, err := client.MGet(ctx, "one", "two+", "three").Result()
	fmt.Println(result, err) // [1 <nil> 3] <nil>

	for i, val := range result {
		fmt.Println(i, val, val == nil)
	}
}

func TestExpire(t *testing.T) {
	initRedis()

	key, val := "hello", "world"
	ok, err := client.Expire(ctx, key, time.Second*1000).Result()
	fmt.Println(ok, err) // false <nil>

	client.Set(ctx, key, val, -1)
	ok, err = client.Expire(ctx, key, time.Second*1000).Result()
	fmt.Println(ok, err) // true <nil>

	client.Del(ctx, key)
}

func TestPipeline(t *testing.T) {
	initRedis()

	// set
	{
		p := client.Pipeline()
		p.Set(ctx, "aaa", 111, -1)
		p.Set(ctx, "bbb", 222, -1)
		p.Set(ctx, "ccc", 333, -1)
		cmds, err := p.Exec(ctx)
		if err != nil {
			log.Fatalf(">> set: pipeline exec err: %v", err)
		}
		for i, v := range cmds {
			val, err2 := v.(*redis.StatusCmd).Result()
			log.Printf(">> set: pipeline exec %d result: %s, %v", i+1, val, err2)
		}
	}

	// get
	{
		p := client.Pipeline()
		p.Get(ctx, "aaa")
		p.Get(ctx, "bbb")
		p.Get(ctx, "ccc")
		cmds, err := p.Exec(ctx)
		if err != nil {
			log.Fatalf(">> get: pipeline exec err: %v", err)
		}
		for i, v := range cmds {
			val, err2 := v.(*redis.StringCmd).Result()
			log.Printf(">> get: pipeline exec %d result: %s, %v", i+1, val, err2)
		}
	}

	// del
	{
		p := client.Pipeline()
		p.Del(ctx, "aaa")
		p.Del(ctx, "bbb")
		p.Del(ctx, "ccc")
		cmds, err := p.Exec(ctx)
		if err != nil {
			log.Fatalf(">> get: pipeline exec err: %v", err)
		}
		for i, v := range cmds {
			val, err2 := v.(*redis.IntCmd).Result()
			log.Printf(">> get: pipeline exec %d result: %d, %v", i+1, val, err2)
		}
	}
}
