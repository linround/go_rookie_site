package main

import (
	"fmt"
	"sync"
)

var wait = sync.WaitGroup{}

func main() {
	wait.Add(1)
	for i := 0; i < 20; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}
	wait.Done()
	wait.Wait()
	fmt.Println("完成")
}
