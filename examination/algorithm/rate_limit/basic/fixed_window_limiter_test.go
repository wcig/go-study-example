package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestFixedWindowLimiter(t *testing.T) {
	l := NewFixedWindowLimiter(time.Second, 2)
	for {
		result := l.Allow()
		fmt.Println(time.Now().Format(DatetimeFormat), result)
		time.Sleep(time.Millisecond * 300)
	}
}
