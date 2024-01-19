package main

import (
	"fmt"
	"os"
	"strings"
)

// 修改 dup1 程序：文件读取。

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	for _, filename := range filenames {
		data, err := os.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
