package main

import (
	"fmt"
	"time"
)

func main() {
	// ...create abort channel...
	abort := make(chan int)
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			fmt.Println(countdown)
			go func() {
				abort <- countdown
			}()
		case <-abort:
			fmt.Println("Launch aborted!")

			return
		}
	}
}
