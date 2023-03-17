package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() { //g7
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() { //g8
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
	time.Sleep(2 * time.Second)
}
