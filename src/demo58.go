package main

import "fmt"

type A struct {
	ax, ay int
}
type B struct {
	A      // 放在同一行，容易造成理解问题
	bx, by float64
}

func main() {
	b := B{A{100, 200}, 300, 400}
	fmt.Println(b)
}
