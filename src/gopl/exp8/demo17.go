package main

import (
	"fmt"
	"time"
)

var abort = make(chan int)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	launch()
	select {
	case <-tick:
		fmt.Println("tick")
	case <-abort:
		{
			fmt.Println("abort")
		}

	}
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

}

func launch() {
	go func() {
		// read a single byte
		abort <- 1000
	}()

}
