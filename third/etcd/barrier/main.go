package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr        = flag.String("addr", "http://localhost:2379", "etcd addresses")
	barrierName = flag.String("name", "my-test-barrier", "barrier name")
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

	// 创建栅栏
	barrier := recipe.NewBarrier(ec, *barrierName)

	// 命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		action := scanner.Text()
		switch action {
		// 持有barrier
		case "hold":
			fmt.Println("before hold")
			if err = barrier.Hold(); err != nil {
				fmt.Printf("hold err: %v", err)
				continue
			}
			fmt.Println("after hold")
		// 等待barrier
		case "wait":
			fmt.Println("before wait")
			if err = barrier.Wait(); err != nil {
				fmt.Printf("wait err: %v", err)
				continue
			}
			fmt.Println("after wait")
		// 释放barrier
		case "release":
			fmt.Println("before release")
			if err = barrier.Release(); err != nil {
				fmt.Printf("release err: %v", err)
				continue
			}
			fmt.Println("after release")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
