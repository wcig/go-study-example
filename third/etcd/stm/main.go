package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr = flag.String("addr", "http://localhost:2379", "etcd addresses")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 创建client
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{*addr}})
	if err != nil {
		log.Fatalf("init etcd client err: %v", err)
	}
	defer func() {
		if err = cli.Close(); err != nil {
			log.Fatalf("close etcd client err: %v", err)
		}
	}()

	// 设置5个账户，每个账号都有100元，总共500元
	const totalAccounts = 5
	for i := 0; i < totalAccounts; i++ {
		k := fmt.Sprintf("accounts/%d", i)
		if _, err = cli.Put(context.TODO(), k, "100"); err != nil {
			log.Fatal(err)
		}
	}

	// STM的应用函数，主要的事务逻辑
	exchange := func(stm concurrency.STM) error {
		// 随机得到两个转账账号
		from, to := rand.Intn(totalAccounts), rand.Intn(totalAccounts)
		if from == to {
			// 自己不和自己转账
			return nil
		}
		// 读取账号的值
		fromK, toK := fmt.Sprintf("accounts/%d", from), fmt.Sprintf("accounts/%d", to)
		fromV, toV := stm.Get(fromK), stm.Get(toK)
		fromInt, _ := strconv.Atoi(fromV)
		toInt, _ := strconv.Atoi(toV)
		// 把源账号一半的钱转账给目标账号
		xfer := fromInt / 2
		fromInt, toInt = fromInt-xfer, toInt+xfer
		// 把转账后的值写回
		stm.Put(fromK, fmt.Sprintf("%d", fromInt))
		stm.Put(toK, fmt.Sprintf("%d", toInt))
		return nil
	}

	// 启动10个goroutine进行转账操作
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				if _, err2 := concurrency.NewSTM(cli, exchange); err2 != nil {
					log.Fatal(err2)
				}
			}
		}()
	}
	wg.Wait()

	// 检查账号最后的数目
	sum := 0
	accounts, err := cli.Get(context.TODO(), "accounts/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range accounts.Kvs {
		// 遍历账号的值
		v, _ := strconv.Atoi(string(kv.Value))
		sum += v
		log.Printf("account %s: %d", kv.Key, v)
	}
	// 总数
	log.Println("account sum is", sum)

	// Output:
	// 2023/12/02 12:17:30 account accts/0: 120
	// 2023/12/02 12:17:30 account accts/1: 34
	// 2023/12/02 12:17:30 account accts/2: 245
	// 2023/12/02 12:17:30 account accts/3: 82
	// 2023/12/02 12:17:30 account accts/4: 19
	// 2023/12/02 12:17:30 account sum is 500
}
