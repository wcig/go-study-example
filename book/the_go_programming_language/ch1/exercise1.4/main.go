package main

import (
	"bufio"
	"fmt"
	"os"
)

// 修改 dup2 程序：输出出现重复行文件名称。

func main() {
	counts := make(map[string]map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		countLiensFromFiles(filenames, counts)
	}
	for filename, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%s: %d: %s\n", filename, n, line)
			}
		}
	}
}

func countLines(file *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(file)
	count := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		if line == "end" {
			break
		}
		count[line]++
	}
	filename := file.Name()
	counts[filename] = count
}

func countLiensFromFiles(filenames []string, counts map[string]map[string]int) {
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
