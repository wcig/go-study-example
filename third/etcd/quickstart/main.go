package main

import (
	"context"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var ec *clientv3.Client

func initEtcd() {
	cfg := clientv3.Config{
		Endpoints:            []string{"localhost:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          0,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		Context:              nil,
		Logger:               nil,
		LogConfig:            nil,
		PermitWithoutStream:  false,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("init etcd client err: %v", err)
	}
	ec = client
	log.Println("init etcd client success")
}

func closeEtcd() {
	if err := ec.Close(); err != nil {
		log.Fatalf("close etcd client err: %v", err)
	}
	log.Println("close etcd client success")
}

func main() {
	initEtcd()
	defer closeEtcd()

	key, val, newVal := "hello", "world", "world2"
	ctx := context.Background()

	// put
	putResponse, err := ec.Put(ctx, key, val)
	if err != nil {
		log.Fatalf("etcd put kv err: %v", err)
	}
	_ = putResponse
	log.Printf("etcd put k: %s, v: %s success", key, val)

	// get
	getResponse, err := ec.Get(ctx, key)
	if err != nil {
		log.Fatalf("etcd put kv err: %v", err)
	}
	for _, kv := range getResponse.Kvs {
		log.Printf("etcd get k: %s, v: %s, version: %d", kv.Key, kv.Value, kv.Version)
	}

	// put (update)
	putResponse, err = ec.Put(ctx, key, newVal)
	if err != nil {
		log.Fatalf("etcd put kv err: %v", err)
	}
	log.Printf("etcd put k: %s, v: %s success", key, val)

	// get
	getResponse, err = ec.Get(ctx, key)
	if err != nil {
		log.Fatalf("etcd put kv err: %v", err)
	}
	for _, kv := range getResponse.Kvs {
		log.Printf("etcd get k: %s, v: %s, version: %d", kv.Key, kv.Value, kv.Version)
	}

	// delete
	deleteResponse, err := ec.Delete(ctx, key)
	if err != nil {
		log.Fatalf("etcd delete kv err: %v", err)
	}
	log.Printf("etcd delete k: %s, number: %d", key, deleteResponse.Deleted)

	// Output:
	// 2023/12/01 22:00:21 init etcd client success
	// 2023/12/01 22:00:21 etcd put k: hello, v: world success
	// 2023/12/01 22:00:21 etcd get k: hello, v: world, version: 1
	// 2023/12/01 22:00:21 etcd put k: hello, v: world success
	// 2023/12/01 22:00:21 etcd get k: hello, v: world2, version: 2
	// 2023/12/01 22:00:21 etcd delete k: hello, number: 1
	// 2023/12/01 22:00:21 close etcd client success
}
