package main

import (
	"fmt"
	"math"
	"time"
)

const (
	a = 1 << iota
	b
	c
)

var (
	d float32 = math.Pi
	e float64 = math.Pi
)

var (
	f float64 = 3 + 0i
)

func main() {
	f = 2
	fmt.Println(f)
	const noDelay time.Duration = 0
	const num time.Month = 1
	fmt.Println(noDelay, num)
	fmt.Println(a, b, c)
}
