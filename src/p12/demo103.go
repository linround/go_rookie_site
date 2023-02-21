package main

import "fmt"

type any interface {
}
type evalFunc func(any) (any, any)

func main() {
	evenF := func(s any) (any, any) {
		os := s.(int)
		ns := os + 2
		return os, ns
	}
	even := buildLazyIntEvaluator(evenF, 0)
	for i := 0; i < 10; i++ {
		fmt.Println("偶数", i, even())
	}
}

// 惰性生产器
// 传入的函数需要计算下一个返回值以及下一个状态参数

func BuildLazyEvaluator(evalFunc2 evalFunc, initState any) func() any {
	// 通道
	retValChan := make(chan any)
	loopF := func() {
		var ac any = initState
		var retVal any
		// 无限循环的go协程
		for {
			retVal, ac = evalFunc2(ac)
			retValChan <- retVal
		}
	}
	retFunc := func() any {
		return <-retValChan
	}
	go loopF()
	return retFunc
}
func buildLazyIntEvaluator(evalFunc2 evalFunc, initState any) func() int {
	ef := BuildLazyEvaluator(evalFunc2, initState)
	return func() int {
		return ef().(int)
	}
}
