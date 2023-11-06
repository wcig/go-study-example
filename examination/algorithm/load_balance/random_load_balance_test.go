package load_balance

import (
	"fmt"
	"testing"
)

func TestRandomLoadBalance(t *testing.T) {
	l := NewRandomLoadBalance(servers)
	for i := 0; i < 10; i++ {
		fmt.Println(i, l.NextServer())
	}

	// Output:
	// 0 server3.com
	// 1 server2.com
	// 2 server2.com
	// 3 server1.com
	// 4 server3.com
	// 5 server2.com
	// 6 server1.com
	// 7 server2.com
	// 8 server2.com
	// 9 server1.com
}
