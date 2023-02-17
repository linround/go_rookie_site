package main

import (
	"math"
	"strconv"
)

func main() {
	t := math.Sqrt(100)
	println("t的根号：", int(t))

	v, ok := mySqr(100)
	if !ok {
		println("出错了", ok)
	}
	println(int(v))

	println(atoi("456123"))

}
func mySqr(f float64) (v float64, ok bool) {

	if f < 0 {
		return v, true
	}
	return math.Sqrt(f), false
}

func atoi(s string) (v int) {
	n, _ := strconv.Atoi(s)
	return n
}
