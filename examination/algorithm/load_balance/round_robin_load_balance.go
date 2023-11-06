package load_balance

import "sync"

// 轮询（Round Robin）算法：将请求按顺序分配给后端服务器，每个请求依次轮流分配到不同的服务器上。适用于服务器性能相近且无状态的情况。
type RoundRobinLoadBalance struct {
	servers []string
	index   int
	mu      sync.Mutex
}

func NewRoundRobinLoadBalance(servers []string) *RoundRobinLoadBalance {
	return &RoundRobinLoadBalance{
		servers: servers,
		index:   0,
		mu:      sync.Mutex{},
	}
}

func (l *RoundRobinLoadBalance) NextServer() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	server := l.servers[l.index]
	l.index = (l.index + 1) % len(l.servers)
	return server
}
