package main

import "fmt"

//惰性生成器的实现

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			if count > 5 {
				break
			}
			count++
		}
	}()
	return yield
}

func generateInter() int {
	return <-resume
}
func main() {
	resume = integers()
	fmt.Println(generateInter())
	fmt.Println(generateInter())
	fmt.Println(generateInter())
	fmt.Println(generateInter())
	fmt.Println(generateInter())
}
