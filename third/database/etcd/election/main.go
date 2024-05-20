package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr      = flag.String("addr", "http://localhost:2379", "etcd addresses")
	electName = flag.String("name", "my-test-elect", "elect name")
	nodeID    = flag.Int("id", 0, "node ID")
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

	// 创建选举
	el := concurrency.NewElection(ss, *electName)

	// 命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		action := scanner.Text()
		switch action {
		case "elect":
			go elect(el, *electName)
		case "proclaim":
			proclaim(el, *electName)
		case "resign":
			resign(el, *electName)
		case "watch":
			go watch(el, *electName)
		case "query":
			query(el, *electName)
		case "rev":
			rev(el, *electName)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}

var count int

// 选主
func elect(el *concurrency.Election, name string) {
	log.Println("campaigning for ID:", *nodeID)
	if err := el.Campaign(context.Background(), name); err != nil {
		log.Printf("campaigning for ID: %d err: %v", *nodeID, err)
		return
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

// 为主设置新值
func proclaim(el *concurrency.Election, name string) {
	log.Println("proclaiming for ID:", *nodeID)
	if err := el.Proclaim(context.Background(),
		fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Printf("proclaiming for ID: %d err: %v", *nodeID, err)
		return
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}

// 重新选主 (其他节点可能被选为主)
func resign(el *concurrency.Election, name string) {
	log.Println("resigning for ID:", *nodeID)
	if err := el.Resign(context.Background()); err != nil {
		log.Printf("resigning for ID: %d err: %v", *nodeID, err)
		return
	}
	log.Println("resigned for ID:", *nodeID)
	count++
}

// 监控
func watch(el *concurrency.Election, name string) {
	ch := el.Observe(context.Background())
	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader change to:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}

// 查询主
func query(el *concurrency.Election, name string) {
	resp, err := el.Leader(context.Background())
	if err != nil {
		log.Printf("query leader err: %v", err)
		return
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

// 查询主revision信息
func rev(el *concurrency.Election, name string) {
	revision := el.Rev()
	log.Println("current rev:", revision)
}
