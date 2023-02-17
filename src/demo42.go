package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getx2andx32(10)
	Prit()
	varNum(1, 2, 5, 8, 9, 6, 3)
}

func getx2andx32(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func Prit() {
	println(num, numx2, numx3)
}

func varNum(a ...int) {
	fmt.Print(a)
}
