package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenRandNum(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(GenRandNum(1, 5))
	}
	// output:
	// 2
	// 3
	// 4
	// 3
	// 3
}

// --------------------------------------------------------------- //

// 生成随机数, 取值范围:[min, max]
func GenRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
