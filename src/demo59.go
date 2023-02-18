package main

import "fmt"

type Ad struct {
	a int
}
type Bd struct {
	a, b int
}
type C struct {
	Ad
	Bd
}

func main() {
	var c C
	c.Ad.a = 100
	this := 100
	fmt.Println(c.Ad)
	fmt.Println(this)
}
