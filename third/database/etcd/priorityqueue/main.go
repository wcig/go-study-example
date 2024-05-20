package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

var (
	addr              = flag.String("addr", "http://localhost:2379", "etcd addresses")
	priorityQueueName = flag.String("name", "my-test-priority-queue", "priority queue name")
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

	// 创建优先队列
	queue := recipe.NewPriorityQueue(ec, *priorityQueueName)

	// 命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		action := items[0]
		switch action {
		case "push":
			if len(items) != 3 {
				fmt.Println("priority-queue: must set value and priority to push")
				continue
			}
			pr, err2 := strconv.Atoi(items[2])
			if err2 != nil {
				fmt.Println("priority-queue: must set priority to push")
				continue
			}
			// 入队
			if err = queue.Enqueue(items[1], uint16(pr)); err != nil {
				fmt.Printf("priority-queue: push value err: %v\n", err)
				continue
			}
			fmt.Printf("priority-queue: push value success: %s\n", items[1])
		case "pop":
			// 出队 (队列为空时阻塞)
			val, err2 := queue.Dequeue()
			if err2 != nil {
				fmt.Printf("priority-queue: pop value err: %v\n", err2)
				continue
			}
			fmt.Printf("priority-queue: pop value success: %s\n", val)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
