package ch105_xtimerate

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

// 限流器: "golang.org/x/time/rate"

func TestSdkTimeLimit(t *testing.T) {
	limiter := rate.NewLimiter(1, 1)

	const num = 10
	for i := 1; i <= num; i++ {
		fmt.Println("add task:", i, time.Now())

		// output, err := baselineDetectWithRateLimit(limiter, i)
		// fmt.Println("output:", output, err)

		n := i
		go func() {
			output, err := runWithTimeLimit(limiter, n)
			fmt.Println("output:", output, err)
		}()
	}

	time.Sleep(time.Minute)
}

func runWithTimeLimit(limiter *rate.Limiter, input int) (output *int, err error) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err = limiter.Wait(ctx)
		output = run(input)
	}()

	wg.Wait()
	fmt.Println("exec task:", input, *output, time.Now())
	return output, err
}

func run(input int) (output *int) {
	n := input + 10
	return &n
}
