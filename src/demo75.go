package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件

func main() {

	readWrite()
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	inputFile, inputError := os.Open(ex + "/src/demo76.text")
	if inputError != nil {
		fmt.Println("读取文件出错了！")
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Println("输入是：", inputString)
		if readerError == io.EOF {
			return
		}
	}

}

// 文件的读取和写入
func readWrite() {
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	inputFile := ex + "/src/demo76.text"
	outputFile := ex + "/src/demo76_output.text"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("读取文件错误 2")
	}
	fmt.Println(string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		fmt.Println("写入文件错误 2")
	}
	fmt.Println("写入完成")

}
