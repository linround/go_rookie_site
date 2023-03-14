package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}
func bigSlowOperation() {
	// 记录函数的进入和退出实际
	// 记录进入 trace即bigSlowOperation的时间
	// 使用defer记录退出bigSlowOperation的时间
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
