package load_balance

import "sync"

// 加权轮询（Weighted Round Robin）算法：为每个服务器分配一个权重值，根据权重比例决定发送请求的概率。权重越高的服务器获取到的请求数量越多。
type WeightRoundRobinLoadBalance struct {
	servers []*WeightServer
	mu      sync.Mutex
}

type WeightServer struct {
	address         string
	weight          int
	currentWeight   int
	effectiveWeight int
}

func NewWeightRoundRobinLoadBalance(addressList []string, weightList []int) *WeightRoundRobinLoadBalance {
	servers := make([]*WeightServer, len(addressList))
	for i := range addressList {
		server := &WeightServer{
			address:         addressList[i],
			weight:          weightList[i],
			currentWeight:   weightList[i],
			effectiveWeight: weightList[i],
		}
		servers[i] = server
	}
	return &WeightRoundRobinLoadBalance{
		servers: servers,
		mu:      sync.Mutex{},
	}
}

func (l *WeightRoundRobinLoadBalance) NextServer() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	totalWeight := 0
	var maxWeightServer *WeightServer
	for i := range l.servers {
		totalWeight += l.servers[i].effectiveWeight
		l.servers[i].currentWeight = l.servers[i].currentWeight + l.servers[i].effectiveWeight
		if maxWeightServer == nil || maxWeightServer.currentWeight < l.servers[i].currentWeight {
			maxWeightServer = l.servers[i]
		}
	}
	maxWeightServer.currentWeight -= totalWeight
	return maxWeightServer.address
}
