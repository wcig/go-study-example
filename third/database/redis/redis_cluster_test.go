package redis

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var clusterClient *redis.ClusterClient

func initClusterClient() {
	// redis 7.x
	addrs := []string{
		"moss:7001",
		"moss:7002",
		"moss:7003",
		"moss:7004",
		"moss:7005",
		"moss:7006",
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

	pong, err := clusterClient.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	nodesResult, err := clusterClient.ClusterNodes(context.Background()).Result()
	fmt.Println(nodesResult, err)

	// Output:
	// PONG <nil>
	// 93ac0e30276bcf6c4c539151dd8d44c47dab9d55 192.168.100.254:7002@17002 myself,master - 0 1698374812000 10 connected 5461-10922
	// 521c5afb196fb3cc57dca9bcf4d27ba507895036 192.168.100.254:7001@17001 slave 5215a0dae794f6f031138350ee1a47c9946eabaf 0 1698374813309 8 connected
	// 45999feab9118d11baceb973f13ccf81eb43b82f 192.168.100.254:7006@17006 master - 0 1698374813007 9 connected 10923-16383
	// 5215a0dae794f6f031138350ee1a47c9946eabaf 192.168.100.254:7004@17004 master - 0 1698374813108 8 connected 0-5460
	// abbd1de6b3dd0ada65095f891f7e55cf62d44a72 192.168.100.254:7003@17003 slave 45999feab9118d11baceb973f13ccf81eb43b82f 0 1698374813309 9 connected
	// d1041ab792a240ac3815fd809cb5526b1b3dfef1 192.168.100.254:7005@17005 slave 93ac0e30276bcf6c4c539151dd8d44c47dab9d55 0 1698374813108 10 connected
	// <nil>
}

func TestClusterMGet(t *testing.T) {
	initClusterClient()

	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	result := multiGet(keys...)
	for i := range result {
		assert.True(t, result[i] == vals[i])
	}
}

func TestClusterMSet(t *testing.T) {
	initClusterClient()

	keys := []string{"one", "two", "three", "four", "five"}
	vals := []interface{}{"1", "2", "3", "4", "5"}

	var pairs []interface{}
	for i := range keys {
		pairs = append(pairs, keys[i], vals[i])
	}
	multiSet(-1, pairs...)
}

func BenchmarkMGet(b *testing.B) {
	initClusterClient()

	keys := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		multiGet(keys...)
	}
}

func BenchmarkGet(b *testing.B) {
	initClusterClient()

	keys := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			clusterClient.Get(ctx, key)
		}
	}
}

func BenchmarkMSet(b *testing.B) {
	initClusterClient()

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
			clusterClient.Set(ctx, keys[i], vals[i], -1)
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
	task.reply, task.err = clusterClient.Get(ctx, task.key).Result()
	task.done <- 1
}

func handleSetTask(task *multiTask) {
	task.reply, task.err = clusterClient.Set(ctx, task.key, task.val, task.expiration).Result()
	task.done <- 1
}

