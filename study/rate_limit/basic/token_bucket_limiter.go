package basic

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	rate   int
	bucket int

	tokens   int
	lastTime time.Time
	mu       sync.Mutex
}

func NewTokenBucketLimiter(rate int, bucket int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		rate:     rate,
		bucket:   bucket,
		lastTime: time.Now(),
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	interval := now.Sub(l.lastTime)
	addTokens := int(float64(interval) * float64(l.rate) / float64(time.Second))
	if addTokens > 0 {
		l.lastTime = now
		l.tokens += addTokens
		if l.tokens > l.bucket {
			l.tokens = l.bucket
		}
	}
	if l.tokens <= 0 {
		return false
	}
	l.tokens--
	return true
}
