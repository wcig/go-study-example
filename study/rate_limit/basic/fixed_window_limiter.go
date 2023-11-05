package basic

import (
	"fmt"
	"sync"
	"time"
)

// 固定窗口限流算法
type FixedWindowLimiter struct {
	window time.Duration // 固定窗口
	limit  int           // 限流数

	count    int        // 计数器
	lastTime time.Time  // 上次请求时间
	mu       sync.Mutex // 锁
}

func NewFixedWindowLimiter(window time.Duration, limit int) *FixedWindowLimiter {
	if window <= 0 || limit <= 0 {
		panic(fmt.Errorf("invalid param"))
	}
	return &FixedWindowLimiter{
		window:   window,
		limit:    limit,
		lastTime: time.Now(),
	}
}

func (l *FixedWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	fmt.Println(">>", now.Format(DatetimeFormat), l.lastTime.Format(DatetimeFormat), now.Sub(l.lastTime) > l.window, l.count)

	if now.Sub(l.lastTime) > l.window {
		l.count = 0
		l.lastTime = now
	}
	if l.count >= l.limit {
		return false
	}
	l.count++
	return true
}

const DatetimeFormat = "2006-01-02T15:04:05.000"