// 代码来源: github.com/go-redis/redis@v6.15.9+incompatible/internal/hashtag/hashtag.go:64
func TestSlotCalculate(t *testing.T) {
	var crc16tab = [256]uint16{
		0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5, 0x60c6, 0x70e7,
		0x8108, 0x9129, 0xa14a, 0xb16b, 0xc18c, 0xd1ad, 0xe1ce, 0xf1ef,
		0x1231, 0x0210, 0x3273, 0x2252, 0x52b5, 0x4294, 0x72f7, 0x62d6,
		0x9339, 0x8318, 0xb37b, 0xa35a, 0xd3bd, 0xc39c, 0xf3ff, 0xe3de,
		0x2462, 0x3443, 0x0420, 0x1401, 0x64e6, 0x74c7, 0x44a4, 0x5485,
		0xa56a, 0xb54b, 0x8528, 0x9509, 0xe5ee, 0xf5cf, 0xc5ac, 0xd58d,
		0x3653, 0x2672, 0x1611, 0x0630, 0x76d7, 0x66f6, 0x5695, 0x46b4,
		0xb75b, 0xa77a, 0x9719, 0x8738, 0xf7df, 0xe7fe, 0xd79d, 0xc7bc,
		0x48c4, 0x58e5, 0x6886, 0x78a7, 0x0840, 0x1861, 0x2802, 0x3823,
		0xc9cc, 0xd9ed, 0xe98e, 0xf9af, 0x8948, 0x9969, 0xa90a, 0xb92b,
		0x5af5, 0x4ad4, 0x7ab7, 0x6a96, 0x1a71, 0x0a50, 0x3a33, 0x2a12,
		0xdbfd, 0xcbdc, 0xfbbf, 0xeb9e, 0x9b79, 0x8b58, 0xbb3b, 0xab1a,
		0x6ca6, 0x7c87, 0x4ce4, 0x5cc5, 0x2c22, 0x3c03, 0x0c60, 0x1c41,
		0xedae, 0xfd8f, 0xcdec, 0xddcd, 0xad2a, 0xbd0b, 0x8d68, 0x9d49,
		0x7e97, 0x6eb6, 0x5ed5, 0x4ef4, 0x3e13, 0x2e32, 0x1e51, 0x0e70,
		0xff9f, 0xefbe, 0xdfdd, 0xcffc, 0xbf1b, 0xaf3a, 0x9f59, 0x8f78,
		0x9188, 0x81a9, 0xb1ca, 0xa1eb, 0xd10c, 0xc12d, 0xf14e, 0xe16f,
		0x1080, 0x00a1, 0x30c2, 0x20e3, 0x5004, 0x4025, 0x7046, 0x6067,
		0x83b9, 0x9398, 0xa3fb, 0xb3da, 0xc33d, 0xd31c, 0xe37f, 0xf35e,
		0x02b1, 0x1290, 0x22f3, 0x32d2, 0x4235, 0x5214, 0x6277, 0x7256,
		0xb5ea, 0xa5cb, 0x95a8, 0x8589, 0xf56e, 0xe54f, 0xd52c, 0xc50d,
		0x34e2, 0x24c3, 0x14a0, 0x0481, 0x7466, 0x6447, 0x5424, 0x4405,
		0xa7db, 0xb7fa, 0x8799, 0x97b8, 0xe75f, 0xf77e, 0xc71d, 0xd73c,
		0x26d3, 0x36f2, 0x0691, 0x16b0, 0x6657, 0x7676, 0x4615, 0x5634,
		0xd94c, 0xc96d, 0xf90e, 0xe92f, 0x99c8, 0x89e9, 0xb98a, 0xa9ab,
		0x5844, 0x4865, 0x7806, 0x6827, 0x18c0, 0x08e1, 0x3882, 0x28a3,
		0xcb7d, 0xdb5c, 0xeb3f, 0xfb1e, 0x8bf9, 0x9bd8, 0xabbb, 0xbb9a,
		0x4a75, 0x5a54, 0x6a37, 0x7a16, 0x0af1, 0x1ad0, 0x2ab3, 0x3a92,
		0xfd2e, 0xed0f, 0xdd6c, 0xcd4d, 0xbdaa, 0xad8b, 0x9de8, 0x8dc9,
		0x7c26, 0x6c07, 0x5c64, 0x4c45, 0x3ca2, 0x2c83, 0x1ce0, 0x0cc1,
		0xef1f, 0xff3e, 0xcf5d, 0xdf7c, 0xaf9b, 0xbfba, 0x8fd9, 0x9ff8,
		0x6e17, 0x7e36, 0x4e55, 0x5e74, 0x2e93, 0x3eb2, 0x0ed1, 0x1ef0,
	}
	const slotNumber = 16384

	crc16sum := func(key string) (crc uint16) {
		for i := 0; i < len(key); i++ {
			crc = (crc << 8) ^ crc16tab[(byte(crc>>8)^key[i])&0x00ff]
		}
		return
	}

	key := "hello"
	slot := int(crc16sum(key)) % slotNumber
	fmt.Println(">> slot:", slot)
}

// "github.com/redis/go-redis/v9" 实现了redis cluster对pipeline支持
func TestClusterPipeline(t *testing.T) {
	initClusterClient()

	// set
	{
		ctx := context.Background()
		p := clusterClient.Pipeline()
		p.Set(ctx, "aaa", 111, -1)
		p.Set(ctx, "bbb", 222, -1)
		p.Set(ctx, "ccc", 333, -1)
		cmds, err := p.Exec(context.Background())
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
		ctx := context.Background()
		p := clusterClient.Pipeline()
		p.Get(ctx, "aaa")
		p.Get(ctx, "bbb")
		p.Get(ctx, "ccc")
		cmds, err := p.Exec(context.Background())
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
		ctx := context.Background()
		p := clusterClient.Pipeline()
		p.Del(ctx, "aaa")
		p.Del(ctx, "bbb")
		p.Del(ctx, "ccc")
		cmds, err := p.Exec(context.Background())
		if err != nil {
			log.Fatalf(">> get: pipeline exec err: %v", err)
		}
		for i, v := range cmds {
			val, err2 := v.(*redis.IntCmd).Result()
			log.Printf(">> get: pipeline exec %d result: %d, %v", i+1, val, err2)
		}
	}
}
