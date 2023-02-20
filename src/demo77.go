package main

import (
	"fmt"
	"os"
)

func main() {

	readFile()
}

// 按照列读取文件
func readFile() {
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	file, err2 := os.Open(ex + "/src/demo77_pro.text")
	// Base用于获取路径中的最后一个元素（不包含后面的分割符）
	//filename := filepath.Base(ex + "/src/demo77_pro.text")
	//fmt.Println(filename)
	if err2 != nil {
		fmt.Println("读取文件出错")
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err3 := fmt.Fscanln(file, &v1, &v2, &v3)
		if err3 != nil {
			fmt.Println("文件扫描出错了")
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)

	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
