package main

import (
	"fmt"
	"time"
)

// 使用完整信号量模式对长度为n的float64切片进行n个计算 并同时完成

type empty interface {
}

func main() {
	const n int = 2
	var em empty
	data := make([]float64, n)
	res := make([]float64, n)
	sem := make(chan empty, n)

	for i, xi := range data {
		fmt.Println("执行", xi, i)
		go func(i int, xi float64) {
			res = append(res, txt(float64(i)))
		}(i, xi)
		sem <- em
	}
	time.Sleep(10e9)
	for i := 0; i < n; i++ {
		// 释放信号量
		<-sem
	}
	println("执行完成")

	fmt.Println(res)
}
func txt(x float64) float64 {
	fmt.Println("开始计算", x)
	time.Sleep(5e9)
	fmt.Println("计算结束", x)
	x++
	return x
}
