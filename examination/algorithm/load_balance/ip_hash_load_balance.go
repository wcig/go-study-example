package load_balance

import (
	"hash/fnv"
	"net"
	"sync"
)

// IP哈希（IP Hash）算法：将请求根据客户端IP地址的哈希值分配给特定的服务器。对于相同的IP地址，始终将请求发送到同一台服务器上，适用于需要保持会话一致性的应用。
type IpHashLoadBalance struct {
	servers []string
	mu      sync.Mutex
}

func NewIpHashLoadBalance(servers []string) *IpHashLoadBalance {
	return &IpHashLoadBalance{
		servers: servers,
		mu:      sync.Mutex{},
	}
}

func (l *IpHashLoadBalance) NextServer(clientIp string) string {
	h := fnv.New32a()
	_, err := h.Write(net.ParseIP(clientIp).To4())
	if err != nil {
		panic(err)
	}
	ipHash := h.Sum32()
	index := int(ipHash) % len(l.servers)
	return l.servers[index]
}
