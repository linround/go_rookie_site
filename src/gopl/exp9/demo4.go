package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex // guards balance

type test struct {
	name string
}

func (t *test) getName(s int) string {
	fmt.Println(" 开始加锁", s)
	mu.Lock()
	fmt.Println(" 已经加锁", s)
	name := t.name
	fmt.Println(" 已经赋值", s)
	defer mu.Unlock()
	return name
}

// 创建一个加法函数试试
// create
var vt test = test{"linyuan"}

func main() {
	go func() {
		fmt.Println("打印值111:", vt.getName(111))
	}()
	go func() {
		fmt.Println("打印值222:", vt.getName(222))
	}()
	time.Sleep(3 * time.Second)
}
