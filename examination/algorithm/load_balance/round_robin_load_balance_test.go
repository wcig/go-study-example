package load_balance

import (
	"fmt"
	"testing"
)

var servers = []string{"server1.com", "server2.com", "server3.com"}

func TestRoundRobinLoadBalance(t *testing.T) {
	l := NewRoundRobinLoadBalance(servers)
	for i := 0; i < 10; i++ {
		fmt.Println(i, l.NextServer())
	}

	// Output:
	// 0 server1.com
	// 1 server2.com
	// 2 server3.com
	// 3 server1.com
	// 4 server2.com
	// 5 server3.com
	// 6 server1.com
	// 7 server2.com
	// 8 server3.com
	// 9 server1.com
}
