package main

import (
	"fmt"
	"time"
)

func chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func main() {

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive C", c)
		case s := <-ch:
			fmt.Println("Receive S", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
