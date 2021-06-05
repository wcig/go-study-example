package ch8_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// context
// 上下文类型，可以在通信中携带截止时间、取消信号和其他请求范围值。

// 变量
func TestErr(t *testing.T) {
	_ = context.Canceled         // context被取消返回该错误
	_ = context.DeadlineExceeded // context截止时间到达返回该错误
}

// 创建context:
// 1.创建一个非nil的空context（没有携带任何信息）
// func Background() Context
// 2.创建一个非nil的空context（当不清楚使用哪个context可使用这个）
// func TODO() Context
// 3.创建一个parent的拷贝的带新Done的context，在cancel函数被调用或父context通道关闭时，返回的context Done通道关闭。（取消此上下文会释放其关联的资源，所以一旦此上下文的操作完成，代码应立即调用cancel）
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// 4.创建一个parent的拷贝context，其截止时间不晚于d
// func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
// 5.等价于WithDeafline(parent, time.Now().Add(timeout))
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// 6.创建一个parent的拷贝context，并携带key、val值
// func WithValue(parent Context, key, val interface{}) Context

// CancelFunc 告诉操作放弃其工作。 CancelFunc 不会等待工作停止。一个 CancelFunc 可以被多个 goroutine 同时调用。在第一次调用之后，对 CancelFunc 的后续调用将不执行任何操作。
// type CancelFunc func()

func TestNewEmptyContext(t *testing.T) {
	c1 := context.Background()
	c2 := context.TODO()
	fmt.Println(c1)
	fmt.Println(c2)
	// output:
	// context.Background
	// context.TODO
}

// 使用可取消上下文来防止 goroutine 泄漏。到示例函数结束时，由 gen 启动的 goroutine 将返回而不会泄漏。
func TestWithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	//  output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

// 创建一个带截止时间的上下文，当截止时间到达上下文错误为超时错误
func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	// output:
	// context deadline exceeded
}

// 创建一个带超时时间的上下文，带超时时间到达上下文错误为超时错误
func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	// output:
	// context deadline exceeded
}

// 创建一个带key、val的上下文，以及如何检索它
func TestWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
	// output:
	// found value: Go
	// key not found: color
}
