package basic

import (
	"fmt"
	"sync"
	"time"
)

// TODO: some problem
type LeakyBucketLimiter struct {
	rate      int
	peakLevel int

	currentLevel int
	lastTime     time.Time
	mu           sync.Mutex
}

func NewLeakyBucketLimiter(rate int, peakLevel int) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		rate:      rate,
		peakLevel: peakLevel,
		lastTime:  time.Now(),
	}
}

func (l *LeakyBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	defer func() {
		fmt.Println(">>", time.Now().Format("2006-01-02T15:04:05.000"), l.lastTime.Format("2006-01-02T15:04:05.000"),
			l.currentLevel)
	}()

	now := time.Now()
	interval := now.Sub(l.lastTime)
	subtractLevel := int(float64(interval) * float64(l.rate) / float64(time.Second))
	if subtractLevel > 0 {
		l.lastTime = now
		l.currentLevel -= subtractLevel
		if l.currentLevel < 0 {
			l.currentLevel = 0
		}
	}
	if l.currentLevel >= l.peakLevel {
		return false
	}
	l.currentLevel++
	return true
}
