package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	target := ex + "/src/demo78_target.text"
	source := ex + "/src/demo78_source.text"
	copyFile(target, source)
}

// 文件内容的复制
func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
