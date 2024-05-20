package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr              = flag.String("addr", "http://localhost:2379", "etcd addresses")
	doubleBarrierName = flag.String("name", "my-test-double-barrier", "double barrier name")
	count             = flag.Int("count", 2, "barrier count")
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

	// 创建栅栏
	barrier := recipe.NewDoubleBarrier(ss, *doubleBarrierName, *count)

	// 命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		action := scanner.Text()
		switch action {
		case "enter":
			fmt.Println("before enter")
			if err = barrier.Enter(); err != nil {
				fmt.Printf("enter err: %v", err)
				continue
			}
			fmt.Println("after enter")
		case "leave":
			fmt.Println("before leave")
			if err = barrier.Leave(); err != nil {
				fmt.Printf("leave err: %v", err)
				continue
			}
			fmt.Println("after leave")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
