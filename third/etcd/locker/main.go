package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr     = flag.String("addr", "http://localhost:2379", "etcd addresses")
	lockName = flag.String("name", "my-test-lock", "lock name")
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
	locker := concurrency.NewLocker(ss, *lockName)

	// 加锁
	log.Println("acquiring etcd lock")
	locker.Lock()
	log.Println("acquired etcd lock")

	// 解锁
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	locker.Unlock()
	log.Println("released etcd lock")
}
