package load_balance

import (
	"fmt"
	"testing"
)

func TestIpHashLoadBalance(t *testing.T) {
	l := NewIpHashLoadBalance(servers)
	clientIpPrefix := "192.168.0."
	for i := 0; i < 10; i++ {
		clientIp := fmt.Sprintf("%s%d", clientIpPrefix, i)
		fmt.Println(i, clientIp, l.NextServer(clientIp))
	}

	// Output:
	// 0 192.168.0.0 server1.com
	// 1 192.168.0.1 server2.com
	// 2 192.168.0.2 server3.com
	// 3 192.168.0.3 server1.com
	// 4 192.168.0.4 server2.com
	// 5 192.168.0.5 server3.com
	// 6 192.168.0.6 server1.com
	// 7 192.168.0.7 server2.com
	// 8 192.168.0.8 server3.com
	// 9 192.168.0.9 server1.com
}
