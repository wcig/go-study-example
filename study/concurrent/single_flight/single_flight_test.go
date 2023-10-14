package single_flight

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

// golang.org/x/sync: singleflight 工具包
// 参考:
// 1.https://www.liwenzhou.com/posts/Go/singleflight/
// 2.https://lailin.xyz/post/go-training-week5-singleflight.html

// func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)
// 并发调用时, 只有第一次真正调用, 其他阻塞等待第一次调用结果
func TestDo(t *testing.T) {
	sg := new(singleflight.Group)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	start := time.Now()
	const loadDataKey = "loadDataDo"

	go func() {
		defer wg.Done()
		v, err, shared := sg.Do(loadDataKey, func() (interface{}, error) {
			data := loadData()
			return data, nil
		})
		if err != nil {
			log.Fatalf(">> query-1 load data err: %v", err)
		}
		data, ok := v.(string)
		if !ok {
			log.Fatalf(">> query-1 transfer data err: %v", err)
		}
		log.Printf(">> query-1 val: %s, shared: %v", data, shared)
	}()

	time.Sleep(50 * time.Millisecond)
	// time.Sleep(250 * time.Millisecond)

	go func() {
		defer wg.Done()
		v, err, shared := sg.Do(loadDataKey, func() (interface{}, error) {
			data := loadData()
			return data, nil
		})
		if err != nil {
			log.Fatalf(">> query-2 load data err: %v", err)
		}
		data, ok := v.(string)
		if !ok {
			log.Fatalf(">> query-2 transfer data err: %v", err)
		}
		log.Printf(">> query-2 val: %s, shared: %v", data, shared)
	}()

	wg.Wait()
	log.Printf(">> all done, time cost: %v", time.Since(start))

	// Output:
	// 1.time.Sleep(50 * time.Millisecond)
	// 2023/10/14 13:33:03 >> load data...
	// 2023/10/14 13:33:03 >> query-1 val: data, shared: true
	// 2023/10/14 13:33:03 >> query-2 val: data, shared: true
	// 2023/10/14 13:33:03 >> all done, time cost: 101.494416ms

	// 2.time.Sleep(250 * time.Millisecond)
	// 2023/10/14 13:31:21 >> load data...
	// 2023/10/14 13:31:21 >> query-1 val: data, shared: false
	// 2023/10/14 13:31:21 >> load data...
	// 2023/10/14 13:31:21 >> query-2 val: data, shared: false
	// 2023/10/14 13:31:21 >> all done, time cost: 351.944875ms
}

func loadData() string {
	log.Printf(">> load data...")
	time.Sleep(100 * time.Millisecond)
	// time.Sleep(200 * time.Millisecond)
	return "data"
}

func TestDoChan(t *testing.T) {
	sg := new(singleflight.Group)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	start := time.Now()
	const loadDataDoChanKey = "loadDataDoChan"

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		val, err := loadDataSingleFlight(ctx, sg, loadDataDoChanKey)
		log.Printf(">> query-1 val: %s, err: %v", val, err)
	}()

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		val, err := loadDataSingleFlight(ctx, sg, loadDataDoChanKey)
		log.Printf(">> query-2 val: %s, err: %v", val, err)
	}()

	wg.Wait()
	log.Printf(">> all done, time cost: %v", time.Since(start))

	// Output:
	// 1.loadData() 耗时100ms
	// 2023/10/14 16:29:17 >> load data...
	// 2023/10/14 16:29:18 >> query-1 val: data, err: <nil>
	// 2023/10/14 16:29:18 >> query-2 val: data, err: <nil>
	// 2023/10/14 16:29:18 >> all done, time cost: 100.781959ms

	// 2.loadData() 耗时300ms
	// 2023/10/14 16:30:32 >> load data...
	// 2023/10/14 16:30:33 >> query-2 val: , err: context deadline exceeded
	// 2023/10/14 16:30:33 >> query-1 val: , err: context deadline exceeded
	// 2023/10/14 16:30:33 >> all done, time cost: 201.199666ms
}

// func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result
// 通过 chan 返回结果
func loadDataSingleFlight(ctx context.Context, sg *singleflight.Group, key string) (string, error) {
	result := sg.DoChan(key, func() (interface{}, error) {
		data := loadData()
		return data, nil
	})
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case r := <-result:
		return r.Val.(string), r.Err
	}
}

// func (g *Group) Forget(key string)
// 忘记 key, 下一次 key 的 Do 函数调用不会等待之前的调用结果
func TestForget(t *testing.T) {
	sg := new(singleflight.Group)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	start := time.Now()
	const loadDataForgetKey = "loadDataForget"

	go func() {
		defer wg.Done()
		v, err, shared := sg.Do(loadDataForgetKey, func() (interface{}, error) {
			data := loadData()
			return data, nil
		})
		if err != nil {
			log.Fatalf(">> query-1 load data err: %v", err)
		}
		data, ok := v.(string)
		if !ok {
			log.Fatalf(">> query-1 transfer data err: %v", err)
		}
		log.Printf(">> query-1 val: %s, shared: %v", data, shared)
	}()

	time.Sleep(50 * time.Millisecond)
	sg.Forget(loadDataForgetKey)

	go func() {
		defer wg.Done()
		v, err, shared := sg.Do(loadDataForgetKey, func() (interface{}, error) {
			data := loadData()
			return data, nil
		})
		if err != nil {
			log.Fatalf(">> query-2 load data err: %v", err)
		}
		data, ok := v.(string)
		if !ok {
			log.Fatalf(">> query-2 transfer data err: %v", err)
		}
		log.Printf(">> query-2 val: %s, shared: %v", data, shared)
	}()

	wg.Wait()
	log.Printf(">> all done, time cost: %v", time.Since(start))

	// Output:
	// 2023/10/14 16:33:09 >> load data...
	// 2023/10/14 16:33:09 >> load data...
	// 2023/10/14 16:33:10 >> query-1 val: data, shared: false
	// 2023/10/14 16:33:10 >> query-2 val: data, shared: false
	// 2023/10/14 16:33:10 >> all done, time cost: 151.342042ms
}
