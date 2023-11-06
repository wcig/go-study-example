package load_balance

import "sync"

// 最少连接（Least Connections）算法：将请求分配给当前连接数最少的服务器，以实现负载均衡。适用于处理长连接或请求处理时间不均匀的情况。
type LeastConnectionLoadBalance struct {
	servers []*Server
	mu      sync.Mutex
}

type Server struct {
	address     string
	connections int64
}

func NewLeastConnectionLoadBalance(servers []string) *LeastConnectionLoadBalance {
	list := make([]*Server, len(servers))
	for i := range servers {
		list[i] = &Server{address: servers[i], connections: 0}
	}
	return &LeastConnectionLoadBalance{
		servers: list,
		mu:      sync.Mutex{},
	}
}

func (l *LeastConnectionLoadBalance) NextServer() *Server {
	l.mu.Lock()
	defer l.mu.Unlock()
	index := 0
	min := l.servers[0].connections
	for i := 1; i < len(l.servers); i++ {
		if l.servers[i].connections < min {
			index = i
			min = l.servers[i].connections
		}
	}
	return l.servers[index]
}

// func HandleRequest() {
// 	server := NextServer()
//
// 	// 增加服务器的当前连接数
// 	server.connections++
//
// 	// 向选定的服务器发送请求
// 	sendRequestToServer(server.Address)
//
// 	// 处理完成后减少服务器的当前连接数
// 	server.connections--
// }
