package main

import (
	"fmt"
	"time"
)

var (
	sema = make(chan struct{}, 1)
)

type balance struct {
	value string
}

var bt balance = balance{"linyuan"}

func (b *balance) getV(s int) string {

	fmt.Println(" 开始加锁", s)
	sema <- struct{}{}
	fmt.Println(" 已经加锁", s)
	value := b.value
	fmt.Println(" 已经赋值", s)
	defer func() {
		<-sema
	}()
	return value
}

func main() {
	go func() {
		fmt.Println("打印值111:", bt.getV(111))
	}()
	go func() {
		fmt.Println("打印值222:", bt.getV(222))
	}()
	time.Sleep(3 * time.Second)
}
