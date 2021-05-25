package ch100_redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var clusterClient *redis.ClusterClient

func init() {
	addrs := []string{
		"10.200.50.49:7001",
		"10.200.50.49:7002",
		"10.200.50.49:7003",
		"10.200.50.49:7004",
		"10.200.50.49:7005",
		"10.200.50.49:7006",
	}
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           addrs,
		DialTimeout:     60 * time.Second,
		ReadTimeout:     60 * time.Second,
		WriteTimeout:    60 * time.Second,
		PoolSize:        1000,
		Password:        "",
		MaxRetries:      2,
		MinRetryBackoff: -1,
		MaxRetryBackoff: -1,
	})

	pong, err := clusterClient.Ping().Result()
	fmt.Println(pong, err)
}

func TestClusterMGet(t *testing.T) {
	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	result := multiGet(keys...)
	for i := range result {
		assert.True(t, result[i] == vals[i])
	}
}

func TestClusterMSet(t *testing.T) {
	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	var pairs []interface{}
	for i := range keys {
		pairs = append(pairs, keys[i], vals[i])
	}
	multiSet(-1, pairs...)
}

func BenchmarkMGet(b *testing.B) {
	keys := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		multiGet(keys...)
	}
}

func BenchmarkGet(b *testing.B) {
	keys := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			clusterClient.Get(key)
		}
	}
}

func BenchmarkMSet(b *testing.B) {
	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	var pairs []interface{}
	for i := range keys {
		pairs = append(pairs, keys[i], vals[i])
	}

	for i := 0; i < b.N; i++ {
		multiSet(-1, pairs)
	}
}

func BenchmarkSet(b *testing.B) {
	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	var pairs []interface{}
	for i := range keys {
		pairs = append(pairs, keys[i], vals[i])
	}

	for i := 0; i < b.N; i++ {
		for i := range keys {
			clusterClient.Set(keys[i], vals[i], -1)
		}
	}
}

func multiGet(keys ...string) []interface{} {
	size := len(keys)
	if size == 0 {
		return []interface{}{}
	}

	tasks := make([]*multiTask, size)
	for i := range keys {
		task := &multiTask{
			key:   keys[i],
			reply: nil,
			err:   nil,
			done:  make(chan int),
		}
		tasks[i] = task
	}

	for i := range tasks {
		go handleGetTask(tasks[i])
	}

	for i := range tasks {
		<-tasks[i].done
	}

	result := make([]interface{}, size)
	for i := range tasks {
		if tasks[i].err == nil {
			result[i] = tasks[i].reply
		}
	}
	return result
}

func multiSet(expiration time.Duration, pairs ...interface{}) {
	size := len(pairs)
	if size == 0 || size%2 != 0 {
		return
	}

	tasks := make([]*multiTask, size/2)
	for i := 0; i < size/2; i++ {
		task := &multiTask{
			key:        pairs[2*i].(string),
			val:        pairs[2*i+1],
			expiration: expiration,
			reply:      nil,
			err:        nil,
			done:       make(chan int),
		}
		tasks[i] = task
	}

	for i := range tasks {
		go handleSetTask(tasks[i])
	}

	for i := range tasks {
		<-tasks[i].done
	}
}

type multiTask struct {
	key        string
	val        interface{}
	expiration time.Duration
	reply      interface{}
	err        error
	done       chan int
}

func handleGetTask(task *multiTask) {
	task.reply, task.err = clusterClient.Get(task.key).Result()
	task.done <- 1
}

func handleSetTask(task *multiTask) {
	task.reply, task.err = clusterClient.Set(task.key, task.val, task.expiration).Result()
	task.done <- 1
}
