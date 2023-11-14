package cyclic_barrier

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/marusama/cyclicbarrier"
)

func TestCyclicBarrierDemo(t *testing.T) {
	const num = 3
	barrier := cyclicbarrier.New(num)
	for i := 0; i < num; i++ {
		id := i
		go func() {
			log.Printf(">> %d start", id)
			time.Sleep(time.Second)
			if err := barrier.Await(context.Background()); err != nil {
				log.Fatalf(">> %d err: %v", id, err)
			}
			log.Printf(">> %d end", id)
		}()
	}
	time.Sleep(3 * time.Second)
	log.Println(">> over")
	// Output:
	// 2023/11/14 17:48:19 >> 2 start
	// 2023/11/14 17:48:19 >> 0 start
	// 2023/11/14 17:48:19 >> 1 start
	// 2023/11/14 17:48:20 >> 2 end
	// 2023/11/14 17:48:20 >> 1 end
	// 2023/11/14 17:48:20 >> 0 end
	// 2023/11/14 17:48:22 >> over
}

func TestCyclicBarrierDemoWithAcion(t *testing.T) {
	const num = 3
	cnt := 0
	barrier := cyclicbarrier.NewWithAction(num, func() error {
		cnt++
		return nil
	})
	for i := 0; i < num; i++ {
		id := i
		go func() {
			log.Printf(">> %d start", id)
			time.Sleep(time.Second)
			if err := barrier.Await(context.Background()); err != nil {
				log.Fatalf(">> %d err: %v", id, err)
			}
			log.Printf(">> %d end", id)
		}()
	}
	time.Sleep(3 * time.Second)
	log.Println(">> over", cnt)
	// Output:
	// 2023/11/14 17:57:26 >> 1 start
	// 2023/11/14 17:57:26 >> 2 start
	// 2023/11/14 17:57:26 >> 0 start
	// 2023/11/14 17:57:27 >> 2 end
	// 2023/11/14 17:57:27 >> 1 end
	// 2023/11/14 17:57:27 >> 0 end
	// 2023/11/14 17:57:29 >> over 1
}
