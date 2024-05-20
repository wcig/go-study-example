package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr     = flag.String("addr", "http://localhost:2379", "etcd addresses")
	lockName = flag.String("name", "my-test-mutex", "lock name")
	action   = flag.String("rw", "w", "r acquiring read lock, w acquiring write lock")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 创建client
	ec, err := clientv3.New(clientv3.Config{Endpoints: []string{*addr}})
	if err != nil {
		log.Fatalf("init etcd client err: %v", err)
	}
	defer func() {
		if err = ec.Close(); err != nil {
			log.Fatalf("close etcd client err: %v", err)
		}
	}()

	// 创建session
	ss, err := concurrency.NewSession(ec)
	if err != nil {
		log.Fatalf("init etcd session err: %v", err)
	}
	defer func() {
		if err = ss.Close(); err != nil {
			log.Fatalf("close etcd session err: %v", err)
		}
	}()

	// 创建锁
	mu := recipe.NewRWMutex(ss, *lockName)

	switch *action {
	case "r":
		testReadLock(mu)
	case "w":
		testWriteLock(mu)
	default:
		log.Fatalln("unknown action")
	}
}

func testReadLock(mu *recipe.RWMutex) {
	// 加锁
	log.Println("acquiring etcd read lock")
	if err := mu.RLock(); err != nil {
		log.Printf("acquired etcd read lock, err: %v", err)
	}
	log.Println("acquired etcd read lock")

	// 解锁
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	if err := mu.RUnlock(); err != nil {
		log.Printf("released etcd read lock, err: %v", err)
	}
	log.Println("released etcd read lock")
}

func testWriteLock(mu *recipe.RWMutex) {
	// 加锁
	log.Println("acquiring etcd write lock")
	if err := mu.Lock(); err != nil {
		log.Printf("acquired etcd write lock, err: %v", err)
	}
	log.Println("acquired etcd write lock")

	// 解锁
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	if err := mu.Unlock(); err != nil {
		log.Printf("released etcd write lock, err: %v", err)
	}
	log.Println("released etcd write lock")
}
