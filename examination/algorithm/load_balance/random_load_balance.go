package load_balance

import (
	"math/rand"
	"sync"
	"time"
)

// 随机（Random）算法：随机选择一个服务器来处理请求，以实现负载均衡。适用于服务器性能相近且无状态的情况。
type RandomLoadBalance struct {
	servers []string
	r       *rand.Rand
	mu      sync.Mutex
}

func NewRandomLoadBalance(servers []string) *RandomLoadBalance {
	return &RandomLoadBalance{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		mu:      sync.Mutex{},
	}
}

func (l *RandomLoadBalance) NextServer() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	index := l.r.Intn(len(l.servers))
	return l.servers[index]
}
