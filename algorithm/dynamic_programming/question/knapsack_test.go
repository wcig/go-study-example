package question

import (
	"fmt"
	"math"
	"testing"
)

// 问题:
// 0-1背包: 给定n个物品, 第之个物品的重量为 wgt[i-1], 价值为 wal[i一1], 和一个容量为cap 的背包.
// 每个物品只能选择一次, 问在限定背包容量下能放入物品的最大价值.
func TestKnapsack(t *testing.T) {
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	c := 50

	// 动态规划
	res := knapsackDP(wgt, val, c)
	fmt.Printf("不超过背包容量的最大物品价值为 %d\n", res)

	// 空间优化后的动态规划
	res = knapsackDPComp(wgt, val, c)
	fmt.Printf("不超过背包容量的最大物品价值为 %d\n", res)
}

func knapsackDP(wgt []int, val []int, cap int) int {
	n := len(wgt)
	// 初始化dp表
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}
	// dp状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= cap; j++ {
			if wgt[i-1] > j {
				// 重量超过剩余背包容量, 不选择物品i
				continue
			} else {
				// 不选和选择物品i两种方案取最大值
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]
}

func knapsackDPComp(wgt []int, val []int, cap int) int {
	n := len(wgt)
	dp := make([]int, cap+1)
	for i := 1; i <= n; i++ {
		// 倒序遍历
		for c := cap; c >= 1; c-- {
			if wgt[i-1] <= c {
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}
