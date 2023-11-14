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
			if err := barrier.Await(context.Background()); err != nil {
				log.Fatalf(">> %d err: %v", id, err)
			}
			log.Printf(">> %d end", id)
		}()
	}
	time.Sleep(3 * time.Second)
	log.Println(">> over")
	// Output:
	// 2023/11/14 17:46:15 >> 2 start
	// 2023/11/14 17:46:15 >> 0 start
	// 2023/11/14 17:46:15 >> 1 start
	// 2023/11/14 17:46:15 >> 1 end
	// 2023/11/14 17:46:15 >> 2 end
	// 2023/11/14 17:46:15 >> 0 end
	// 2023/11/14 17:46:18 >> over
}
