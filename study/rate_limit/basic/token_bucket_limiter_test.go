package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucketLimiter(t *testing.T) {
	l := NewTokenBucketLimiter(5, 10)
	for i := 0; i < 20; i++ {
		allow := l.Allow()
		fmt.Println(i+1, time.Now().Format("2006-01-02T15:04:05.000"), allow)
		time.Sleep(time.Millisecond * 100)
	}
	// Output:
	// 1 2023-11-05T17:17:10.550 false
	// 2 2023-11-05T17:17:10.651 false
	// 3 2023-11-05T17:17:10.752 true
	// 4 2023-11-05T17:17:10.852 false
	// 5 2023-11-05T17:17:10.953 true
	// 6 2023-11-05T17:17:11.055 false
	// 7 2023-11-05T17:17:11.156 true
	// 8 2023-11-05T17:17:11.256 false
	// 9 2023-11-05T17:17:11.357 true
	// 10 2023-11-05T17:17:11.459 false
	// 11 2023-11-05T17:17:11.560 true
	// 12 2023-11-05T17:17:11.660 false
	// 13 2023-11-05T17:17:11.761 true
	// 14 2023-11-05T17:17:11.863 false
	// 15 2023-11-05T17:17:11.964 true
	// 16 2023-11-05T17:17:12.064 false
	// 17 2023-11-05T17:17:12.165 true
	// 18 2023-11-05T17:17:12.266 false
	// 19 2023-11-05T17:17:12.367 true
	// 20 2023-11-05T17:17:12.468 false
}

func TestTokenBucketLimiterWithSleep(t *testing.T) {
	l := NewTokenBucketLimiter(5, 10)
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		allow := l.Allow()
		fmt.Println(i+1, time.Now().Format("2006-01-02T15:04:05.000"), allow)
		time.Sleep(time.Millisecond * 100)
	}
	// Output:
	// 1 2023-11-05T17:44:27.479 true
	// 2 2023-11-05T17:44:27.581 true
	// 3 2023-11-05T17:44:27.682 true
	// 4 2023-11-05T17:44:27.783 true
	// 5 2023-11-05T17:44:27.884 true
	// 6 2023-11-05T17:44:27.985 true
	// 7 2023-11-05T17:44:28.086 true
	// 8 2023-11-05T17:44:28.187 true
	// 9 2023-11-05T17:44:28.288 true
	// 10 2023-11-05T17:44:28.389 false
	// 11 2023-11-05T17:44:28.491 true
	// 12 2023-11-05T17:44:28.592 false
	// 13 2023-11-05T17:44:28.692 true
	// 14 2023-11-05T17:44:28.793 false
	// 15 2023-11-05T17:44:28.894 true
	// 16 2023-11-05T17:44:28.996 false
	// 17 2023-11-05T17:44:29.097 true
	// 18 2023-11-05T17:44:29.198 false
	// 19 2023-11-05T17:44:29.299 true
	// 20 2023-11-05T17:44:29.400 false
}
