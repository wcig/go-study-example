// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_backtracking

package question

// 问题: 给定一个正整数数组 nums 和一个目标正整数 target，请找出所有可能的组合，使得组合中的元素和等于 target。
//      给定数组无重复元素，每个元素可以被选取多次。请以列表形式返回这些组合，列表中不应包含重复组合。
/* 回溯算法：子集和 I */
func backtrackSubsetSumINaive(total, target int, state, choices *[]int, res *[][]int) {
	// 子集和等于 target 时，记录解
	if target == total {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	// 遍历所有选择
	for i := 0; i < len(*choices); i++ {
		// 剪枝：若子集和超过 target ，则跳过该选择
		if total+(*choices)[i] > target {
			continue
		}
		// 尝试：做出选择，更新元素和 total
		*state = append(*state, (*choices)[i])
		// 进行下一轮选择
		backtrackSubsetSumINaive(total+(*choices)[i], target, state, choices, res)
		// 回退：撤销选择，恢复到之前的状态
		*state = (*state)[:len(*state)-1]
	}
}

/* 求解子集和 I（包含重复子集） */
func subsetSumINaive(nums []int, target int) [][]int {
	state := make([]int, 0) // 状态（子集）
	total := 0              // 子集和
	res := make([][]int, 0) // 结果列表（子集列表）
	backtrackSubsetSumINaive(total, target, &state, &nums, &res)
	return res
}
