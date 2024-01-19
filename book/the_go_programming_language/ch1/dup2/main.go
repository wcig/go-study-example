package main

import (
	"bufio"
	"fmt"
	"os"
)

// 修改 dup1 程序：支持从标准输入读取，也支持从文件读取。

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		countLiensFromFiles(filenames, counts)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if line == "end" {
			break
		}
		counts[line]++
	}
}

func countLiensFromFiles(filenames []string, counts map[string]int) {
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(file, counts)
		_ = file.Close()
	}
}
