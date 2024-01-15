package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestSlidingWindowLimiter(t *testing.T) {
	l := NewSlidingWindowLimiter(time.Second, 200*time.Millisecond, 5)
	for {
		result := l.Allow()
		fmt.Println(time.Now().Format(DatetimeFormat), result)
		time.Sleep(time.Millisecond * 100)
	}
}
