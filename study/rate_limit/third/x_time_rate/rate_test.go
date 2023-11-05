package x_time_rate

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestFirst(t *testing.T) {
	r := rate.Every(time.Millisecond * 300)
	l := rate.NewLimiter(r, 3)
	for i := 0; i < 20; i++ {
		allow := l.Allow()
		fmt.Println(i+1, time.Now().Format("2006-01-02T15:04:05.000"), allow)
		time.Sleep(time.Millisecond * 100)
	}
	// Output:
	// 1 2023-11-05T16:55:15.609 true
	// 2 2023-11-05T16:55:15.710 true
	// 3 2023-11-05T16:55:15.811 true
	// 4 2023-11-05T16:55:15.913 true
	// 5 2023-11-05T16:55:16.014 false
	// 6 2023-11-05T16:55:16.115 false
	// 7 2023-11-05T16:55:16.216 true
	// 8 2023-11-05T16:55:16.317 false
	// 9 2023-11-05T16:55:16.418 false
	// 10 2023-11-05T16:55:16.519 true
	// 11 2023-11-05T16:55:16.620 false
	// 12 2023-11-05T16:55:16.722 false
	// 13 2023-11-05T16:55:16.822 true
	// 14 2023-11-05T16:55:16.923 false
	// 15 2023-11-05T16:55:17.024 false
	// 16 2023-11-05T16:55:17.125 true
	// 17 2023-11-05T16:55:17.226 false
	// 18 2023-11-05T16:55:17.327 false
	// 19 2023-11-05T16:55:17.428 true
	// 20 2023-11-05T16:55:17.529 false
}
