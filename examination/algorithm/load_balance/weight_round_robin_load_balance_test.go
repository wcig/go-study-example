package load_balance

import (
	"fmt"
	"testing"
)

func TestWeightRoundRobinLoadBalance(t *testing.T) {
	l := NewWeightRoundRobinLoadBalance(servers, []int{4, 2, 1})
	m := make(map[string]int)
	for i := 0; i < 14; i++ {
		server := l.NextServer()
		fmt.Println(i, server)
		if _, ok := m[server]; ok {
			m[server]++
		} else {
			m[server] = 1
		}
	}
	fmt.Println(m)

	// Output:
	// 0 server1.com
	// 1 server2.com
	// 2 server1.com
	// 3 server1.com
	// 4 server3.com
	// 5 server1.com
	// 6 server2.com
	// 7 server1.com
	// 8 server2.com
	// 9 server1.com
	// 10 server1.com
	// 11 server3.com
	// 12 server1.com
	// 13 server2.com
	// map[server1.com:8 server2.com:4 server3.com:2]
}
