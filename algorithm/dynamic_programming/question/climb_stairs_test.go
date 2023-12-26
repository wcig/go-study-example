package question

import (
	"fmt"
	"testing"
)

// 问题: 爬楼梯: 给定一个共有 n 阶的楼梯, 你每步可以上 1 或 2 阶, 请问有多少种方案可以爬到楼顶?
func TestClimbStairs(t *testing.T) {
	n := 9
	res := climbStairs(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)

	res = climbStairsComp(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

// 动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 初始化dp表
	dp := make([]int, n+1)
	// dp初始值
	dp[1] = 1
	dp[2] = 2
	// dp状态转移
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 动态规划 (空间优化版)
func climbStairsComp(n int) int {
	if n <= 2 {
		return n
	}
	// dp初始值
	d1 := 1
	d2 := 2
	// dp状态转移
	for i := 3; i <= n; i++ {
		tmp := d1 + d2
		d1 = d2
		d2 = tmp
	}
	return d2
}
