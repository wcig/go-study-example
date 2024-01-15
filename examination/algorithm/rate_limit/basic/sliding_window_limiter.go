package basic

import (
	"fmt"
	"sync"
	"time"
)

// 滑动窗口限流算法
type SlidingWindowLimiter struct {
	window       time.Duration // 大窗口
	smallWindow  time.Duration // 小窗口
	smallWindows int           // 小窗口数量
	buckets      map[int64]int // 各小窗口统计数 (key: 小窗口起始时间, val: 小窗口统计数)
	limit        int           // 大窗口限流数

	lastTime time.Time  // 上次请求时间
	mu       sync.Mutex // 锁
}

func NewSlidingWindowLimiter(window, smallWindow time.Duration, limit int) *SlidingWindowLimiter {
	if window <= 0 || smallWindow <= 0 || smallWindow > window || limit <= 0 {
		panic(fmt.Errorf("invalid param"))
	}
	return &SlidingWindowLimiter{
		window:       window,
		smallWindow:  smallWindow,
		smallWindows: int(window / smallWindow),
		buckets:      map[int64]int{},
		limit:        limit,
		lastTime:     time.Now(),
	}
}

func (l *SlidingWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 当前时间小窗口开始时间
	curSmallWindow := time.Now().UnixNano() / int64(l.smallWindow) * int64(l.smallWindow)
	// 整个滑动窗口中起始小窗口开始时间
	startSmallWindow := curSmallWindow - int64(l.smallWindow)*(int64(l.smallWindows)-1)
	// 统计所有小窗口总数
	total := 0
	for smallWindow, count := range l.buckets {
		if smallWindow < startSmallWindow {
			// 过期小窗口删除
			delete(l.buckets, smallWindow)
		} else {
			total += count
		}
	}
	if total >= l.limit {
		return false
	}
	l.buckets[curSmallWindow]++
	return true
}
