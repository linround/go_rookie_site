package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println(counts, "counts:")
}

// f 是被打开的文件
func countLines(f *os.File, counts map[string]int) {
	// 使用scanner进行读取

	// 统计相同的行数
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println(input)
}
