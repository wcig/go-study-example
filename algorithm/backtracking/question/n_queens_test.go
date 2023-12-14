// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_backtracking

package question

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	n := 4
	res := nQueens(n)

	fmt.Println("输入棋盘长宽为 ", n)
	fmt.Println("皇后放置方案共有 ", len(res), " 种")
	for _, state := range res {
		fmt.Println("--------------------")
		for _, row := range state {
			fmt.Println(row)
		}
	}

	// Output:
	// 输入棋盘长宽为  4
	// 皇后放置方案共有  2  种
	// --------------------
	// [# Q # #]
	// [# # # Q]
	// [Q # # #]
	// [# # Q #]
	// --------------------
	// [# # Q #]
	// [Q # # #]
	// [# # # Q]
	// [# Q # #]
}
