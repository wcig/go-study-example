package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr      = flag.String("addr", "http://localhost:2379", "etcd addresses")
	queueName = flag.String("name", "my-test-queue", "queue name")
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

	// 创建队列
	queue := recipe.NewQueue(ec, *queueName)

	// 命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		action := items[0]
		switch action {
		case "push":
			if len(items) != 2 {
				fmt.Println("queue: must set value to push")
				continue
			}
			// 入队
			if err = queue.Enqueue(items[1]); err != nil {
				fmt.Printf("queue: push value err: %v\n", err)
			}
			fmt.Printf("queue: push value success: %s\n", items[1])
		case "pop":
			// 出队 (队列为空时阻塞)
			val, err2 := queue.Dequeue()
			if err2 != nil {
				fmt.Printf("queue: pop value err: %v\n", err2)
				continue
			}
			fmt.Printf("queue: pop value success: %s\n", val)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
