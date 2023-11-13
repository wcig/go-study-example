package ch8_context

import (
	"context"
	"log"
	"testing"
	"time"
)

// ctx的key/val值在整个链条都可以获取
func TestValContext(t *testing.T) {
	ctx1 := context.WithValue(context.Background(), "key", "val")
	log.Println(">> ctx1 key:", ctx1.Value("key"))

	ctx2, cancel := context.WithCancel(ctx1)
	defer cancel()
	log.Println(">> ctx2 key:", ctx2.Value("key"))

	ctx3, timeoutCancel := context.WithTimeout(ctx2, time.Second)
	defer timeoutCancel()
	log.Println(">> ctx3 key:", ctx3.Value("key"))

	log.Println(ctx1)
	log.Println(ctx2)
	log.Println(ctx3)

	// Output:
	// 2023/11/13 20:27:22 >> ctx1 key: val
	// 2023/11/13 20:27:22 >> ctx2 key: val
	// 2023/11/13 20:27:22 >> ctx3 key: val
	// 2023/11/13 20:27:22 context.Background.WithValue(type string, val val)
	// 2023/11/13 20:27:22 context.Background.WithValue(type string, val val).WithCancel
	// 2023/11/13 20:27:22 context.Background.WithValue(type string, val val).WithCancel.WithDeadline(2023-11-13 20:27:23.195013 +0800 CST m=+1.001581042 [999.925083ms])
}
