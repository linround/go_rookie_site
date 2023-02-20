package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	// 调用sha1.new 创建了一个新的hash.Hash对象
	// 用来计算sha1校验值
	hasher := sha1.New()
	//通过调用 io.WriteString 或hasher.write
	// 将给定的[]byte添加到当前的hash.Hash对象中
	io.WriteString(hasher, "test")
	b := []byte{}
	res := hasher.Sum(b)
	fmt.Printf("输出值: %x \n", res)
	fmt.Printf("Result: %d\n", res)
	hasher.Reset()
	data := "We shall overcome!"
	n, err := hasher.Write([]byte(data))
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
