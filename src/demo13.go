package main

import "fmt"

type Circle struct {
	radius float64
}

var balance = [4]float64{1, 2, 5, 7}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) getLength() float64 {
	return 3.14 * c.radius * 2
}
func main() {
	var c1 Circle
	c1.radius = 18
	fmt.Print(c1.getArea())
	println()
	fmt.Print(c1.getLength())
}
