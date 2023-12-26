package question

import (
	"fmt"
	"math"
	"testing"
)

// 问题:
// 零钱兑换: 给定 n 种硬币, 第 i 种硬币的面值为 coins[i-1], 目标金额为 amt, 每种硬币可以重复选取.
// 问能凑出目标金额的最少硬币数量, 如果无法凑出目标金额则返回 -1.

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 4

	// 动态规划
	res := coinChangeDP(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)

	// 空间优化后的动态规划
	res = coinChangeDPComp(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)
}

func coinChangeDP(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amt+1)
	}
	for j := 1; j <= amt; j++ {
		dp[0][j] = max
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amt; j++ {
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coins[i-1]]+1)))
			}
		}
	}
	if dp[n][amt] != max {
		return dp[n][amt]
	}
	return -1
}

func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	dp := make([]int, amt+1)
	for i := 1; i <= amt; i++ {
		dp[i] = max
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amt; j++ {
			if coins[i-1] > j {
				dp[j] = dp[j]
			} else {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1)))
			}
		}
	}
	if dp[amt] != max {
		return dp[amt]
	}
	return -1
}
