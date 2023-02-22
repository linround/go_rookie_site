package main

import (
	"bufio"
	"fmt"
	"os"
)

//读取文件

func readFile() {
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	file, err2 := os.Open(ex + "/src/p16/main8.text")

	if err2 != nil {
		fmt.Println("读取文件错误")
	}
	defer file.Close()
	// 定义一个reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		fmt.Println(str)
		if err != nil {
			return
		}

	}
}
func main() {
	readFile()

	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	file, _ := os.Open(ex + "/src/p16/main8.text")
	readFileBySlice(file)
}
func readFileBySlice(f *os.File) {
	defer f.Close()
	const NBUF = 512
	var buf [NBUF]byte
	for {
		// 将文件读取到buf中
		switch nr, _ := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Println("读取错误")
		case nr == 0:
			return
		case nr > 0:
			if nw, _ := f.Write(buf[0:nr]); nw != nr {
				fmt.Println("写入已经完成", nw)
			}
		}
	}
}
