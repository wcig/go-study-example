package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输出标准输入中出现次数大于 1 的行，前面是次数，后面是内容。

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		if line == "end" {
			break
		}
		counts[line]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
