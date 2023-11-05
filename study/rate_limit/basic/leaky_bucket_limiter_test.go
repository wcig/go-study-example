package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestLeakyBucketLimiter(t *testing.T) {
	l := NewLeakyBucketLimiter(5, 10)
	for i := 0; i < 20; i++ {
		allow := l.Allow()
		fmt.Println(i+1, time.Now().Format("2006-01-02T15:04:05.000"), allow)
		time.Sleep(time.Millisecond * 100)
	}
}
