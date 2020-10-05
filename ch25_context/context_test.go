package ch25_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 创建空Context
func TestEmptyContext(t *testing.T) {
	c1 := context.Background()
	fmt.Println(c1)

	c2 := context.TODO()
	fmt.Println(c2)
}

// output:
// context.Background
// context.TODO

// CancelContext
func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go handleWithCancel(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("main ctx over")
	cancel()
}

func handleWithCancel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	default:
		time.Sleep(2 * time.Second)
		fmt.Println("process request over")
	}
}

// output:
// process request over
// main ctx over

// DeadlineContext
func TestDeadlineContext(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go handleWithDeadline(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	time.Sleep(10 * time.Second)
}

func handleWithDeadline(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	default:
		time.Sleep(2 * time.Second)
		fmt.Println("process request over")
	}
}

// output:
// process request over
// main context deadline exceeded

// TimeoutContext
func TestTimeoutContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handleWithTimeout(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handleWithTimeout(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

// output:
// process request with 500ms
// main context deadline exceeded

// ValueContext
func TestValueContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, "name", "valueContext")
	defer cancel()

	handleWithValue(valCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("main ctx over")
}

func handleWithValue(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	default:
		time.Sleep(2 * time.Second)
		fmt.Println("process request over:", ctx.Value("name"))
	}
}

// output:
// process request over: valueContext
// main ctx over
